package main

import "testing"

func tempHandler(rw ResponseWriter, r *Request)         {}
func tempLoginHandler(rw ResponseWriter, r *Request)    {}
func tempRegisterHandler(rw ResponseWriter, r *Request) {}

func TestRouter(t *testing.T) {
	router := NewRouter()
	router.addRoute("/api/auth", "GET", tempHandler)
	router.addRoute("/api/auth/login", "POST", tempLoginHandler)
	router.addRoute("/api/auth/register", "POST", tempRegisterHandler)

	tests := []struct {
		path            string
		method          string
		expectedHandler string
	}{
		{
			path:            "/api/auth",
			method:          "GET",
			expectedHandler: "tempHandler",
		},
		{
			path:            "/api/auth/login",
			method:          "POST",
			expectedHandler: "tempLoginHandler",
		},
		{
			path:            "/api/auth/register",
			method:          "POST",
			expectedHandler: "tempRegisterHandler",
		},
	}

	for _, tt := range tests {
		handler := router.getHandler(tt.path, tt.method)
		if handler == nil {
			t.Fatalf("expected handler does not match. want=%s, got=nil", tt.expectedHandler)
		}
	}
}
