package jsonrpc

const (
	TestRpcUrl = "http://domain"
)

type BaseTestRpcSvr struct {
}

func (*BaseTestRpcSvr) BaseUrl() string {
	return TestRpcUrl
}
