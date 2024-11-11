package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/ahmadjavaidwork/bud/request"
)

type Server struct {
	Addr   string
	router *Router
}

func NewServer(addr string) *Server {
	return &Server{
		Addr:   addr,
		router: NewRouter(),
	}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp4", s.Addr)
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
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	r := request.InitRequest()

	allHeadersParsed := false
	for {
		if allHeadersParsed {
			break
		}

		n, err := conn.Read(buffer)
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		allHeadersParsed = r.ParseRequestMessage(buffer)
	}

	for h, v := range r.Headers {
		fmt.Printf("%s: %s\n", h, v)
	}
	handler := s.router.getHandler(r.Path, r.Method)
	if handler != nil {
		handler()
	}
	conn.Close()
}

func (s *Server) addHandler(pattern string, handler func()) {
	route := strings.Split(pattern, " ")
	s.router.addRoute(route[1], route[0], handler)
}
