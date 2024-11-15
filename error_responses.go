package main

func closeConnectionHandler(w ResponseWriter) {
	w.Header().Add("Connection", "close")
}

func notFoundHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusNotFound)
	w.Write([]byte("Not Found"))
	closeConnectionHandler(w)
}

func methodNotAllowed(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusMethodNotAllowed)
	w.Write([]byte("Not Allowed"))
	w.Header().Add("Allow", "GET")
	closeConnectionHandler(w)
}

func internalServerErrorHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
	closeConnectionHandler(w)
}
