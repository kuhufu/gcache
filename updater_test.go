package gcache

import (
	"fmt"
	"testing"
	"time"
)

var s = NewUpdaterOf(cache)

func TestNewSchedulerOf(t *testing.T) {
	s.Interval("k1", 1, func() []byte {
		return []byte(time.Now().String())
	}, true)

	for {
		fmt.Println(cache.Get("k1"))
		time.Sleep(time.Second)
	}
}

func TestScheduler_Interval(t *testing.T) {
	s.Interval("k1", 1, func() []byte {
		return []byte(time.Now().String())
	}, true)

	for {
		fmt.Println(cache.Get("k1"))
		time.Sleep(time.Second)
	}
}

func TestScheduler_Timeout(t *testing.T) {
	s.Start()
	s.Timeout("k1", 1, func() []byte {
		return []byte(time.Now().String())
	}, true)

	var value, _ = cache.Get("k1")
	fmt.Println(string(value))
	time.Sleep(2 * time.Second)
	value, _ = cache.Get("k1")
	fmt.Println(string(value))
}
