package network

import (
	"errors"
	"net"
	"testing"
)

type StubListenner struct {
	CloseCalled bool
}

func (l StubListenner) Accept() (net.Conn, error) {
	return nil, nil
}

func (l StubListenner) Addr() net.Addr {
	return nil
}

func (l StubListenner) Close() error {
	return errors.New("Server closed.")
}
func TestTCPServer(t *testing.T) {
	t.Run("Run returns error if cannot listen the network", func(t *testing.T) {
		tcpserver := &TCPServer{
			Addr:   "test",
			server: &StubListenner{},
		}
		err := tcpserver.Run()
		if err == nil {
			t.Error("Expected to get an error")
		}

	})
}
