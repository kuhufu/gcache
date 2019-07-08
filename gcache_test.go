package gcache

import (
	"fmt"
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	StartSchedule()
	Interval(cache, "k1", 2, func() []byte {
		return []byte(time.Now().String())
	}, true)

	for {
		value, _ := cache.Get("k1")
		fmt.Println(string(value))
		time.Sleep(time.Second * 2)
	}
}
