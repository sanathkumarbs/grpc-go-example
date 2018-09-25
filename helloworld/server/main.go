// Generate Proto
// go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Packages
// https://tour.golang.org/basics/1

package main

// Imports
// https://tour.golang.org/basics/2

import (
	"log"
	"net"

	// Full paths of import from $GOPATH/src directory
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// In python, this is similar to `import lib as pb`
	pb "github.com/sanathkumarbs/grpc-go-example/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

// Constants
// https://tour.golang.org/basics/15

const (
	port = ":50051"
)

// structs are similar to a "class"
// https://tour.golang.org/moretypes/2

// Defining a server to implement helloworld.server
type server struct{}

// SayHello is a method. Similar to a method in a class.

// This method operates on server s.
// Consider `server` as the classname, `s` as its object inside the method
// In goland, the (s *server) are called recievers

// To truly have s as an object, you would need a reference
// for the `server`, so you send it as a reference `*server`

// To just pass by value use func (s server) SayHello() {...}

// Method args and returns
// func (<recievers>) methodname (args) (returntypes) {...}
// https://tour.golang.org/basics/6

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Primer on pointers
	// http://piotrzurek.net/2013/09/20/pointers-in-go.html
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

// main is a function (not a method!).
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
