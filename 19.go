/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(1901, time.January, 1, 12, 0, 0, 0, time.UTC)
	end := time.Date(2000, time.December, 31, 12, 0, 0, 0, time.UTC)
	oneDay := time.Hour * 24
	counter := 0

	for start != end {
		if start.Weekday() == time.Sunday && start.Day() == 1 {
			counter++
		}
		start = start.Add(oneDay)
	}
	fmt.Printf("%d Sundays fell on the first of the month during the twentieth century.\n", counter)
}
