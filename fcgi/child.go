package fcgi

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/goodplayer/fastcgi-plus/fcgi/innerapi"
)

type child struct {
	conn        net.Conn
	recordChan  chan reqWriter
	r           *bufio.Reader
	requests    map[uint16]*statefulRequest
	requestLock sync.RWMutex
}

func (this *child) reset() {
	//TODO
	this.conn = nil
	this.recordChan = nil
	this.r = nil
	this.requests = nil
}

func startChildHandleLoop(conn net.Conn, handler http.Handler, fcgis *fcgiServer) error {
	c := getChild()
	c.conn = conn
	c.r = bufio.NewReader(conn)
	c.recordChan = make(chan reqWriter, 64)
	c.requests = make(map[uint16]*statefulRequest)

	go c.childHandleProcessor(fcgis.childProcessor)
	go c.outboundProcessor()

	//TODO
	return nil
}

func (this *child) release() {
	this.conn.Close()
}

func (this *child) childHandleProcessor(cp innerapi.ChildProcessor) {
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
		bizErr := this.packetDispatching(req, requests, cp)
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

func (this *child) packetDispatching(req request, reqMap map[uint16]*statefulRequest, cp innerapi.ChildProcessor) error {
	reqId := req.Header.getRequestId()
	ptype := req.Header.Type
	if reqId == 0 {
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
		this.requestLock.RLock()
		r, ok := reqMap[reqId]
		this.requestLock.RUnlock()
		if !ok { // request id not exist
			if ptype != _FCGI_BEGIN_REQUEST {
				// not begin request, just send end request
				end := end_request_message
				end.setAppStatus(1)
				end.setProtocolStatus(_FCGI_UNKNOWN_ROLE)
				end.setRequestId(reqId)
				this.recordChan <- end
				return nil // may not return error
			} else {
				if req.ContentData[0] != 0 || req.ContentData[1] != _FCGI_RESPONDER {
					// non responder role. currently only support responder
					end := end_request_message
					end.setAppStatus(1)
					end.setProtocolStatus(_FCGI_UNKNOWN_ROLE)
					end.setRequestId(reqId)
					this.recordChan <- end
					return nil // may not return error
				}
				// always support keepalive, so no flags check
				// init a request
				sreq := getStatefulRequest()
				this.requestLock.Lock()
				reqMap[reqId] = sreq
				this.requestLock.Unlock()
				r = sreq
				r.state = _STATEFUL_REQUEST_STATE_READING_PARAM
				return nil
			}
		} else {
			// request id exists.
			switch r.state {
			case _STATEFUL_REQUEST_STATE_READING_PARAM:
				// reading param
				if len(req.ContentData) == 0 {
					cp.ProcessParam(r)
					r.state = _STATEFUL_REQUEST_STATE_READING_STDIN
				} else {
					err := parseNvPair(r, req.ContentData)
					if err != nil {
						end := end_request_message
						end.setAppStatus(4)
						end.setProtocolStatus(_FCGI_UNKNOWN_ROLE)
						end.setRequestId(reqId)
						this.recordChan <- end
						this.requestLock.Lock()
						delete(reqMap, reqId)
						this.requestLock.Unlock()
						return nil
					}
				}
			case _STATEFUL_REQUEST_STATE_READING_STDIN:
				// reading stdin
				if len(req.ContentData) == 0 {
					r.state = _STATEFUL_REQUEST_STATE_READING_DONE
				} else {
					//TODO
				}
			case _STATEFUL_REQUEST_STATE_READING_DATA:
				// currently not support reading data
				end := end_request_message
				end.setAppStatus(3)
				end.setProtocolStatus(_FCGI_UNKNOWN_ROLE)
				end.setRequestId(reqId)
				this.recordChan <- end
				this.requestLock.Lock()
				delete(reqMap, reqId)
				this.requestLock.Unlock()
				return nil
			default:
				// others may be error state
				end := end_request_message
				end.setAppStatus(2)
				end.setProtocolStatus(_FCGI_UNKNOWN_ROLE)
				end.setRequestId(reqId)
				this.recordChan <- end
				this.requestLock.Lock()
				delete(reqMap, reqId)
				this.requestLock.Unlock()
				return nil
			}
		}
	}
	return nil
}
