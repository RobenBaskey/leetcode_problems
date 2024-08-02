package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	mySlice := make([]map[string]int, 0) // [{key : value}]
	start := 0

	/// map => {p:1,w:1}

	for start < len(s) {
		fmt.Println(mySlice)
		//sData := string(s[start])

		//data := mySlice[0] //a

		// if data == 0 {
		// 	dataMap[sData]++
		// 	myValue += sData
		// 	fmt.Println("add ", dataMap)
		// } else {
		// 	dataMap = removeDataFromMap(dataMap, sData)
		// 	dataMap[sData]++
		// 	fmt.Println("remove", dataMap)
		// }

		// if result < len(dataMap) {
		// 	result = len(dataMap)
		// }

		fmt.Println(result)
		fmt.Println()

		start++
	}
	return result
}

func checkData(myData string, data string) bool {
	for _, key := range myData {
		return data == string(key)
	}
	return false
}

func removeDataFromMap(myData string, data string) string {
	for _, key := range myData {
		if string(key) == data {

		}
	}

	return ""
}
