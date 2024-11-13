package main

import (
	"testing"
)

type expectedHeader struct {
	name  string
	value string
}

type requestTest struct {
	expectedStartLine   string
	expectedMethodName  string
	expectedPath        string
	expectedHTTPVersion string
	expectedHeaders     []expectedHeader
}

const RAW_REQUEST = "GET /api/auth HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n"

func TestParseFullRequest(t *testing.T) {
	req := InitRequest()
	req.ParseRequestMessage([]byte(RAW_REQUEST))

	test := requestTest{
		expectedStartLine:   "GET /api/auth HTTP/1.1\r\n",
		expectedMethodName:  "GET",
		expectedPath:        "/api/auth",
		expectedHTTPVersion: "HTTP/1.1",
		expectedHeaders: []expectedHeader{
			{
				name:  "Host",
				value: "example.com",
			},
			{
				name:  "Connection",
				value: "close",
			},
		},
	}

	testRequestStartLine(t, test, req)
	testRequestHeader(t, test, req)
}

func TestParseStreamRequest(t *testing.T) {
	req := InitRequest()
	for _, c := range []byte(RAW_REQUEST) {
		req.ParseRequestMessage([]byte{c})
	}

	test := requestTest{
		expectedStartLine:   "GET /api/auth HTTP/1.1\r\n",
		expectedMethodName:  "GET",
		expectedPath:        "/api/auth",
		expectedHTTPVersion: "HTTP/1.1",
		expectedHeaders: []expectedHeader{
			{
				name:  "Host",
				value: "example.com",
			},
			{
				name:  "Connection",
				value: "close",
			},
		},
	}

	testRequestStartLine(t, test, req)
	testRequestHeader(t, test, req)
}

func testRequestStartLine(
	t *testing.T,
	reqT requestTest,
	req *Request,
) {
	t.Helper()

	if reqT.expectedStartLine != req.startLine() {
		t.Fatalf("start line did not match. want='%s' got='%s'",
			reqT.expectedStartLine, req.startLine())
	}

	if reqT.expectedMethodName != req.Method {
		t.Fatalf("method name did not match. want='%s' got='%s'",
			reqT.expectedMethodName, req.Method)
	}

	if reqT.expectedPath != req.Path {
		t.Fatalf("path did not match. want='%s' got='%s'",
			reqT.expectedPath, req.Path)
	}

	if reqT.expectedHTTPVersion != req.V {
		t.Fatalf("HTTP version did not match. want='%s' got='%s'",
			reqT.expectedHTTPVersion, req.V)
	}
}

func testRequestHeader(
	t *testing.T,
	reqT requestTest,
	req *Request,
) {
	t.Helper()

	for _, h := range reqT.expectedHeaders {
		headerValue, ok := req.Headers[h.name]
		if !ok {
			t.Fatalf("'%s' header is not present in request", h.name)
		}
		if headerValue != h.value {
			t.Fatalf("value of header='%s' did not match. want='%s' got='%s'",
				h.name, h.value, headerValue)
		}
	}
}
