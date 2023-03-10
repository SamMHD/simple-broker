<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# receiver

```go
import "github.com/SamMHD/simple-broker/receiver"
```

package receiver implements the HTTP gateway for the Broker. it receives messages over HTTP and forwards them to the Broker using gRPC. HTTP requests are served by gin.

## Index

- [type Server](<#type-server>)
  - [func NewServer(config util.Config) (*Server, error)](<#func-newserver>)
  - [func (server *Server) Start() error](<#func-server-start>)


## type [Server](<https://github.com/SamMHD/simple-broker/blob/main/receiver/server.go#L16-L20>)

Server serves HTTP requests for the receiver

```go
type Server struct {
    // contains filtered or unexported fields
}
```

### func [NewServer](<https://github.com/SamMHD/simple-broker/blob/main/receiver/server.go#L23>)

```go
func NewServer(config util.Config) (*Server, error)
```

NewServer creates a new HTTP server and set up routing.

### func \(\*Server\) [Start](<https://github.com/SamMHD/simple-broker/blob/main/receiver/server.go#L83>)

```go
func (server *Server) Start() error
```

Start runs the Gin HTTP server it will return an error if the server fails to start, otherwise it will block the thread and wait for the server to stop. in case of peacefull stop, it will return nil.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
