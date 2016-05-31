package fcgi

import "testing"

func TestRequestHeaderGetterSetter(t *testing.T) {
	rh := requestHeader{}
	if rh.getRequestId() != 0 {
		t.Fatal("reqeuset header is not 0.")
	}
	rh.setRequestId(12345)
	if rh.getRequestId() != 12345 {
		t.Fatal("request header is not 12345. actual:", rh.getRequestId())
	}
}

func TestUnknownTypeMessage(t *testing.T) {
	ty := unknown_type_packet
	t.Log(ty.toBytes())
	ty.setType(1)
	t.Log(ty.toBytes())
	t.Log(unknown_type_packet.toBytes())
}
