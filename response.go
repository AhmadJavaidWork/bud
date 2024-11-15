package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
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
	startLine     string
	body          []byte
}

func newResponse(conn net.Conn, w io.Writer, req *Request) *response {
	return &response{
		conn:          conn,
		w:             bufio.NewWriter(w),
		req:           req,
		handlerHeader: make(Header),
		data:          []byte{},
		body:          []byte{},
		startLine:     "",
	}
}

func (res *response) Write(data []byte) (int, error) {
	res.body = append(res.body, data...)
	return len(res.body), nil
}

func (res response) WriteHeader(statusCode int) {
	res.startLine = fmt.Sprintf("%s %d %s\r\n", res.req.V, statusCode, StatusText(statusCode))
	res.w.Write([]byte(res.startLine))
}

func (res *response) Header() Header { return res.handlerHeader }

func (res *response) flushResponse() {
	res.Header().Set("Date", time.Now().UTC().Format(TimeFormat))

	contentType := http.DetectContentType(res.body)
	res.Header().Set("Content-Type", contentType)
	res.Header().Set("Content-Length", strconv.Itoa(len(res.body)))

	res.w.Write([]byte(res.Header().String()))
	res.w.Write(res.body)
	res.w.Flush()
}
