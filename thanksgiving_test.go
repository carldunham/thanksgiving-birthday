package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFallsOnThanksgivingAfter(t *testing.T) {
	cases := []struct {
		Name     string
		After    int
		Day      int
		Expected int
		Err      error
	}{
		{
			Name:     "current",
			After:    2022,
			Day:      24,
			Expected: 2022,
			Err:      nil,
		},
		{
			Name:     "day too early",
			After:    2022,
			Day:      17,
			Expected: 0,
			Err:      ErrInvalidDay,
		},
		{
			Name:     "day too late",
			After:    2022,
			Day:      29,
			Expected: 0,
			Err:      ErrInvalidDay,
		},
		{
			Name:     "negative year",
			After:    -1,
			Day:      22,
			Expected: 0,
			Err:      ErrInvalidYear,
		},
		{
			Name:     "first",
			After:    2001,
			Day:      24,
			Expected: 2005,
			Err:      nil,
		},
		{
			Name:     "second",
			After:    2006,
			Day:      24,
			Expected: 2011,
			Err:      nil,
		},
		{
			Name:     "third",
			After:    2012,
			Day:      24,
			Expected: 2016,
			Err:      nil,
		},
		{
			Name:     "fourth",
			After:    2017,
			Day:      24,
			Expected: 2022,
			Err:      nil,
		},
		{
			Name:     "future",
			After:    2023,
			Day:      24,
			Expected: 2033,
			Err:      nil,
		},
		{
			Name:     "theoretical first",
			After:    0,
			Day:      24,
			Expected: 5,
			Err:      nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			actual, err := fallsOnThanksgivingAfter(c.After, c.Day)

			require.ErrorIs(t, err, c.Err)
			assert.Equal(t, c.Expected, actual)
		})
	}
}

func TestAgainstBruteForceApproach(t *testing.T) {
	for day := 22; day <= 28; day++ {
		for year := 1900; year <= 2200; year++ {
			t.Run(fmt.Sprintf("year=%d day=%d", year, day), func(t *testing.T) {
				expected := year
				for time.Date(expected, 11, day, 0, 0, 0, 0, time.Local).Weekday() != time.Thursday {
					expected++
				}
				actual, err := fallsOnThanksgivingAfter(year, day)

				require.NoError(t, err)
				assert.Equal(t, expected, actual)
			})
		}
	}
}
