package fcgi

import (
	"sync"
	"sync/atomic"
)

var (
	childCache = sync.Pool{
		New: func() interface{} {
			return new(child)
		},
	}

	bufItemCache = sync.Pool{
		New: func() interface{} {
			return &BufItem{}
		},
	}

	statefulRequestCache = sync.Pool{
		New: func() interface{} {
			return new(statefulRequest)
		},
	}
)

type BufItem struct {
	data        [default_buffer_size]byte
	ref         int32
	releaseFunc func(*BufItem)
}

func getBufItem() *BufItem {
	b := bufItemCache.Get().(*BufItem)
	b.releaseFunc = returnBufItem
	b.ref = 1
	return b
}

func returnBufItem(buf *BufItem) {
	bufItemCache.Put(buf)
}

func (this *BufItem) GetBuffer() []byte {
	return this.data[:]
}

func (this *BufItem) Retain() {
	if atomic.LoadInt32(&this.ref) <= 0 {
		panic("BufItem is released. should not retain!")
	}
	atomic.AddInt32(&this.ref, 1)
}

func (this *BufItem) Release() {
	i := atomic.AddInt32(&this.ref, -1)
	if i == 0 {
		this.releaseFunc(this)
	} else if i < 0 {
		panic("BufItem ref < 0")
	}
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

func getStatefulRequest() *statefulRequest {
	r := statefulRequestCache.Get().(*statefulRequest)
	r.reset()
	return r
}

func returnStatefulRequest(request *statefulRequest) {
	request.reset()
	statefulRequestCache.Put(request)
}
