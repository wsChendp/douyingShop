.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrity
gen-demo-thrity:
	@cd demo/demo_thrity && cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_thrity --service demo_thrity --idl ../../idl/echo.thrity

.PHONY: demo-link-fix
demo-link-fix:
	cd demo/demo_proto && golangci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend
gen-frontend:
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../id
