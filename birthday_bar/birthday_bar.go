package main

import "fmt"

func main() {
	s := []int32{2, 2, 2, 1, 3, 2, 2, 3, 3, 1, 4, 1, 3, 2, 2, 1, 2, 2, 4, 2, 2, 3, 5, 3, 4, 3, 2, 1, 4, 5, 4}
	var d int32 = 10
	var m int32 = 4

	result := birthday(s, d, m)
	fmt.Println("My result is ", result)
}

func birthday(s []int32, d int32, m int32) int32 {
	var result int32 = 0

	if m == 1 {
		for _, data := range s {
			if data == d {
				result++
			}
		}

		return result
	}

	length := len(s)

	if m > int32(length) {
		return 0
	}

	for i := 0; i <= length-int(m); i++ {
		subData := s[i:(int(m) + i)]
		var sum int32 = 0
		for _, value := range subData {
			sum += value
		}
		if sum == d {
			result++
		}
	}
	return result

}
