package fcgi

func checkIsManagingRequest(t byte) bool {
	switch t {
	case _FCGI_GET_VALUES:
		return true
	default:
		return false
	}
}
