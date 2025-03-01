package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"time"
)

func Middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// do something before calling next
		begin := time.Now()
		err = next(ctx, req, resp)
		end := time.Now()
		// do something after calling next
		fmt.Println("middleware", end.Sub(begin), "begin", begin, "end", end)
		return err
	}
}
