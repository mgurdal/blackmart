package network

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/mgurdal/blackmarkt/store"
	"github.com/mgurdal/blackmarkt/user"
)

// Server defines the minimum contract our
// TCP and UDP server implementations must satisfy.
type Server interface {
	Run() error
	Close() error
}

// NewServer creates a new TCPServer
func NewServer(addr string) (Server, error) {
	return &TCPServer{
		Addr: addr,
	}, nil
}

// TCPServer holds the structure of our TCP
// implementation.
type TCPServer struct {
	Addr   string
	server net.Listener
}

// Run starts the TCP Server.
func (t *TCPServer) Run() (err error) {
	t.server, err = net.Listen("tcp", t.Addr)
	log.Printf("Running TCP Server at %s", t.Addr)
	if err != nil {
		return err
	}
	defer t.Close()

	for {

		conn, err := t.server.Accept()

		if err != nil {
			err = errors.New("could not accept connection")
			break
		}
		if conn == nil {
			err = errors.New("could not create connection")
			break
		}
		usr := &user.User{
			ID:   uuid.New(),
			Conn: conn,
		}
		storage := store.GetStore()
		storage.Register(usr.ID, usr)
		go t.handleConnection(usr)
	}
	return
}

// Close shuts down the TCP Server
func (t *TCPServer) Close() (err error) {
	return t.server.Close()
}

// handleConnections deals with the business logic of
// each connection and their requests.
func (t *TCPServer) handleConnection(usr *user.User) {
	defer usr.Conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(usr.Conn), bufio.NewWriter(usr.Conn))
	for {
		req, err := rw.ReadString('\n')
		if err != nil {
			rw.WriteString("failed to read input")
			rw.Flush()
			return
		}
		log.Printf("Request received from user %s: %s", usr.ID, req)
		rw.WriteString(fmt.Sprintf("Request received from user %s: %s", usr.ID, req))
		rw.Flush()
	}
}
