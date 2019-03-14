package main

import (
	"fmt"
	"go-jsonrpc-demo/jsonrpc/test"

	"go-jsonrpc-demo/jsonrpc"
)

func main() {
	reqArgs := &jsonrpc.RpcReqArgs{
		Method: test.MethodOne,
		Params: []interface{}{4071},
	}

	testSvr := new(test.RpcSvr)
	result, err := jsonrpc.Call(testSvr, reqArgs)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
