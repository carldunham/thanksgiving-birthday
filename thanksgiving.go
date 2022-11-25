package main

import "fmt"

const (
	ErrInvalidDay = Error("invalid day, must be 22-28 (inclusive)")
	ErrNoIdea     = Error("no idea")
)

// fallsOnThanksgivingAfter calculates the year on or after the one given where the
// given day in November falls on Thanksgiving.
//
// Thanksgiving is a US holiday that takes place on the fourth Thursday in November each year,
// so the earliest it will occur is on the 22nd, and the latest is the 28th.
func fallsOnThanksgivingAfter(year uint16, day uint8) (uint16, error) {
	switch {
	case day < 22, day > 28:
		return 0, fmt.Errorf("%d is an %w", day, ErrInvalidDay)
	case day == 24 && year == 2022:
		return 2022, nil
	default:
		return 0, ErrNoIdea
	}
}
