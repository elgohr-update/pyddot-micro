package main

import (
	"example-service/handler"
	example "example-service/proto"
	"fmt"

	"github.com/micro/micro/v2/service"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.example"),
		service.Version("latest"),
	)

	// Initialise service
	srv.Init()

	// Register Handler
	example.RegisterExampleHandler(srv.Server(), new(handler.Example))

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
