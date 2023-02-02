package timer

import "time"

type TTimer struct {
	running bool
	start   time.Time
	worked  int
}

// Start launches the timer or resumes timer which is on pause
func (t *TTimer) Start() {
	if t.running {
		return
	}
	t.running = true
	t.start = time.Now()
}

// GetTime returns how many seconds passed since timer was started
func (t *TTimer) GetTime() int {
	if t.running {
		return t.worked + int(time.Since(t.start).Seconds())
	} else {
		return t.worked
	}
}

// Pause sets timer on pause. Timer could be resumed by Start method
func (t *TTimer) Pause() {
	t.worked = t.GetTime()
	t.running = false
}

// Reset resets time on the timer
func (t *TTimer) Reset() {
	t.worked = 0
	t.running = false
}
