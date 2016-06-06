package innerapi

type ChildProcessor interface {
}

type NvPair interface {
	GetKeyValue() ([]byte, []byte)
}
