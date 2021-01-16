package main

import (
	proto "cly/protos"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()
	greeterClient := proto.NewGreeterService("Greeter", service.Client())
	for i := 0; i < 20000; i++ {

		rsp, err := greeterClient.Hello(context.TODO(), &proto.Request{Name: "jackson"})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(rsp.Greeting)
	}
}
