package main

func closeConnectionHandler(w ResponseWriter, r *Request) {
	w.Header().Add("Connection", "close")
}

func notFoundHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusNotFound)
	closeConnectionHandler(w, r)
}

func methodNotAllowed(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusMethodNotAllowed)
	w.Header().Add("Allow", "GET")
	closeConnectionHandler(w, r)
}

func internalServerErrorHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusInternalServerError)
	closeConnectionHandler(w, r)
}
