package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Request struct {
	V          string // http version
	Method     string
	Headers    map[string]string
	Path       string
	Body       []byte
	prevBuffer []byte // to be processed buffer
}

func InitRequest() *Request {
	return &Request{Headers: map[string]string{}}
}

func (req *Request) ParseRequestMessage(buffer []byte) bool {
	req.parseStartLine(buffer)
	allHeadersParsed := req.parseHeaders()
	req.parseBody(allHeadersParsed)
	return allHeadersParsed
}

func (req *Request) GetHeader(name string) string {
	if h, ok := req.Headers[name]; ok {
		return h
	}
	return ""
}

func (req *Request) setHeader(name string, value string) {
	req.Headers[name] = value
}

func (req *Request) ContainsBody() (bool, error) {
	length, err := req.contentLength()
	if err != nil {
		return false, err
	}
	return length > 0, nil
}

func (req *Request) contentLength() (int, error) {
	cl := req.GetHeader("Content-Length")
	if cl == "" {
		return 0, nil
	}
	length, err := strconv.ParseInt(cl, 0, 64)
	if err != nil {
		return 0, err
	}
	if length > int64(math.Pow(2, 16)) {
		return 0, fmt.Errorf("max content length exceeded: %d", length)
	}
	return int(length), nil
}

func (req *Request) isStartLineParsed() bool {
	return req.Method != "" && req.Path != "" && req.V != ""
}

func (req *Request) startLine() string {
	return fmt.Sprintf("%s %s %s\r\n", req.Method, req.Path, req.V)
}

func (req *Request) parseStartLine(buffer []byte) {
	req.prevBuffer = append(req.prevBuffer, buffer...)
	if req.isStartLineParsed() {
		return
	}

	l := 0
	if !strings.Contains((string(req.prevBuffer)), "\r\n") {
		return
	}

	for r := 0; req.prevBuffer[r] != '\n'; r++ {
		if req.prevBuffer[r] != ' ' && req.prevBuffer[r] != '\r' {
			continue
		}

		if req.Method == "" {
			req.Method = string(req.prevBuffer[l:r])
		} else if req.Path == "" {
			req.Path = string(req.prevBuffer[l:r])
		} else {
			req.V = string(req.prevBuffer[l:r])
		}
		l = r + 1
	}

	req.prevBuffer = req.prevBuffer[l+1:]
}

func (req *Request) parseHeaders() bool {
	if !req.isStartLineParsed() {
		return false
	}

	if !strings.Contains((string(req.prevBuffer)), "\r\n") {
		return false
	}

	l := 0
	allheadersParsed := false

	for r := 0; r < len(req.prevBuffer); r++ {
		if req.prevBuffer[r] != '\r' {
			continue
		}
		if req.prevBuffer[l] == '\n' {
			l++
		}

		pair := strings.Split(string(req.prevBuffer[l:r]), ":")
		if len(pair[0]) == 0 {
			allheadersParsed = true
			break
		}

		if pair[1][0] == ' ' {
			pair[1] = pair[1][1:]
		}
		req.setHeader(pair[0], pair[1])
		l = r + 1
	}

	req.prevBuffer = req.prevBuffer[l:]
	return allheadersParsed
}

func (req *Request) parseBody(allHeadersParsed bool) error {
	if !allHeadersParsed {
		return nil
	}

	contentLength, err := req.contentLength()
	if err != nil {
		return err
	}

	if contentLength+2 != len(req.prevBuffer) {
		return nil
	}

	req.prevBuffer = req.prevBuffer[2:]
	req.Body = req.prevBuffer[:contentLength]

	return nil
}
