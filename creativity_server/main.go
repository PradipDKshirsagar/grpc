package main

import (
	"context"
	"creativitygrpc/creativity"
	"creativitygrpc/creativity_server/db"
	"creativitygrpc/creativity_server/user"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PONG")
}

var router = mux.NewRouter()

func init() {
	db.DBConnetion()
}

type server struct{}

func (s *server) GetUser(ctx context.Context, in *creativity.HelloRequest) (*creativity.HelloReply, error) {
	log.Printf("Received: %v", in.Id)
	var us user.User
	us, err := user.ReadService(in.Id)
	msg := fmt.Sprintf("%v", us)
	return &creativity.HelloReply{Message: msg}, err
}

func main() {
	defer db.GetDB().Close()

	router.HandleFunc("/ping", pingHandler)

	http.Handle("/", router)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	creativity.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
