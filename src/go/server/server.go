package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gonzaloserrano/grpc-go/src/go/clock"

	"google.golang.org/grpc"
)

const port = ":8888"

type server struct{}

// Get implements the clock.ClockServer interface
func (s *server) WhatTimeIsIt(req *clock.ClockRequest, stream clock.Clock_WhatTimeIsItServer) error {
	fmt.Printf("Hello %s\n", req.Name)
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			println("tick")
			resp := &clock.ClockResponse{time.Now().String()}
			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	clock.RegisterClockServer(s, &server{})
	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	println("serving at port", port)
	for {
	}
}
