package main

import (
	"fmt"
	"sort"
)

func main() {
	seats := []int{12, 14, 19, 19, 12}
	students := []int{19, 2, 17, 20, 7}
	result := minMovesToSeat(seats, students)
	fmt.Println("Final result is = ", result)
}

func minMovesToSeat(seats []int, students []int) int {
	result := 0
	sort.Ints(seats)
	sort.Ints(students)

	for i := range students {
		if seats[i] > students[i] {
			result += seats[i] - students[i]
		} else {
			result += students[i] - seats[i]
		}
	}
	return result
}
