package timer

import "time"

type TTimer struct {
	isStarted bool
	start     time.Time
	//finish    time.Time
}

// Start launches the timer
func (t *TTimer) Start() {
	t.isStarted = true
	t.start = time.Now()
}

// GetTime returns how many seconds passed since timer was started
func (t *TTimer) GetTime() int {
	if t.isStarted {
		return int(time.Since(t.start).Seconds())
	} else {
		return 0
	}
}
