package gcache

import (
	"fmt"
	"testing"
	"time"
)

var s = NewUpdaterOf(memc)

func TestNewSchedulerOf(t *testing.T) {
	s.Interval("k1", 1, func() interface{} {
		return time.Now().String()
	}, true)

	for {
		fmt.Println(memc.Get("k1"))
		time.Sleep(time.Second)
	}
}

func TestScheduler_Interval(t *testing.T) {
	s.Interval("k1", 1, func() interface{} {
		return time.Now().String()
	}, true)

	for {
		fmt.Println(memc.Get("k1"))
		time.Sleep(time.Second)
	}
}

func TestScheduler_Timeout(t *testing.T) {
	s.Start()
	s.Timeout("k1", 1, func() interface{} {
		return time.Now().String()
	}, true)
	fmt.Println(memc.Get("k1").String())
	time.Sleep(time.Second * 4)
	fmt.Println(memc.Get("k1").String())
}
