package oncer

import (
	"fmt"
	"runtime"
	"sync"
)

var caller = &sync.Map{}

func Do(fn func(), withline ...bool) {
	var (
		file string
		line int
		ok   bool
		once *sync.Once
	)

	if len(withline) > 0 && withline[0] {
		_, file, line, ok = runtime.Caller(1)
		if !ok {
			return
		}
		file = fmt.Sprintf("%s:%d", file, line)
	} else {
		_, file, _, ok = runtime.Caller(1)
		if !ok {
			return
		}
	}
	if v, ok := caller.Load(file); ok {
		once = v.(*sync.Once)
	} else {
		once = new(sync.Once)
		caller.Store(file, once)
	}
	once.Do(fn)
}
