package main

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	HEAD   = "HEAD"
	PATCH  = "PATCH"
	TRACE  = "TRACE"
)

func MethodsLookUp(name string) string {
	switch name {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	case "HEAD":
		return HEAD
	case "PATCH":
		return PATCH
	case "TRACE":
		return TRACE
	default:
		return ""
	}
}
