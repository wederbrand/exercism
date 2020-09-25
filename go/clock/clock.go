// Package clock represents a simple clock.
package clock

import (
	"fmt"
)

const minutesPerDay = 24 * 60

// Clock represents a simple clock.
type Clock struct {
	minutes int
}

// New returns a new simple clock.
func New(hour int, minute int) Clock {
	minutes := hour*60 + minute

	minutes %= minutesPerDay

	if minutes < 0 {
		minutes += minutesPerDay
	}

	return Clock{minutes}
}

// String returns a textual representation of a simple clock with format HH:MM.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add adds any number of minutes to the simple clock, even negative values.
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

// Subtract subtracts any number of minutes from the simple clock, even positive values.
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}
