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

const RAW_REQUEST_WITHOUT_BODY = "GET /api/auth HTTP/1.1\r\nMax-Forwards: 0\r\nUser-Agent: SomeAgent/29.92.9\r\nAccept: */*\r\nAuthorization: 9cvklsjdflakd1762a-c741lj-4ljf8ljb-94iuouoi1a-fpououb2c0a9abe4d\r\nHost: example.com\r\nAccept-Encoding: gzip, deflate, br\r\nConnection: keep-alive\r\n\r\n"
const RAW_REQUEST_WITH_BODY = "GET /api/auth HTTP/1.1\r\nContent-Length: 20\r\nMax-Forwards: 0\r\nUser-Agent: SomeAgent/29.92.9\r\nAccept: */*\r\nAuthorization: 9cvklsjdflakd1762a-c741lj-4ljf8ljb-94iuouoi1a-fpououb2c0a9abe4d\r\nHost: example.com\r\nAccept-Encoding: gzip, deflate, br\r\nConnection: keep-alive\r\n\r\n{\n    \"id\": \"1000\"\n}"

var test = requestTest{
	expectedStartLine:   "GET /api/auth HTTP/1.1\r\n",
	expectedMethodName:  "GET",
	expectedPath:        "/api/auth",
	expectedHTTPVersion: "HTTP/1.1",
	expectedHeaders: []expectedHeader{
		{
			name:  "Max-Forwards",
			value: "0",
		},
		{
			name:  "User-Agent",
			value: "SomeAgent/29.92.9",
		},
		{
			name:  "Accept",
			value: "*/*",
		},
		{
			name:  "Authorization",
			value: "9cvklsjdflakd1762a-c741lj-4ljf8ljb-94iuouoi1a-fpououb2c0a9abe4d",
		},
		{
			name:  "Host",
			value: "example.com",
		},
		{
			name:  "Accept-Encoding",
			value: "gzip, deflate, br",
		},
		{
			name:  "Connection",
			value: "keep-alive",
		},
	},
}

func TestParseFullRequestWithoutBody(t *testing.T) {
	req := InitRequest()
	req.ParseRequestMessage([]byte(RAW_REQUEST_WITHOUT_BODY))

	testRequestStartLine(t, test, req)
	testRequestHeader(t, test, req)
}

func TestParseStreamRequestWithoutBody(t *testing.T) {
	for i := 1; i <= len(RAW_REQUEST_WITHOUT_BODY); i++ {
		req := InitRequest()
		for j := 0; j < len(RAW_REQUEST_WITHOUT_BODY); j += i {
			end := min(len(RAW_REQUEST_WITHOUT_BODY), j+i)
			req.ParseRequestMessage([]byte(RAW_REQUEST_WITHOUT_BODY)[j:end])
		}

		testRequestStartLine(t, test, req)
		testRequestHeader(t, test, req)
	}
}

func TestParseFullRequestWithBody(t *testing.T) {
	req := InitRequest()
	req.ParseRequestMessage([]byte(RAW_REQUEST_WITH_BODY))

	testRequestStartLine(t, test, req)
	testRequestHeader(t, test, req)
	testRequestBody(t, []byte("{\n    \"id\": \"1000\"\n}"), req)
}

func TestParseStreamRequestWithBody(t *testing.T) {
	for i := 1; i <= len(RAW_REQUEST_WITH_BODY); i++ {
		req := InitRequest()
		for j := 0; j < len(RAW_REQUEST_WITH_BODY); j += i {
			end := min(len(RAW_REQUEST_WITH_BODY), j+i)
			req.ParseRequestMessage([]byte(RAW_REQUEST_WITH_BODY)[j:end])
		}

		testRequestStartLine(t, test, req)
		testRequestHeader(t, test, req)
		testRequestBody(t, []byte("{\n    \"id\": \"1000\"\n}"), req)
	}
}

func testRequestStartLine(
	t *testing.T,
	reqT requestTest,
	req *Request,
) {
	t.Helper()

	if reqT.expectedStartLine != req.startLine() {
		t.Fatalf("start line does not match. want='%s' got='%s'",
			reqT.expectedStartLine, req.startLine())
	}

	if reqT.expectedMethodName != req.Method {
		t.Fatalf("method name does not match. want='%s' got='%s'",
			reqT.expectedMethodName, req.Method)
	}

	if reqT.expectedPath != req.Path {
		t.Fatalf("path does not match. want='%s' got='%s'",
			reqT.expectedPath, req.Path)
	}

	if reqT.expectedHTTPVersion != req.V {
		t.Fatalf("HTTP version does not match. want='%s' got='%s'",
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
			t.Fatalf("value of header='%s' does not match. want='%s' got='%s'",
				h.name, h.value, headerValue)
		}
	}
}

func testRequestBody(
	t *testing.T,
	expectedBody []byte,
	req *Request,
) {
	t.Helper()
	if len(expectedBody) != len(req.Body) {
		t.Fatalf("req body does not match. want='%q' got='%q'",
			expectedBody, req.Body)
	}

	for i := 0; i < len(expectedBody); i++ {
		if expectedBody[i] != req.Body[i] {
			t.Fatalf("req body does not match. want='%q' got='%q'",
				expectedBody, req.Body)
		}
	}
}
