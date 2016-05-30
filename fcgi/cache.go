package fcgi

import "sync"

var (
	bufCache = sync.Pool{
		New: func() interface{} {
			return make([]byte, default_buffer_size)
		},
	}

	childCache = sync.Pool{
		New: func() interface{} {
			return new(child)
		},
	}
)

func getBuffer() []byte {
	return bufCache.Get().([]byte)
}

func returnBuffer(b []byte) {
	bufCache.Put(b)
}

func getChild() *child {
	c := childCache.Get().(*child)
	c.reset()
	return c
}

func returnChild(c *child) {
	c.reset()
	childCache.Put(c)
}
