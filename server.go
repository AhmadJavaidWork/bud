package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

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

	r := InitRequest()
	w := newResponseWriter(conn, conn, r)

	allHeadersParsed := false
	for {
		if allHeadersParsed {
			break
		}

		_, err := conn.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			internalServerErrorHandler(w, r)
			log.Fatal(err)
		}

		allHeadersParsed = r.ParseRequestMessage(buffer)
	}

	for h, v := range r.Headers {
		fmt.Printf("%s: %s\n", h, v)
	}

	if r.Method != GET {
		methodNotAllowed(w, r)
		w.flushResponse()
		conn.Close()
		return
	}

	handler := s.router.getHandler(r.Path, r.Method)
	if handler == nil {
		handler = notFoundHandler
	}
	handler(w, r)
	w.flushResponse()

	conn.Close()
}

func (s *Server) addHandler(pattern string, handler Handler) {
	route := strings.Split(pattern, " ")
	s.router.addRoute(route[1], route[0], handler)
}
