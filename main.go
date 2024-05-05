package main

import (
	"log"

	"github.com/valrobichaux/Distributed-File-Storage/p2p"
)

func main() {
	TCPTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		//TODO: onPeer func
	}

	tcpTransport := p2p.NewTCPTransport(TCPTransportOpts)

	fileServeropts := FileServerOpts{
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}
	s := NewFileServer(fileServeropts)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	select {}
}
