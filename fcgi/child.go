package fcgi

import (
	"bufio"
	"net"
	"net/http"
)

type child struct {
	conn       net.Conn
	recordChan chan reqWriter
	r          *bufio.Reader
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
	defer this.reset()
	defer this.release()
	for {
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
		bizErr, ioErr := this.packetDispatching(req)
		if bizErr != nil {
			logError("packet dispatching occurs biz error. exit inbound loop.", err)
			this.release()
			//readToEOF(r) // perhaps we don't need read to EOF
			break
		} else if ioErr != nil {
			logError("packet dispatching occurs io error. exit inbound loop.", err)
			break
		}
	}
}

func (this *child) outboundProcessor() {
	//TODO
}

func (this *child) packetDispatching(req request) (error, error) {
	//TODO
	if req.Header.getRequestId() == 0 {
		t := req.Header.Type
		ok := checkIsManagingRequest(t)
		if !ok {
			logError("requestId 0 with non-managing request")
		}
		// we don't support managing request now, send unknow type packet
		un := unknownTypeMessage{}
		un.setType(req.Header.Type)
		this.recordChan <- un
	}
	return nil, nil
}
