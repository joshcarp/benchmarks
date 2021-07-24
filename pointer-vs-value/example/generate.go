//go:generate bash -c "docker run --rm -v $(pwd):/example:rw joshcarp/protoc -I./example/ --go-grpc_out=paths=source_relative:example --go_out=paths=source_relative:example example.proto"

package example

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type ServerImplValue struct {
	ExampleServiceServer
}

func (s ServerImplValue) GetExample(ctx context.Context, in *Example) (*Example, error) {
	return &Example{Name: "Hello " + in.Name, Whatever: "Hello" + in.Whatever}, nil
}

func ServeRand(srv ExampleServiceServer) (int, net.Listener, error) {
	server := grpc.NewServer()
	RegisterExampleServiceServer(server, srv)
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, nil, err
	}
	go server.Serve(ln)
	return ln.Addr().(*net.TCPAddr).Port, ln, nil
}

type ServerImplPointer struct{
	ExampleServiceServer
}

func (s *ServerImplPointer) GetExample(ctx context.Context, in *Example) (*Example, error) {
	return &Example{Name: "Hello " + in.Name, Whatever: "Hello" + in.Whatever}, nil
}
