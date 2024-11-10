package main

const (
	StatusContinue           = 100
	StatusSwitchingProtocols = 101

	StatusOK                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206

	StatusMultipleChoices   = 300
	StatusMovedPermanently  = 301
	StatusFound             = 302
	StatusSeeOther          = 303
	StatusNotModified       = 304
	StatusUseProxy          = 305
	StatusTemporaryRedirect = 307

	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusPaymentRequired              = 402
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusMethodNotAllowed             = 405
	StatusNotAcceptable                = 406
	StatusProxyAuthRequired            = 407
	StatusRequestTimeout               = 408
	StatusConflict                     = 409
	StatusGone                         = 410
	StatusLengthRequired               = 411
	StatusPreconditionFailed           = 412
	StatusRequestEntityTooLarge        = 413
	StatusURITooLong                   = 414
	StatusUnsupportedMediaType         = 415
	StatusRequestedRangeNotSatisfiable = 416
	StatusExpectationFailed            = 417

	StatusInternalServerError     = 500
	StatusNotImplemented          = 501
	StatusBadGateway              = 502
	StatusServiceUnavailable      = 503
	StatusGatewayTimeout          = 504
	StatusHTTPVersionNotSupported = 505
)

func StatusText(code int) string {
	switch code {
	case StatusContinue:
		return "Continue"
	case StatusSwitchingProtocols:
		return "Switching Protocols"
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusAccepted:
		return "Accepted"
	case StatusNonAuthoritativeInfo:
		return "Non-Authoritative Informantion"
	case StatusNoContent:
		return "No Content"
	case StatusResetContent:
		return "Reset Content"
	case StatusPartialContent:
		return "Partial Content"
	case StatusMultipleChoices:
		return "Multiple Choices"
	case StatusMovedPermanently:
		return "Moved Permanently"
	case StatusFound:
		return "Found"
	case StatusSeeOther:
		return "See Other"
	case StatusNotModified:
		return "Not Modified"
	case StatusUseProxy:
		return "Use Proxy"
	case StatusTemporaryRedirect:
		return "Temporary Redirect"
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"
	case StatusNotAcceptable:
		return "Not Acceptable"
	case StatusProxyAuthRequired:
		return "Proxy Authentication Required"
	case StatusRequestTimeout:
		return "Request Timeout"
	case StatusConflict:
		return "Conflict"
	case StatusGone:
		return "Gone"
	case StatusLengthRequired:
		return "Length Required"
	case StatusPreconditionFailed:
		return "Precondition Failed"
	case StatusRequestEntityTooLarge:
		return "Request Entity Too Large"
	case StatusURITooLong:
		return "Request URI Too Long"
	case StatusUnsupportedMediaType:
		return "Unsupported Media Type"
	case StatusRequestedRangeNotSatisfiable:
		return "Request Range Not Satisfiable"
	case StatusExpectationFailed:
		return "Expectation Failed"
	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusNotImplemented:
		return "Not Implemented"
	case StatusBadGateway:
		return "Bad Gateway"
	case StatusServiceUnavailable:
		return "Service Unavailable"
	case StatusGatewayTimeout:
		return "Gateway Timeout"
	case StatusHTTPVersionNotSupported:
		return "HTTP Version Not Supported"
	default:
		return ""
	}
}
