package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro/v2"

	proto "cly/protos"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, request *proto.Request, response *proto.Response) error {
	response.Greeting = "hello" + request.Name
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("Greeter"),
	)
	//
	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
