package fcgi

import (
	"net"
	"net/http"
)

type child struct {
	conn net.Conn
}

func startChildHandleLoop(conn net.Conn, handler http.Handler, fcgis *fcgiServer) error {
	c := getChild()
	c.conn = conn

	go c.childHandleProcessor()

	//TODO
	return nil
}

func (this *child) reset() {
	//TODO
}

func (this *child) childHandleProcessor() {
	//TODO
}

func packetDispatching(req *request) {
	//TODO
}
