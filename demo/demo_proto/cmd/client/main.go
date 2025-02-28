package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		fmt.Println(11)
		log.Fatal(err)
	}
	fmt.Println("连接成功")
	c, err := echo.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		fmt.Println(22)
		log.Fatal(err)
	}
	fmt.Println("创建服务类成功")

	res, err := c.Echo(context.Background(), &pbapi.Request{Message: "hello"})
	if err != nil {
		fmt.Println(33)
		log.Println(err)
	}
	fmt.Println("%v", res)
}
