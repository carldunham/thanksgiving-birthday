package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFallsOnThanksgivingAfter(t *testing.T) {
	cases := []struct {
		Name     string
		After    uint16
		Day      uint8
		Expected uint16
		Err      error
	}{
		{
			Name:     "default",
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
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			actual, err := fallsOnThanksgivingAfter(c.After, c.Day)

			require.ErrorIs(t, err, c.Err)
			assert.Equal(t, c.Expected, actual)
		})
	}
}
