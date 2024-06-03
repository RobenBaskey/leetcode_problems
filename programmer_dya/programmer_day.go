package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := dayOfProgrammer(1918)
	fmt.Println(result)
}

func dayOfProgrammer(year int32) string {
	monthSlice := []int{31, 28, 31, 30, 31, 30, 31, 31}
	if year == 1918 {
		myYear := fmt.Sprintf("%d", year)
		myStr := strconv.Itoa(26) + ".09." + myYear
		return myStr
	}
	if year <= 1918 {
		if year%4 == 0 {
			monthSlice[1] = 29
		}
	} else {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			monthSlice[1] = 29
		}
	}

	sum := 0
	for _, data := range monthSlice {
		sum = sum + data
	}
	day := 256 - sum
	myYear := fmt.Sprintf("%d", year)
	myStr := strconv.Itoa(day) + ".09." + myYear

	return myStr
}
