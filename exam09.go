package main

import (
	"fmt"
)

func main() {
	const (
		day1 = iota + 1
		day2
		day3
		day4
		day5
		day6
		day7
	)
	fmt.Println(day1, day2, day3, day4, day5, day6, day7)
	const a, b, c = iota, iota + 1, iota + 1
	fmt.Println(a, b, c)
}
