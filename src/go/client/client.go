package main

import (
	"flag"
	"fmt"

	"github.com/gonzaloserrano/grpc-go/src/go/clock"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const address = "localhost:8888"

func main() {
	user := flag.String("user", "unknown", "user id to init data, rpc calls metadata etc")
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := clock.NewClockClient(conn)

	resp, err := c.WhatTimeIsIt(context.Background(), &clock.ClockRequest{*user})
	if err != nil {
		panic(err)
	}
	for {
		msg, err := resp.Recv()
		if err != nil {
			println("Disconnected")
			return
		}
		fmt.Printf("Received `%s`\n", msg.Message)
	}
}
