package main

import (
	"flag"
	"log"

	"github.com/kki7823/kube-autocomplete-server/server"
)

func main() {
	runServer()
}

func runServer() {
	var serverPort string
	flag.StringVar(&serverPort, "server-port", "8081", "server port")
	o := &server.Option{
		Port: serverPort,
	}

	s := server.NewServer(o)
	err := s.Run()
	if err != nil {
		log.Fatal("unable to start erver", err)
	}
}
