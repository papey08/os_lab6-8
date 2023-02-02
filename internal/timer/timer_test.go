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
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(correctAnswers[i]))
			answers[i] = timers[i].GetTime()
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

// TestTTimer_Pause tests
func TestTTimer_Pause(t *testing.T) {
	correctAnswers := []int{0, 1, 1, 2}
	answers := make([]int, len(correctAnswers))
	// test 01
	wg := new(sync.WaitGroup)
	wg.Add(len(answers))
	// test 01
	go func(i int) {
		defer wg.Done()
		timer := new(TTimer)
		answers[i] = timer.GetTime()
	}(0)
	// test 02
	go func(i int) {
		defer wg.Done()
		timer := new(TTimer)
		timer.Start()
		time.Sleep(time.Second * 1)
		timer.Pause()
		time.Sleep(time.Second * 1)
		answers[i] = timer.GetTime()
	}(1)
	// test 03
	go func(i int) {
		defer wg.Done()
		timer := new(TTimer)
		timer.Start()
		time.Sleep(time.Second * 1)
		timer.Pause()
		time.Sleep(time.Second * 1)
		timer.Start()
		answers[i] = timer.GetTime()
	}(2)
	// test 04
	go func(i int) {
		defer wg.Done()
		timer := new(TTimer)
		timer.Start()
		time.Sleep(time.Second * 1)
		timer.Pause()
		time.Sleep(time.Second * 1)
		timer.Start()
		time.Sleep(time.Second * 1)
		answers[i] = timer.GetTime()
	}(3)
	wg.Wait()
	for i, x := range answers {
		if correctAnswers[i] != x {
			t.Errorf("Test %d of %d: expected %d, found %d\n", i+1,
				len(correctAnswers), correctAnswers[i], answers[i])
		}
	}
}
