package fcgi

import (
	"net"
	"net/http"
)

type fcgiServer struct {
	option *FcgiServerOption
}

type FcgiServerOption struct {
}

func NewFcgiServer(option *FcgiServerOption) *fcgiServer {
	fcgis := new(fcgiServer)
	fcgis.option = option
	return fcgis
}

func (this *fcgiServer) Serve(l net.Listener, handler http.Handler) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		err = startChildHandleLoop(conn, handler, this)
		if err != nil {
			return err
		}
	}
	return nil
}
