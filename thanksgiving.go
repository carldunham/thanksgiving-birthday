package main

import (
	"fmt"
	"time"
)

const (
	ErrInvalidDay  = Error("invalid day, must be 22-28 (inclusive)")
	ErrInvalidYear = Error("invalid year, must be >= 0")
	ErrNoIdea      = Error("no idea")
)

// fallsOnThanksgivingAfter calculates the year on or after the one given where the
// given day in November falls on Thanksgiving.
//
// Thanksgiving is a US holiday that takes place on the fourth Thursday in November each year,
// so the earliest it will occur is on the 22nd, and the latest is the 28th.
func fallsOnThanksgivingAfter(year int, day int) (int, error) {
	switch {
	case day < 22, day > 28:
		return 0, fmt.Errorf("%d is an %w", day, ErrInvalidDay)
	case year < 0:
		return 0, fmt.Errorf("%d is an %w", year, ErrInvalidYear)
	default:
		delta := int(time.Thursday - time.Date(year, 11, day, 0, 0, 0, 0, time.Local).Weekday())
		switch {
		case delta == 0:
			return year, nil
		case delta < 0:
			delta += 7
		}
		target := year + delta
		next := target - leapYearsBetween(year, target)
		next += leapYearsBetween(next, target)

		if isLeapYear(next) {
			return fallsOnThanksgivingAfter(next, day)
		}
		return next, nil
	}
}

func leapYearsBetween(from, to int) int {
	var ret int

	for y := from + 1; y < to; y++ {
		if isLeapYear(y) {
			ret++
		}
	}
	return ret
}

func isLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	default:
		return year%4 == 0
	}
}
