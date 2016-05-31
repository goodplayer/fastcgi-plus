package fcgi

const (
	_FCGI_VERSION_1 byte = 1

	//TODO record type
	// at least one stream record in each direction
	// empty stream record means the stream record finished
	// 1. management records
	_FCGI_GET_VALUES        = 9  // web to app(in)
	_FCGI_GET_VALUES_RESULT = 10 // app to web(out)
	_FCGI_UNKNOWN_TYPE      = 11 // app to web(out)
	// application records
	_FCGI_BEGIN_REQUEST = 1 // web to app(in)
	_FCGI_ABORT_REQUEST = 2 // web to app(in)
	_FCGI_PARAMS        = 4 //stream record, web to app(in)
	_FCGI_STDIN         = 5 //stream record, web to app(in)
	_FCGI_DATA          = 8 //stream record, web to app(in)
	_FCGI_STDOUT        = 6 //stream record, app to web(out)
	_FCGI_STDERR        = 7 //stream record, app to web(out)
	_FCGI_END_REQUEST   = 3 // app to web(out)
	// 2. other
	_FCGI_MAXTYPE = _FCGI_UNKNOWN_TYPE

	// protocol status
	_FCGI_REQUEST_COMPLETE = 0
	_FCGI_CANT_MPX_CONN    = 1
	_FCGI_OVERLOADED       = 2
	_FCGI_UNKNOWN_ROLE     = 3

	// role
	_FCGI_RESPONDER  = 1
	_FCGI_AUTHORIZER = 2
	_FCGI_FILTER     = 3

	_FCGI_NULL_REQUEST_ID = 0

	// flag
	_FCGI_KEEP_CONN = 1

	// values
	_FCGI_MAX_CONNS  = "FCGI_MAX_CONNS"
	_FCGI_MAX_REQS   = "FCGI_MAX_REQS"
	_FCGI_MPXS_CONNS = "FCGI_MPXS_CONNS"
)

const (
	default_buffer_size = 64 * 1024
)

var (
	unknown_type_packet unknownTypeMessage = [16]byte{
		_FCGI_VERSION_1,
		_FCGI_UNKNOWN_TYPE,
		0, 0,
		0, 8,
		0, 0,

		0, // type - predefine = 0, must change
		0, 0, 0, 0, 0, 0, 0,
	}
)
