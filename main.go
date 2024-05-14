package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/valrobichaux/Distributed-File-Storage/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	TCPTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		//TODO: onPeer func
	}

	tcpTransport := p2p.NewTCPTransport(TCPTransportOpts)

	storageRoot := strings.TrimPrefix(listenAddr, ":")

	fileServeropts := FileServerOpts{
		StorageRoot:       storageRoot + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}
	s := NewFileServer(fileServeropts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")

	go func() { log.Fatal(s1.Start()) }()

	time.Sleep(2 * time.Second)

	go s2.Start()
	time.Sleep(2 * time.Second)

	// data := bytes.NewReader([]byte("My big data filer here!"))
	// s2.Store("myprivatedata", data)

	r, err := s2.Get("myprivatedata")
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	select {}
}
