package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"time"
)

type ResponseWriter interface {
	Write(data []byte) (int, error)
	Header() Header
	WriteHeader(statusCode int)
}

type response struct {
	conn          net.Conn
	w             *bufio.Writer
	req           *Request
	handlerHeader Header
	data          []byte
	startLine     bytes.Buffer
}

func newResponseWriter(conn net.Conn, req *Request) *response {
	return &response{
		conn:          conn,
		w:             bufio.NewWriter(conn),
		req:           req,
		handlerHeader: make(Header),
		data:          []byte{},
		startLine:     bytes.Buffer{},
	}
}

func (r *response) Write(data []byte) (int, error) {
	r.data = append(r.data, data...)
	return len(r.data), nil
}

func (r response) WriteHeader(statusCode int) {
	startLine := fmt.Sprintf("%s %s %s\r\n", r.req.V, strconv.Itoa(statusCode), StatusText(statusCode))
	r.w.Write([]byte(startLine))
}

func (r *response) Header() Header { return r.handlerHeader }

func (r *response) flushResponse() {
	r.Header().Set("Date", time.Now().UTC().Format(TimeFormat))
	r.w.Write([]byte(r.Header().String()))
	r.w.Write(r.data)
	r.w.Flush()
}
