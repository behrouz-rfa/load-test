package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"load-test/internal/domain"
	"math/rand"
	"testing"
)

func TestCounter_AddGet(t *testing.T) {
	counter := &domain.Counter{}
	for i := 0; i < 20; i++ {
		counter.Add("1", 1)
	}
	load := counter.Get("1")

	assert.Equal(t, load, int64(20))

}
func TestCounter_AddGetConncurrent(t *testing.T) {
	counter := &domain.Counter{}
	for i := 0; i < 5000; i++ {
		go func() {
			randomIndex := rand.Intn(len(items))
			pick := items[randomIndex]
			counter.Add(fmt.Sprint(pick), 1)
		}()
	}
	counter.Iter()
}

var items = []int{100, 101, 102, 103, 200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 425, 426, 428, 429, 431, 451, 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511}
