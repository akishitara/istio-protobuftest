package main

import (
	"log"
	"net"
	"strconv"

	"../config"

	pb "../protocol"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var setupConfig = config.ConfigLoad("/config/config.toml")
var listenPort = ":" + strconv.Itoa(setupConfig.Server.Port)
var redisAddress string = setupConfig.Redis.Address + ":" + strconv.Itoa(setupConfig.Redis.Port)
var (
	booksTitle map[string]string = map[string]string{
		"aaa": "version 1",
		"bbb": "version 2",
	}
)

type server struct{}

func (s *server) SearchTitle(ctx context.Context, in *pb.BooksRequest) (*pb.BooksReply, error) {
	log.Printf("New Request: %v \n", in.String())
	log.Printf("Reply: %v \n", booksTitle[in.GetName()])
	return &pb.BooksReply{Title: "Title: " + booksTitle[in.GetName()] + "!" + "redis:" + redisGet(redisAddress)}, nil
}

func redisGet(redisAddress string) string {
	c, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		log.Print(err)
	}
	defer c.Close()
	s, err := redis.String(c.Do("GET", "1"))
	if err != nil {
		log.Print(err)
		return "null"
	}
	log.Print(s)
	return s
}

func main() {
	log.Print(listenPort)
	lis, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Print(err)
	}
	s := grpc.NewServer()
	pb.RegisterBooksServerServer(s, &server{})
	s.Serve(lis)
}
