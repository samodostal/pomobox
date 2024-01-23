package main

import (
	"main/button"
	"main/clock"
	"main/counter"
	"time"
)

func main() {
	button := button.New()
	counter := counter.New()

	clock := clock.New()
	clock.SetTime("50:00")
	clock.SetHardwareSecondInMs(640)

	isPomoRunning := false

	minutes, seconds := clock.GetTimeInt()
	totalTimeMs := 0
	buttonDelayMs := 0

	for {
		if isPomoRunning {
			if totalTimeMs > 0 && totalTimeMs%clock.GetHardwareSecondInMs() == 0 {
				if seconds == 0 && minutes > 0 {
					minutes--
					seconds = 60
				}
				if seconds != 0 {
					seconds--
				}
				if seconds == 0 && minutes == 0 {
					isPomoRunning = false
					minutes, seconds = clock.GetTimeInt()
					counter.Increment()
				}
			}
		}

		if !button.Get() {
			if buttonDelayMs == 0 {
				isPomoRunning = !isPomoRunning
				buttonDelayMs = 200
			}
		}

    if buttonDelayMs > 0 {
      buttonDelayMs -= 1
    }

		clock.ResetPins()
		clock.DisplayDigit(seconds%10, 0)
		time.Sleep(time.Millisecond)

		clock.ResetPins()
		clock.DisplayDigit(seconds/10, 1)
		time.Sleep(time.Millisecond)

		clock.ResetPins()
		clock.DisplayDigit(minutes%10, 2)
		time.Sleep(time.Millisecond)

		clock.ResetPins()
		clock.DisplayDigit(minutes/10, 3)
		time.Sleep(time.Millisecond)

		totalTimeMs += 4
	}
}
