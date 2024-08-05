package proxy

import (
	"context"
	"github.com/hashicorp/yamux"
	"github.com/demelostar/ljpos-li/pkg/proxy/netstack"
	"github.com/sirupsen/logrus"
)

const (
	MaxConnectionHandler = 4096
)

type LjposTunnel struct {
	nstack *netstack.NetStack
}

func NewLjposTunnel(stackSettings netstack.StackSettings) (*LjposTunnel, error) {
	// Create a new stack, but without connPool.
	// The connPool will be created when using the *start* command
	nstack, err := netstack.NewStack(stackSettings, nil)
	if err != nil {
		return nil, err
	}
	return &LjposTunnel{nstack: nstack}, nil
}

func (t *LjposTunnel) HandleSession(session *yamux.Session, ctx context.Context) {

	// Create a new, empty, connpool to store connections/packets
	connPool := netstack.NewConnPool(MaxConnectionHandler)
	t.nstack.SetConnPool(&connPool)

	// Cleanup pool if channel is closed
	defer connPool.Close()

	for {
		select {
		case <-ctx.Done():
			t.Close()
			return
		case <-connPool.CloseChan: // pool closed, we can't process packets!
			logrus.Infof("Connection pool closed")
			t.Close()
			return
		case tunnelPacket := <-connPool.Pool: // Process connections/packets
			go netstack.HandlePacket(t.nstack.GetStack(), tunnelPacket, session)
		}
	}
}

func (t *LjposTunnel) GetStack() *netstack.NetStack {
	return t.nstack
}

func (t *LjposTunnel) Close() {
	t.nstack.Close()
}
