package helpers

import (
	"sync"
)

var seqMutex sync.Mutex

func ExecuteSequentially() func() {
	seqMutex.Lock()
	return func() {
		seqMutex.Unlock()
	}
}