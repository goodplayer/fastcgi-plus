package innerapi

import "io"

type ChildProcessor interface {
	ProcessParam(ParamContainer) (interface{}, error)
	CreateChildContainer() ChildContainer
	ServeRequest(req interface{}, body io.ReadCloser)
}

type NvPair interface {
	GetKeyValue() ([]byte, []byte)
	GetKeyValueString() (string, string)
}

type ParamContainer interface {
	Set([]byte, []byte)
	Get([]byte) []byte
	GetString(string) string
	GetNonFcgiParam() map[string][]byte
}

type ChildContainer struct {
	io.ReadCloser
	*io.PipeWriter
}
