package rpc

type TestRpcSvr struct {
	*BaseTestRpcSvr
}

func (psvr *TestRpcSvr) BaseUrl() string {
	return psvr.BaseTestRpcSvr.BaseUrl()
}

func (psvr *TestRpcSvr) ServiceName() string {
	return "rpc/testSvr"
}

func (psvr *TestRpcSvr) New() *TestRpcSvr {
	psvr.BaseTestRpcSvr = new(BaseTestRpcSvr)
	return psvr
}
