package innerapi

type ChildProcessor interface {
	ProcessParam(ParamContainer) (interface{}, error)
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
