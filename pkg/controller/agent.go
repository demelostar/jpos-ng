package controller

import (
	"fmt"
	"github.com/hashicorp/yamux"
	"github.com/demelostar/ljpos-li/pkg/protocol"
	"net"
)

var AgentCounter = 0
var ListenerCounter = 0

type LjposAgent struct {
	Id        int
	Name      string
	Network   []protocol.NetInterface
	Session   *yamux.Session
	CloseChan chan bool
	Interface string
	Running   bool
}

type Listener struct {
	Agent        LjposAgent
	Network      string
	ListenerAddr string
	RedirectAddr string

	Session    net.Conn
	ListenerID int32
}

func (l Listener) String() string {
	return fmt.Sprintf("#%d [%s] (%s) [Agent] %s => [Proxy] %s", l.Agent.Id, l.Agent.Name, l.Network, l.ListenerAddr, l.RedirectAddr)
}

func (la *LjposAgent) String() string {
	return fmt.Sprintf("#%d - %s - %s", la.Id, la.Name, la.Session.RemoteAddr())
}

func NewAgent(session *yamux.Session) (*LjposAgent, error) {
	yamuxConnectionSession, err := session.Open()
	if err != nil {
		return nil, err
	}

	infoRequest := protocol.InfoRequestPacket{}

	protocolEncoder := protocol.NewEncoder(yamuxConnectionSession)
	protocolDecoder := protocol.NewDecoder(yamuxConnectionSession)

	if err := protocolEncoder.Encode(protocol.Envelope{
		Type:    protocol.MessageInfoRequest,
		Payload: infoRequest,
	}); err != nil {
		return nil, err
	}

	if err := protocolDecoder.Decode(); err != nil {
		return nil, err
	}

	response := protocolDecoder.Envelope.Payload
	reply := response.(protocol.InfoReplyPacket)
	AgentCounter++
	return &LjposAgent{
		Id:        AgentCounter,
		Name:      reply.Name,
		Network:   reply.Interfaces,
		Session:   session,
		CloseChan: make(chan bool),
	}, nil
}
