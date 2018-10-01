package ledlib

import (
	"time"
)

type Timer struct {
	start       time.Time
	last_update time.Time
	interval    time.Duration
}

func NewTimer(interval time.Duration) *Timer {
	return &Timer{time.Now(), time.Now(), interval}
}

func (t *Timer) ResetTimer() {
	t.last_update = time.Now()
}

func (t *Timer) IsPast() bool {
	sub := time.Now().Sub(t.last_update)
	if sub > t.interval {
		t.ResetTimer()
		return true
	}
	return false
}

func (t *Timer) GetElapsed() time.Duration {
	return time.Now().Sub(t.last_update)
}

func (t *Timer) GetPastCount() uint64 {
	sub := time.Now().Sub(t.start)
	return uint64(sub / t.interval)
}
