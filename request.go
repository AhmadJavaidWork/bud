package main

import (
	"fmt"
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
	allHeadersParsed := false
	cummulativeBuffer := append(req.prevBuffer, buffer...)
	l := 0

	for r := 0; r < len(cummulativeBuffer) && !req.startLineDone(); r++ {
		if cummulativeBuffer[r] != ' ' && cummulativeBuffer[r] != '\r' {
			continue
		}
		if req.Method == "" && cummulativeBuffer[r] == ' ' {
			req.Method = string(cummulativeBuffer[l:r])
		} else if req.Path == "" && cummulativeBuffer[r] == ' ' {
			req.Path = string(cummulativeBuffer[l:r])
		} else if req.V == "" && cummulativeBuffer[r] == '\r' {
			req.V = string(cummulativeBuffer[l:r])
			cummulativeBuffer = cummulativeBuffer[r+1:]
			break
		}
		l = r + 1
	}

	if cummulativeBuffer[0] == '\n' {
		cummulativeBuffer = cummulativeBuffer[1:]
	}
	l = 0

	for r := 0; r < len(cummulativeBuffer); r++ {
		if cummulativeBuffer[r] != '\r' {
			continue
		}
		if cummulativeBuffer[l] == '\n' {
			l++
		}
		s := strings.Split(string(cummulativeBuffer[l:r]), ":")
		if len(s[0]) == 0 {
			allHeadersParsed = true
			break
		}
		if s[1][0] == ' ' {
			s[1] = s[1][1:]
		}
		req.setHeader(s[0], s[1])
		l = r + 1
	}

	req.prevBuffer = cummulativeBuffer[l:]

	return allHeadersParsed
}

func (r *Request) GetHeader(name string) string {
	if h, ok := r.Headers[name]; ok {
		return h
	}
	return ""
}

func (r *Request) setHeader(name string, value string) {
	r.Headers[name] = value
}

func (r *Request) ContainsBody() (bool, error) {
	cl := r.GetHeader("Content-Length")
	if cl == "" {
		return false, nil
	}
	length, err := strconv.ParseInt(cl, 0, 64)
	if err != nil {
		return false, err
	}
	return length > 0, nil
}

func (req *Request) startLineDone() bool {
	return req.Method != "" && req.Path == "" && req.V != ""
}

func (r *Request) startLine() string {
	return fmt.Sprintf("%s %s %s\r\n", r.Method, r.Path, r.V)
}
