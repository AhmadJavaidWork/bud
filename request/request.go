package request

import (
	"strconv"
	"strings"
)

const lineBreak = "\r\n"

type Request struct {
	V       string
	Method  string
	Headers map[string]string
	Path    string
	Body    []byte
}

func InitRequest() *Request {
	return &Request{Headers: map[string]string{}}
}

func (r *Request) MakeRequest(input string) bool {
	allHeadersRead := false

	for i, l := range strings.Split(input, lineBreak) {
		if len(l) == 0 {
			allHeadersRead = true
			break
		}
		if i == 0 {
			startLine := strings.Split(l, " ")
			r.Method = MethodsLookUp(startLine[0])
			r.Path = startLine[1]
			r.V = startLine[2]
		} else {
			headerLine := strings.Split(l, ":")
			if headerLine[1][0] == ' ' {
				headerLine[1] = headerLine[1][1:]
			}
			r.setHeader(headerLine[0], headerLine[1])
		}
	}

	return allHeadersRead
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
