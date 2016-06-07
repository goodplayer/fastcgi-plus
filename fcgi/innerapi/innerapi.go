package innerapi

type ChildProcessor interface {
	ProcessParam(ParamContainer)
}

type NvPair interface {
	GetKeyValue() ([]byte, []byte)
	GetKeyValueString() (string, string)
}

type ParamContainer interface {
	Set([]byte, []byte)
	Get([]byte) []byte
}
