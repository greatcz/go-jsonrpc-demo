package rpc

type BaseTestRpcSvr struct {
}

func (btrs *BaseTestRpcSvr) BaseUrl() string {
	return "http://domain/"
}
