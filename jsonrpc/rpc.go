package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	Version   = "2.0"
	DefaultID = 1
)

type RpcSvr interface {
	BaseUrl() string
	ServiceName() string
}

type RpcReqArgs struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      uint8         `json:"id"`
}

type RpcResp struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      uint8       `json:"id"`
	Error   interface{} `json:"error"`
}

func Call(rpc RpcSvr, args *RpcReqArgs) (result interface{}, err error) {
	url := fmt.Sprintf("%s/%s", rpc.BaseUrl(), rpc.ServiceName())
	args.JsonRpc = Version
	args.ID = DefaultID

	data, err := json.Marshal(args)
	if nil != err {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if nil != err {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	// todo statusCode check
	body, _ := ioutil.ReadAll(resp.Body)
	rpcResp := RpcResp{}
	err = json.Unmarshal(body, &rpcResp)
	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	return rpcResp.Result, nil
}
