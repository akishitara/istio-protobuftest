package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../config"

	pb "../protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	setupConfig = config.ConfigLoad("/config/config.toml")
	serverAddr  = setupConfig.Server.Address + ":" + strconv.Itoa(setupConfig.Server.Port)
	listenPort  = ":" + strconv.Itoa(setupConfig.Client.Port)
)

func main() {
	log.Print(listenPort)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(listenPort, nil)
	if err != nil {
		log.Print(err)
	}
	log.Print(listenPort)
}

func handler(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, getGrpcBooks())
}

func getGrpcBooks() string {
	flag.Parse()
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer conn.Close()
	c := pb.NewBooksServerClient(conn)
	resp, err := c.SearchTitle(context.Background(),
		&pb.BooksRequest{Name: "bbb", Page: 1})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	return resp.Title
}
