package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

const (
	TotalMinutesPerHours = 60
	TotalMinutesPerDay   = 1440
)

func New(h, m int) Clock {
	totalMinutes := h*TotalMinutesPerHours + m
	totalMinutes %= TotalMinutesPerDay

	if totalMinutes < 0 {
		totalMinutes += TotalMinutesPerDay
	}

	return Clock{
		totalMinutes / TotalMinutesPerHours,
		totalMinutes % TotalMinutesPerHours,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
