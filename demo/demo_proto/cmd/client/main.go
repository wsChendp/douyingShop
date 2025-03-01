package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	c, err := echo.NewClient("demo_proto", client.WithResolver(r),
		//client.WithTransportProtocol(transport.GRPC), // 不使用 gRPC
		client.WithShortConnection(), // 使用短链接
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)

	if err != nil {
		panic(err)
	}

	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")

	res, err := c.Echo(ctx, &pbapi.Request{Message: "error"})
	var bizErr *kerrors.GRPCBizStatusError

	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Println("%#v", bizErr)
		}
		klog.Fatal(err)
	}
	fmt.Println("%v", res)
}
