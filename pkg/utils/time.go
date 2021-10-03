package utils

import (
	"context"
	"time"

	"github.com/araddon/dateparse"
	"go.uber.org/zap"
)

// ToTime converts string to time
func ToTime(in string) (time.Time, error) {
	t, err := dateparse.ParseLocal(in)
	if err != nil {
		Logger.Error("failed to parse time", zap.Error(err))
		return time.Time{}, err
	}

	return t, nil
}

func StartTimeAndLoc() (time.Time, *time.Location) {
	now := time.Now()
	yyyy, mm, dd := now.Date()
	return time.Date(yyyy, mm, dd, 9, 45, 2, 0, now.Location()), now.Location()
}

func EndTime() time.Time {
	now := time.Now()
	yyyy, mm, dd := now.Date()
	return time.Date(yyyy, mm, dd, 15, 0, 0, 0, now.Location())
}

func CurrentTime() string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

// Every sends the time to the returned channel at the specified intervals
func Every(ctx context.Context, start time.Time, interval time.Duration) <-chan time.Time {
	// Create the channel which we will return
	stream := make(chan time.Time, 3)

	// Calculating the first start time in the future
	// Need to check if the time is zero (e.g. if time.Time{} was used)
	if !start.IsZero() {
		diff := time.Until(start)
		if diff < 0 {
			total := diff - interval
			times := total / interval * -1

			start = start.Add(times * interval)
		}
	}

	// Run this in a goroutine, or our function will block until the first event
	go func() {

		// Run the first event after it gets to the start time
		timer := time.NewTimer(time.Until(start))
		defer timer.Stop() // Make sure to stop the timer when we're done

		// Listen on both the timer and the context done channel.
		// Useful if the context is closed before the first timer
		select {
		case t := <-timer.C:
			stream <- t
		case <-ctx.Done():
			close(stream)
			return
		}

		// Open a new ticker
		ticker := time.NewTicker(interval)
		defer ticker.Stop() // Make sure to stop the ticker when we're done

		// Listen on both the ticker and the context done channel to know when to stop
		for {
			select {
			case t2 := <-ticker.C:
				stream <- t2
			case <-ctx.Done():
				close(stream)
				return
			}
		}
	}()

	return stream
}

func WaitUntil(ctx context.Context, t time.Time) {
	diff := t.Sub(time.Now())
	if diff <= 0 {
		return
	}

	WaitFor(ctx, diff)
}

// WaitFor will block for the specified duration or the context is cancelled
func WaitFor(ctx context.Context, diff time.Duration) {
	timer := time.NewTimer(diff)
	defer timer.Stop()

	select {
	case <-timer.C:
		return
	case <-ctx.Done():
		return
	}
}
