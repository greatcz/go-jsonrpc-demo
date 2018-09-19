package main

import (
	"fmt"

	"go-jsonrpc-demo/rpc"
)

func main() {
	reqArgs := new(rpc.RpcReqArgs)
	reqArgs.Method = "one"
	reqArgs.Params = []interface{}{4071}

	testSvr := new(rpc.TestRpcSvr).New()
	result, err := rpc.Call(testSvr, reqArgs)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
