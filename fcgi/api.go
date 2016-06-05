package fcgi

import (
	"net"
	"net/http"

	fcgihttp "github.com/goodplayer/fastcgi-plus/fcgi/http"
	"github.com/goodplayer/fastcgi-plus/fcgi/innerapi"
)

type fcgiServer struct {
	option         *FcgiServerOption
	childProcessor innerapi.ChildProcessor
}

type FcgiServerOption struct {
}

func NewFcgiServer(option *FcgiServerOption) *fcgiServer {
	fcgis := new(fcgiServer)
	fcgis.option = option
	return fcgis
}

func (this *fcgiServer) Serve(l net.Listener, handler http.Handler) error {
	this.childProcessor = fcgihttp.NewHttpProcessor()
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
