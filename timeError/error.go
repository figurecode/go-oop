package timeError

import (
	"fmt"
	"time"
)

type TimeError struct {
	Time time.Time
	Text string
}

func (te TimeError) Error() string {
	return fmt.Sprintf("%v: %v", te.Time.Format("2006/01/02 15:04:05"), te.Text)
}

func NewTimeError(text string) TimeError {
	return TimeError{
		Time: time.Now(),
		Text: text,
	}
}
