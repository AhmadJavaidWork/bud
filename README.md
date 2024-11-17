# bud

http implementation in go

`NewServer(addr string) *Server`

Creates a new `Server` and returns a pointer to it

`Type Server`

```
Addr            string  // tcp4 address to open the connection on e.g 127.0.0.1:8080
router          *Router // handles routing
openConnections int     // counts open connections at any moment
```

- `ListenAndServe() error`

  Opens a tcp listener on the provided server.Addr and start accepting new connections.

- `AddHandler(pattern string, handler Handler)`

  Registers a new handler for the pattern string.

- `handleConnection() error`

  Handles a connection. It reads HTTP messages, parses them into `Request` and serves them accordingly. It also closes connections if the HTTP message is malformed or the message form is not supported.
