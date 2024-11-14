package main

import (
	"io"
	"testing"
)

func TestResponseWriteHeader(t *testing.T) {
	req := InitRequest()
	req.ParseRequestMessage([]byte(RAW_REQUEST_WITHOUT_BODY))
	rw := newResponseWriter(nil, io.Discard, req)

	rw.Header().Set("Connection", "close")
	rw.Header().Set("Content-Type", "text/plain")
	rw.Header().Set("Content-Length", "0")

	tests := []expectedHeader{
		{
			name:  "Connection",
			value: "close",
		},
		{
			name:  "Content-Type",
			value: "text/plain",
		},
		{
			name:  "Content-Length",
			value: "0",
		},
	}

	testResponseHeader(t, tests, rw)

	rw.flushResponse()
}

func testResponseHeader(
	t *testing.T,
	resT []expectedHeader,
	rw ResponseWriter,
) {
	t.Helper()

	for _, h := range resT {
		headerValue := rw.Header().Get(h.name)
		if headerValue != h.value {
			t.Fatalf("value of header='%s' did not match. want='%s' got='%s'",
				h.name, h.value, headerValue)
		}
	}
}
