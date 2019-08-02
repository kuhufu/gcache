package gcache

import (
	"github.com/kuhufu/scheduler"
	"time"
)

type Updater struct {
	scheduler *scheduler.Scheduler
	store     CacheStore
}

func NewUpdater(store CacheStore) *Updater {
	return &Updater{
		scheduler: scheduler.New(),
		store:     store,
	}
}

func (u *Updater) Start() {
	u.scheduler.Start()
}

func (u *Updater) Stop() {
	u.scheduler.Stop()
}

//@param immediately 是否立刻执行一次fetch函数
func (u *Updater) Interval(key string, seconds int, fetch func() interface{}, immediately bool) {
	if immediately {
		_ = u.store.Set(key, fetch(), -1)
	}
	u.scheduler.AddIntervalFunc(time.Duration(seconds)*time.Second, func() {
		_ = u.store.Set(key, fetch(), -1)
	})
}

//@param immediately 是否立刻执行一次fetch函数
func (u *Updater) Timeout(key string, seconds int, fetch func() interface{}, immediately bool) {
	if immediately {
		_ = u.store.Set(key, fetch(), -1)
	}
	u.scheduler.AddTimeoutFunc(time.Duration(seconds)*time.Second, func() {
		_ = u.store.Set(key, fetch(), -1)
	})
}
