package main

import "bytes"

type Header map[string]string

func (h Header) Add(key, value string) {
	h[key] = value
}

func (h Header) Set(key, value string) {
	h[key] = value
}

func (h Header) Get(key string) string {
	if v, ok := h[key]; ok {
		return v
	}
	return ""
}

func (h Header) String() string {
	var out bytes.Buffer

	for k, v := range h {
		out.WriteString(k + ": " + v + "\r\n")
	}
	out.WriteString("\r\n")

	return out.String()
}
