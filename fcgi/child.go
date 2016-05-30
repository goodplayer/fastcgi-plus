package fcgi

import (
	"bufio"
	"net"
	"net/http"
)

type child struct {
	conn       net.Conn
	recordChan chan *request
	r          *bufio.Reader
}

func startChildHandleLoop(conn net.Conn, handler http.Handler, fcgis *fcgiServer) error {
	c := getChild()
	c.conn = conn
	c.r = bufio.NewReader(conn)

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
		err = packetDispatching(req)
		if err != nil {
			logError("packet dispatching error. exit inbound loop.", err)
			break
		}
	}
}

func (this *child) outboundProcessor() {
	//TODO
}

func packetDispatching(req request) error {
	//TODO
	if req.Header.getRequestId() == 0 {

	}
	return nil
}
