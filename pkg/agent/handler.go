package agent

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/user"
	"syscall"
	"time"

	"github.com/demelostar/ljpos-li/pkg/agent/neterror"
	"github.com/demelostar/ljpos-li/pkg/agent/smartping"
	"github.com/demelostar/ljpos-li/pkg/protocol"
	"github.com/demelostar/ljpos-li/pkg/relay"
	"github.com/sirupsen/logrus"
)

var listenerConntrack map[int32]net.Conn
var listenerMap map[int32]interface{}
var connTrackID int32
var listenerID int32

func init() {
	listenerConntrack = make(map[int32]net.Conn)
	listenerMap = make(map[int32]interface{})
}

// Listener is the base class implementing listener sockets for Ljpos
type Listener struct {
	net.Listener
}

// NewListener register a new listener
func NewListener(network string, addr string) (Listener, error) {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return Listener{}, err
	}
	return Listener{lis}, nil
}

// ListenAndServe fill new listener connections to a channel
func (s *Listener) ListenAndServe(connTrackChan chan int32) error {
	for {
		conn, err := s.Accept()
		if err != nil {
			return err
		}
		connTrackID++
		connTrackChan <- connTrackID
		listenerConntrack[connTrackID] = conn
	}
}

// Close request the main listener to exit
func (s *Listener) Close() error {
	return s.Listener.Close()
}

// UDPListener is the base class implementing UDP listeners for Ljpos
type UDPListener struct {
	*net.UDPConn
}

// NewUDPListener register a new UDP listener
func NewUDPListener(network string, addr string) (UDPListener, error) {
	udpaddr, err := net.ResolveUDPAddr(network, addr)
	if err != nil {
		return UDPListener{}, nil
	}

	udplis, err := net.ListenUDP(network, udpaddr)
	if err != nil {
		return UDPListener{}, err
	}
	return UDPListener{udplis}, err
}

func HandleConn(conn net.Conn) {
	decoder := protocol.NewDecoder(conn)
	if err := decoder.Decode(); err != nil {
		logrus.Error(err)
		return
	}

	e := decoder.Envelope.Payload
	switch decoder.Envelope.Type {

	case protocol.MessageConnectRequest:
		connRequest := e.(protocol.ConnectRequestPacket)
		encoder := protocol.NewEncoder(conn)

		logrus.Debugf("Got connect request to %s:%d", connRequest.Address, connRequest.Port)
		var network string
		if connRequest.Transport == protocol.TransportTCP {
			network = "tcp"
		} else {
			network = "udp"
		}
		if connRequest.Net == protocol.Networkv4 {
			network += "4"
		} else {
			network += "6"
		}

		var d net.Dialer
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		targetConn, err := d.DialContext(ctx, network, fmt.Sprintf("%s:%d", connRequest.Address, connRequest.Port))
		defer cancel()

		var connectPacket protocol.ConnectResponsePacket
		if err != nil {

			var serr syscall.Errno
			if errors.As(err, &serr) {
				// Magic trick ! If the error syscall indicate that the system responded, send back a RST packet!
				if neterror.HostResponded(serr) {
					connectPacket.Reset = true
				}
			}

			connectPacket.Established = false
		} else {
			connectPacket.Established = true
		}
		if err := encoder.Encode(protocol.Envelope{
			Type:    protocol.MessageConnectResponse,
			Payload: connectPacket,
		}); err != nil {
			logrus.Error(err)
			return
		}
		if connectPacket.Established {
			relay.StartRelay(targetConn, conn)
		}
	case protocol.MessageHostPingRequest:
		pingRequest := e.(protocol.HostPingRequestPacket)
		encoder := protocol.NewEncoder(conn)

		pingResponse := protocol.HostPingResponsePacket{Alive: smartping.TryResolve(pingRequest.Address)}

		if err := encoder.Encode(protocol.Envelope{
			Type:    protocol.MessageHostPingResponse,
			Payload: pingResponse,
		}); err != nil {
			logrus.Error(err)
			return
		}
	case protocol.MessageInfoRequest:
		var username string
		encoder := protocol.NewEncoder(conn)
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "UNKNOWN"
		}

		userinfo, err := user.Current()
		if err != nil {
			username = "Unknown"
		} else {
			username = userinfo.Username
		}

		netifaces, err := net.Interfaces()
		if err != nil {
			logrus.Error("could not get network interfaces")
			return
		}
		infoResponse := protocol.InfoReplyPacket{
			Name:       fmt.Sprintf("%s@%s", username, hostname),
			Interfaces: protocol.NewNetInterfaces(netifaces),
		}

		if err := encoder.Encode(protocol.Envelope{
			Type:    protocol.MessageInfoReply,
			Payload: infoResponse,
		}); err != nil {
			logrus.Error(err)
			return
		}
	case protocol.MessageListenerCloseRequest:
		// Request to close a listener
		closeRequest := e.(protocol.ListenerCloseRequestPacket)
		encoder := protocol.NewEncoder(conn)

		var err error
		if lis, ok := listenerMap[closeRequest.ListenerID]; ok {
			if l, ok := lis.(net.Listener); ok {
				l.Close()
			}
			if l, ok := lis.(*net.UDPConn); ok {
				l.Close()
			}
		} else {
			err = errors.New("invalid listener id")
		}

		listenerResponse := protocol.ListenerCloseResponsePacket{
			Err: err != nil,
		}
		if err != nil {
			listenerResponse.ErrString = err.Error()
		}

		if err := encoder.Encode(protocol.Envelope{
			Type:    protocol.MessageListenerCloseResponse,
			Payload: listenerResponse,
		}); err != nil {
			logrus.Error(err)
		}

	case protocol.MessageListenerRequest:
		listenRequest := e.(protocol.ListenerRequestPacket)
		encoder := protocol.NewEncoder(conn)
		connTrackChan := make(chan int32)
		stopChan := make(chan error)

		if listenRequest.Network == "tcp" {
			listener, err := NewListener(listenRequest.Network, listenRequest.Address)
			if err != nil {
				listenerResponse := protocol.ListenerResponsePacket{
					ListenerID: 0,
					Err:        true,
					ErrString:  err.Error(),
				}
				if err := encoder.Encode(protocol.Envelope{
					Type:    protocol.MessageListenerResponse,
					Payload: listenerResponse,
				}); err != nil {
					logrus.Error(err)
				}
				return
			}
			listenerMap[listenerID] = listener.Listener
			listenerResponse := protocol.ListenerResponsePacket{
				ListenerID: listenerID,
				Err:        false,
				ErrString:  "",
			}
			if err := encoder.Encode(protocol.Envelope{
				Type:    protocol.MessageListenerResponse,
				Payload: listenerResponse,
			}); err != nil {
				logrus.Error(err)
			}
			go func() {
				if err := listener.ListenAndServe(connTrackChan); err != nil {
					stopChan <- err
				}
			}()
			defer listener.Close()

		} else if listenRequest.Network == "udp" {
			udplistener, err := NewUDPListener(listenRequest.Network, listenRequest.Address)
			if err != nil {
				listenerResponse := protocol.ListenerResponsePacket{
					ListenerID: 0,
					Err:        true,
					ErrString:  err.Error(),
				}
				if err := encoder.Encode(protocol.Envelope{
					Type:    protocol.MessageListenerResponse,
					Payload: listenerResponse,
				}); err != nil {
					logrus.Error(err)
				}
				return
			}
			listenerMap[listenerID] = udplistener.UDPConn
			listenerResponse := protocol.ListenerResponsePacket{
				ListenerID: listenerID,
				Err:        false,
				ErrString:  "",
			}
			if err := encoder.Encode(protocol.Envelope{
				Type:    protocol.MessageListenerResponse,
				Payload: listenerResponse,
			}); err != nil {
				logrus.Error(err)
			}
			go relay.StartRelay(conn, udplistener)
		}

		listenerID++
		if listenRequest.Network == "tcp" {
			for {
				var bindResponse protocol.ListenerBindReponse
				select {
				case err := <-stopChan:
					logrus.Error(err)
					bindResponse = protocol.ListenerBindReponse{
						SockID:    0,
						Err:       true,
						ErrString: err.Error(),
					}
				case connTrackID := <-connTrackChan:
					bindResponse = protocol.ListenerBindReponse{
						SockID: connTrackID,
						Err:    false,
					}
				}

				if err := encoder.Encode(protocol.Envelope{
					Type:    protocol.MessageListenerBindResponse,
					Payload: bindResponse,
				}); err != nil {
					logrus.Error(err)
				}

				if bindResponse.Err {
					break
				}

			}
		}
	case protocol.MessageListenerSockRequest:
		sockRequest := e.(protocol.ListenerSockRequestPacket)
		encoder := protocol.NewEncoder(conn)

		var sockResponse protocol.ListenerSockResponsePacket
		if _, ok := listenerConntrack[sockRequest.SockID]; !ok {
			// Handle error
			sockResponse.ErrString = "invalid or unexistant SockID"
			sockResponse.Err = true
		}

		if err := encoder.Encode(protocol.Envelope{
			Type:    protocol.MessageListenerSockResponse,
			Payload: sockResponse,
		}); err != nil {
			logrus.Error(err)
			return
		}

		if sockResponse.Err {
			return
		}

		netConn := listenerConntrack[sockRequest.SockID]
		relay.StartRelay(netConn, conn)

	case protocol.MessageClose:
		os.Exit(0)

	}
}
