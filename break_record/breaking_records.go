package main

import (
	"fmt"
)

func main() {
	a := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	result := getBreakRecord(a)
	for _, data := range result {
		fmt.Println("Data ", data)
	}
}

func getBreakRecord(scores []int32) []int32 {
	var breakRecordData int32 = 0
	var lowRecordData int32 = 0
	var mostRecord int32 = 0
	var lowRecord int32 = 0
	result := []int32{}
	if len(scores) != 0 {
		breakRecordData = scores[0]
		lowRecordData = scores[0]
		for _, data := range scores {
			if breakRecordData < data {
				mostRecord++
				breakRecordData = data
			}

			if lowRecordData > data {
				lowRecord++
				lowRecordData = data
			}
		}
	}
	result = append(result, mostRecord)
	result = append(result, lowRecord)
	return result
}
