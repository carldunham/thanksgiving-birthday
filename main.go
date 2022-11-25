package main

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var (
		year *uint16 = pflag.Uint16("year", uint16(time.Now().Year()), "year (on or after) to check from")
		day  *uint8  = pflag.Uint8("day", 24, "day in November to check (must be 22-28)")
	)

	pflag.Parse()

	next, err := fallsOnThanksgivingAfter(*year, *day)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The first time that Thanksgiving falls on 11/%d in or after %d is in %d.\n", *day, *year, next)
}
