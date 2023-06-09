// Synchronized package
package syncs

import (
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestAsyncWait(t *testing.T) {
	var atomicResultOne atomic.Value
	var atomicResultTwo atomic.Value
	var syncsGroup AsyncGroup

	syncsGroup.Async(func() {
		atomicResultOne.Store("hello")
	})
	syncsGroup.Async(func() {
		atomicResultTwo.Store("world")
	})

	syncsGroup.AsyncWait()

	assert.Equal(t, "hello", atomicResultOne.Load())
	assert.Equal(t, "world", atomicResultTwo.Load())
}
