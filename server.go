package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
const ReadDeadline = time.Second * 5

type Server struct {
	Addr            string
	router          *Router
	ReadDeadline    time.Duration
	openConnections int
}

func NewServer(addr string) *Server {
	return &Server{
		Addr:         addr,
		router:       NewRouter(),
		ReadDeadline: 5 * time.Second,
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
	defer func() {
		conn.Close()
		s.openConnections--
	}()

	s.openConnections++

	for {
		buffer := make([]byte, 1024)

		req := InitRequest()
		res := newResponse(conn, conn, req)

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
			delete(req.Headers, "Connection")
		}
		handler(res, req)
		res.flushResponse()

		if !req.shouldKeepAlive() {
			break
		}
	}
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

func (s *Server) SetReadDeadline(readDeadline time.Duration) {
	s.ReadDeadline = readDeadline
}
