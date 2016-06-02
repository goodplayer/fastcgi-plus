package fcgi

import (
	"io"
	"reflect"
	"unsafe"
)

type reqWriter interface {
	write(io.Writer) (int, error)
}

type request struct {
	Header      requestHeader
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
	return uint16(uint16(this.RequestIdMSB)<<8) | uint16(this.RequestIdLSB)
}

func (this *requestHeader) setContentLength(length uint16) {
	this.ContentLengthLSB = byte(length)
	this.ContentLengthMSB = byte(length >> 8)
}

func (this *requestHeader) getContentLength() uint16 {
	return uint16(uint16(this.ContentLengthMSB)<<8) | uint16(this.ContentLengthLSB)
}

func (this *requestHeader) read(r io.Reader) (bool, error) {
	h := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(this)),
		Len:  8,
		Cap:  8,
	}
	n, err := io.ReadFull(r, *(*[]byte)(unsafe.Pointer(&h)))
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

type _unknownTypeMessage [16]byte

func (this *_unknownTypeMessage) setType(t byte) {
	this[8] = t
}

func (this *_unknownTypeMessage) toBytes() []byte {
	return (*this)[:]
}

func (this _unknownTypeMessage) write(w io.Writer) (int, error) {
	return w.Write(this[:])
}

type _endRequestMessage [16]byte

func (this *_endRequestMessage) setRequestId(reqId uint16) {
	(*this)[2] = byte(reqId >> 8)
	(*this)[3] = byte(reqId)
}

func (this *_endRequestMessage) setAppStatus(appStatus int32) {
	(*this)[8] = byte(appStatus >> 24)
	(*this)[9] = byte(appStatus >> 16)
	(*this)[10] = byte(appStatus >> 8)
	(*this)[11] = byte(appStatus)
}

func (this *_endRequestMessage) setProtocolStatus(protocolStatus byte) {
	(*this)[12] = byte(protocolStatus)
}

func (this _endRequestMessage) write(w io.Writer) (int, error) {
	return w.Write(this[:])
}

const (
	_STATEFUL_REQUEST_STATE_READING_PARAM = 1
	_STATEFUL_REQUEST_STATE_READING_STDIN = 2
	_STATEFUL_REQUEST_STATE_READING_DATA  = 3
	_STATEFUL_REQUEST_STATE_READING_DONE  = 4
)

type statefulRequest struct {
	//TODO
	reqId uint16
	state byte
}

func (this *statefulRequest) reset() {
	//TODO
}
