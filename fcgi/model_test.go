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
