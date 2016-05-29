package fcgi

const (
	_FCGI_VERSION_1 byte = 1

	//TODO record type
	// management records
	_FCGI_GET_VALUES        = 9  // web to app
	_FCGI_GET_VALUES_RESULT = 10 // app to web
	_FCGI_UNKNOWN_TYPE      = 11 // app to web
	// application records
	_FCGI_BEGIN_REQUEST = 1 // web to app
	_FCGI_PARAMS        = 4 //stream record, web to app
	_FCGI_STDIN         = 5 //stream record, web to app
	_FCGI_DATA          = 8 //stream record, web to app
	_FCGI_STDOUT        = 6 //stream record, app to web
	_FCGI_STDERR        = 7 //stream record, app to web
	_FCGI_ABORT_REQUEST = 2 // web to app
	_FCGI_END_REQUEST   = 3 // app to web
	// other
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
	default_buffer_size = 4096
)
