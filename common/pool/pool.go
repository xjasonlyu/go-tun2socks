package pool

import (
	"sync"
)

const (
	// io.Copy default buffer size is 32 KiB
	// define a buffer of 20 KiB to reduce the memory of each TCP relay
	bufferSize = 20 * 1024
)

// BufPool provide buffer for relay
var BufPool = sync.Pool{New: func() interface{} { return make([]byte, bufferSize) }}
