package timer

import (
	"sync"
	"testing"
	"time"
)

// TestTTimer_GetTime creates some timers and tests GetTime method using time.Sleep and goroutines
func TestTTimer_GetTime(t *testing.T) {
	correctAnswers := []int{5, 6, 5, 1, 1, 1, 0, 0, 0}
	answers := make([]int, len(correctAnswers))
	timers := make([]*TTimer, len(correctAnswers))
	for i := range timers {
		timers[i] = new(TTimer)
	}
	wg := new(sync.WaitGroup)
	for _, timer := range timers {
		timer.Start()
	}
	for i := range timers {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Second * time.Duration(correctAnswers[i]))
			answers[i] = timers[i].GetTime()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := range answers {
		if answers[i] != correctAnswers[i] {
			t.Errorf("Test %d of %d: expected %d, found %d\n", i+1,
				len(answers), correctAnswers[i], answers[i])
		}
	}
}
