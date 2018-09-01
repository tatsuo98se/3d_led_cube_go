package ledlib

import "time"

type Timer struct {
	last_update int64
	interval    int64
}

func (t *Timer) getUpdateTime() int64 {
	return t.last_update + t.interval
}

func NewTimer(intervalInMsec int64) *Timer {
	return &Timer{time.Now().UnixNano(), intervalInMsec * 1000 * 1000}
}

func (t *Timer) ResetTimer() {
	t.last_update = time.Now().UnixNano()
}

func (t *Timer) IsPast() bool {
	now := time.Now().UnixNano()
	if now > t.getUpdateTime() {
		t.ResetTimer()
		return true
	} else {
		return false
	}
}
