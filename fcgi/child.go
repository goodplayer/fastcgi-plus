package fcgi

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

type child struct {
	conn       net.Conn
	recordChan chan reqWriter
	r          *bufio.Reader
	requests   map[uint16]*statefulRequest
}

func startChildHandleLoop(conn net.Conn, handler http.Handler, fcgis *fcgiServer) error {
	c := getChild()
	c.conn = conn
	c.r = bufio.NewReader(conn)
	c.recordChan = make(chan reqWriter, 64)

	go c.childHandleProcessor()
	go c.outboundProcessor()

	//TODO
	return nil
}

func (this *child) release() {
	this.conn.Close()
}

func (this *child) reset() {
	//TODO
}

func (this *child) childHandleProcessor() {
	r := this.r
	loop := true
	requests := this.requests
	defer this.reset()
	defer this.release()
	for loop {
		header := requestHeader{}
		close, err := header.read(r)
		if close {
			break
		}
		if err != nil {
			logError("read header error.", err)
			break
		}
		req := request{
			Header: header,
		}
		cl := int(header.getContentLength()) & 0xFFFF
		var bi *BufItem
		if header.getContentLength() > 0 {
			bi = getBufItem()
			b := bi.GetBuffer()[:cl]
			n, err := io.ReadFull(r, b[:cl])
			if err != nil {
				if n < cl {
					logError("packet dispatching occurs io error. exit inbound loop.", err)
					bi.Release()
					break
				} else {
					loop = false // exit loop, but process last request
				}
			}
			req.ContentData = b[:cl]
		}
		bizErr := this.packetDispatching(req, requests)
		bi.Release() // user should use retain/release for custom reason
		if bizErr != nil {
			logError("packet dispatching occurs biz error. exit inbound loop.", err)
			//readToEOF(r) // perhaps we don't need read to EOF
			break
		}
	}
}

func (this *child) outboundProcessor() {
	//TODO
}

func (this *child) packetDispatching(req request, reqMap map[uint16]*statefulRequest) error {
	//TODO
	if req.Header.getRequestId() == 0 {
		t := req.Header.Type
		ok := checkIsManagingRequest(t)
		if !ok {
			logError("requestId 0 with non-managing request")
		}
		// we don't support managing request now, send unknow type packet
		un := unknown_type_packet
		un.setType(req.Header.Type)
		this.recordChan <- un
	} else {
		//TODO
	}
	return nil
}
