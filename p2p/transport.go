package p2p

import "net"

// Peer is an interface that represents the remote node
type Peer interface {
	RemoteAddr() net.Addr
	Close() error
}

// Transport is anything that handles the communication
// between the nodes in the network.  This can be of the form (TCP, UDP, Websockets, ...)
type Transport interface {
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
