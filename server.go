package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/ahmadjavaidwork/bud/request"
)

type Server struct {
	Addr string
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		// Wait for a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	r := request.InitRequest()

	allHeadersRead := false
	for {
		if allHeadersRead {
			break
		}

		n, err := conn.Read(buffer)
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		allHeadersRead = r.MakeRequest(string(buffer))
	}

	conn.Write(buffer)
}
