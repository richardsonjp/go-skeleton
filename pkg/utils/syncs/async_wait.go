// Synchronized package
package syncs

import (
	"sync"
)

type AsyncGroup struct {
	wait sync.WaitGroup
}

// Async Do the async wait process functionality
func (async *AsyncGroup) Async(callback func()) {
	async.wait.Add(1)
	GoRecover(func() {
		defer async.wait.Done()
		callback()
	})
}

// AsyncWait Wait until all process in async group is finished
func (async *AsyncGroup) AsyncWait() {
	async.wait.Wait()
}
