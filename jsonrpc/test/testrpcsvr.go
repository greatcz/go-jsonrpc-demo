package test

import "go-jsonrpc-demo/jsonrpc"

const (
	ServiceNameTest = "rpc/test"
)

const (
	MethodOne = "one"
)

type RpcSvr struct {
	*jsonrpc.BaseTestRpcSvr
}

func (*RpcSvr) ServiceName() string {
	return ServiceNameTest
}
