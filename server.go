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
			fmt.Println("error accepting connection: ", err)
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	req := InitRequest()
	res := newResponseWriter(conn, conn, req)

	allHeadersParsed := false
	for {
		contentLength, err := req.contentLength()
		if err != nil {
			log := fmt.Sprintf("error reading content length: %s", err.Error())
			requestErrorHandler(res, req, log, internalServerErrorHandler)
			return
		}
		if contentLength == len(req.Body) && allHeadersParsed {
			break
		}

		n, err := conn.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log := fmt.Sprintf("error reading request: %s", err.Error())
			requestErrorHandler(res, req, log, internalServerErrorHandler)
			return
		}

		allHeadersParsed = req.ParseRequestMessage(buffer[:n])
	}

	for h, v := range req.Headers {
		fmt.Printf("%s: %s\n", h, v)
	}

	if req.Method != GET {
		log := "method not allowed"
		requestErrorHandler(res, req, log, methodNotAllowed)
		return
	}

	handler := s.router.getHandler(req.Path, req.Method)
	if handler == nil {
		handler = notFoundHandler
	}
	handler(res, req)
	res.flushResponse()
}

func (s *Server) addHandler(pattern string, handler Handler) {
	route := strings.Split(pattern, " ")
	s.router.addRoute(route[1], route[0], handler)
}

func requestErrorHandler(rw ResponseWriter, req *Request, err string, handler Handler) {
	handler(rw, req)
	fmt.Println(err)
	res, ok := rw.(*response)
	if !ok {
		return
	}
	res.flushResponse()
}
