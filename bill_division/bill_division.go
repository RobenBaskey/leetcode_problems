package main

import "fmt"

func main() {
	bill := []int32{3, 10, 2, 9}
	bonAppetit(bill, 1, 7)
}

func bonAppetit(bill []int32, k int32, b int32) {
	var sum int32
	var i int32
	for i = 0; i < int32(len(bill)); i++ {
		if i != k {
			sum = sum + bill[i]
		}
	}

	if (sum / 2) == b {
		fmt.Println("Bon Appetit")
		return
	}

	perBill := sum / 2
	var overCharged int32
	if b > perBill {
		overCharged = b - perBill
	} else {
		overCharged = perBill - b
	}
	fmt.Println(overCharged)

}
