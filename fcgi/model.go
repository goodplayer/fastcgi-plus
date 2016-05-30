package fcgi

import (
	"io"
	"reflect"
	"unsafe"
)

type request struct {
	Header      *requestHeader
	ContentData []byte
	PaddingData []byte
}

type requestHeader struct {
	Version          byte
	Type             byte
	RequestIdMSB     byte
	RequestIdLSB     byte
	ContentLengthMSB byte
	ContentLengthLSB byte
	PaddingLength    byte
	Reserved         byte
}

func (this *requestHeader) reset() {
	*(*int64)(unsafe.Pointer(this)) = 0
}

func (this *requestHeader) setRequestId(id uint16) {
	this.RequestIdLSB = byte(id)
	this.RequestIdMSB = byte(id >> 8)
}

func (this *requestHeader) getRequestId() uint16 {
	return uint16(this.RequestIdMSB<<8) | uint16(this.RequestIdLSB)
}

func (this *requestHeader) setContentLength(length uint16) {
	this.ContentLengthLSB = byte(length)
	this.ContentLengthMSB = byte(length >> 8)
}

func (this *requestHeader) getContentLength() uint16 {
	return uint16(this.ContentLengthMSB<<8) | uint16(this.ContentLengthLSB)
}

func (this *requestHeader) read(r io.Reader) (bool, error) {
	h := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(this)),
		Len:  8,
		Cap:  8,
	}
	n, err := io.ReadFull(r, []byte(*(*[]byte)(unsafe.Pointer(&h))))
	if n == 0 && err == io.EOF {
		return true, nil
	}
	return false, err
}

type nameValuePair11 struct {
	NameLength  int8
	ValueLength int8
	NameData    []byte
	ValueData   []byte
}

type nameValuePair14 struct {
	NameLength    int8
	ValueLengthB3 int8
	ValueLengthB2 byte
	ValueLengthB1 byte
	ValueLengthB0 byte
	NameData      []byte
	ValueData     []byte
}

type nameValuePair41 struct {
	NameLengthB3 int8
	NameLengthB2 byte
	NameLengthB1 byte
	NameLengthB0 byte
	ValueLength  int8
	NameData     []byte
	ValueData    []byte
}

type nameValuePair44 struct {
	NameLengthB3  int8
	NameLengthB2  byte
	NameLengthB1  byte
	NameLengthB0  byte
	ValueLengthB3 int8
	ValueLengthB2 byte
	ValueLengthB1 byte
	ValueLengthB0 byte
	NameData      []byte
	ValueData     []byte
}

type unknownTypeBody struct {
	Type     byte
	Reserved [7]byte
}

type beginRequestBody struct {
	RoleMSB  byte
	RoleLSB  byte
	Flags    byte
	Reserved [5]byte
}

type endRequestBody struct {
	AppStatusB3    byte
	AppStatusB2    byte
	AppStatusB1    byte
	AppStatusB0    byte
	ProtocolStatus byte
	Reserved       [3]byte
}
