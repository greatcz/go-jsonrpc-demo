package rpc

type TestRpcSvr struct {
	*BaseTestRpcSvr
}

func (tsvr *TestRpcSvr) BaseUrl() string {
	return tsvr.BaseTestRpcSvr.BaseUrl()
}

func (tsvr *TestRpcSvr) ServiceName() string {
	return "rpc/testSvr"
}

func (tsvr *TestRpcSvr) New() *TestRpcSvr {
	tsvr.BaseTestRpcSvr = new(BaseTestRpcSvr)
	return tsvr
}
