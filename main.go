package main

func main() {
	server := &Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
