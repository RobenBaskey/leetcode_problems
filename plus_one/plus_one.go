package main

import "fmt"

func main() {
	result := plusOne([]int{9, 9, 9})
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	lastItem := digits[len(digits)-1]
	lastItem = lastItem + 1
	digits = digits[:len(digits)-1]
	if lastItem >= 10 {
		i := len(digits)
		inHand := lastItem / 10
		newArr := []int{}
		newArr = append(newArr, lastItem%10)
		for i > 0 {
			lastItem = digits[i-1]
			if lastItem+inHand >= 10 {
				inHand = (lastItem + inHand) / 10
				newArr = insertIntoFirst(newArr, (lastItem+inHand)%10)
			} else {
				newArr = insertIntoFirst(newArr, lastItem+inHand)
				inHand = 0
			}
			i--
		}
		if inHand != 0 {
			newArr = insertIntoFirst(newArr, inHand)
		}
		return newArr
	} else {
		digits = append(digits, lastItem)
	}
	return digits
}

func insertIntoFirst(slice []int, value int) []int {
	return append([]int{value}, slice...)
}
