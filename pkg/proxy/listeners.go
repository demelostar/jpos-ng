package proxy

import (
	"context"
	"errors"
	"github.com/hashicorp/yamux"
	"github.com/demelostar/ljpos-li/pkg/protocol"
	"github.com/demelostar/ljpos-li/pkg/relay"
	"github.com/sirupsen/logrus"
	"io"
	"net"
)

func ListenerStop(sess *yamux.Session, listenerId int32) error {
	// Open Yamux connection
	yamuxConnectionSession, err := sess.Open()
	if err != nil {
		return err
	}

	ljposProtocol := protocol.NewEncoderDecoder(yamuxConnectionSession)

	// Send close request
	closeRequest := protocol.ListenerCloseRequestPacket{ListenerID: listenerId}
	if err := ljposProtocol.Encode(protocol.Envelope{
		Type:    protocol.MessageListenerCloseRequest,
		Payload: closeRequest,
	}); err != nil {
		return err
	}

	// Process close response
	if err := ljposProtocol.Decode(); err != nil {
		return err

	}
	response := ljposProtocol.Envelope.Payload

	if err := response.(protocol.ListenerCloseResponsePacket).Err; err != false {
		return errors.New(response.(protocol.ListenerCloseResponsePacket).ErrString)
	}
	return nil
}

type LjposListener struct {
	ID      int32
	ctx     context.Context
	sess    *yamux.Session
	Conn    net.Conn
	addr    string
	network string
	to      string
}

func NewListener(sess *yamux.Session, addr string, network string, to string) (LjposListener, error) {
	// Open a new Yamux Session
	conn, err := sess.Open()
	if err != nil {
		return LjposListener{}, err
	}

	ljposProtocol := protocol.NewEncoderDecoder(conn)

	// Request to open a new port on the agent
	listenerPacket := protocol.ListenerRequestPacket{Address: addr, Network: network}
	if err := ljposProtocol.Encode(protocol.Envelope{
		Type:    protocol.MessageListenerRequest,
		Payload: listenerPacket,
	}); err != nil {
		return LjposListener{}, err
	}

	// Get response from agent
	if err := ljposProtocol.Decode(); err != nil {
		return LjposListener{}, err
	}
	response := ljposProtocol.Envelope.Payload.(protocol.ListenerResponsePacket)
	if err := response.Err; err != false {
		return LjposListener{}, errors.New(response.ErrString)
	}
	return LjposListener{ID: response.ListenerID, sess: sess, Conn: conn, addr: addr, network: network, to: to}, nil
}

func (l *LjposListener) StartRelay() error {
	if l.network == "tcp" {
		return l.relayTCP()
	} else if l.network == "udp" {
		return l.relayUDP()
	}
	return errors.New("invalid network")
}

func (l *LjposListener) relayTCP() error {
	ljposProtocol := protocol.NewEncoderDecoder(l.Conn)
	for {
		// Wait for BindResponses
		if err := ljposProtocol.Decode(); err != nil {
			if err == io.EOF {
				// Listener closed.
				return nil
			}
			return err
		}

		// We received a new BindResponse!
		response := ljposProtocol.Envelope.Payload.(protocol.ListenerBindReponse)

		if err := response.Err; err != false {
			return errors.New(response.ErrString)
		}

		logrus.Debugf("New socket opened : %d", response.SockID)

		// relay connection
		go func(sockID int32) {
			forwarderSession, err := l.sess.Open()
			if err != nil {
				logrus.Error(err)
				return
			}

			forwarderProtocolEncDec := protocol.NewEncoderDecoder(forwarderSession)

			// Request socket access
			socketRequestPacket := protocol.ListenerSockRequestPacket{SockID: sockID}
			if err := forwarderProtocolEncDec.Encode(protocol.Envelope{
				Type:    protocol.MessageListenerSockRequest,
				Payload: socketRequestPacket,
			}); err != nil {
				logrus.Error(err)
				return
			}
			if err := forwarderProtocolEncDec.Decode(); err != nil {
				logrus.Error(err)
				return
			}

			response := forwarderProtocolEncDec.Envelope.Payload
			if err := response.(protocol.ListenerSockResponsePacket).Err; err != false {
				logrus.Error(response.(protocol.ListenerSockResponsePacket).ErrString)
				return
			}
			// Got socket access!

			logrus.Debug("Listener relay established!")

			// Dial the "to" target
			lconn, err := net.Dial(l.network, l.to)
			if err != nil {
				logrus.Error(err)
				return
			}

			// relay connections
			if err := relay.StartRelay(lconn, forwarderSession); err != nil {
				logrus.Error(err)
				return
			}
		}(response.SockID)

	}

}

func (l *LjposListener) relayUDP() error {
	// Dial the "to" target
	lconn, err := net.Dial(l.network, l.to)
	if err != nil {
		return err
	}
	// Relay conn
	err = relay.StartPacketRelay(lconn, l.Conn)
	if err != nil {
		return err
	}
	return nil
}
