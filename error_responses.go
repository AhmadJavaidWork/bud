package main

func closeConnectionHandler(w ResponseWriter) {
	w.Header().Add("Connection", "close")
}

func notFoundHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusNotFound)
	closeConnectionHandler(w)
}

func methodNotAllowed(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusMethodNotAllowed)
	w.Header().Add("Allow", "GET")
	closeConnectionHandler(w)
}

func internalServerErrorHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusInternalServerError)
	closeConnectionHandler(w)
}
