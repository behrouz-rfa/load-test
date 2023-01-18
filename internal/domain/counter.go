package domain

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter Stores counts associated with a key.
type Counter struct {
	m sync.Map
}

// Get Retrieves the count without modifying it
func (c *Counter) Get(key string) int64 {
	count, ok := c.m.Load(key)
	if ok {
		return atomic.LoadInt64(count.(*int64))
	}
	return 0
}

// Add Adds value to the stored underlying value if it exists.
// If it does not exist, the value is assigned to the key.
func (c *Counter) Add(key string, value int64) int64 {
	count, loaded := c.m.LoadOrStore(key, &value)
	if loaded {
		return atomic.AddInt64(count.(*int64), value)
	}
	return *count.(*int64)
}

// DeleteAndGetLastValue Deletes the value associated with the key and retrieves it.
func (c *Counter) DeleteAndGetLastValue(key string) (int64, bool) {

	lastValue, loaded := c.m.LoadAndDelete(key)
	if loaded {
		return *lastValue.(*int64), loaded
	}
	return 0, false
}

func (c *Counter) Iter() {
	c.m.Range(func(key, value any) bool {
		fmt.Println("StatusCode: ", key, "counter: ", atomic.LoadInt64(value.(*int64)))
		return true
	})
}
