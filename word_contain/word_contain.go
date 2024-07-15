package main

func main() {
	array := []string{"abc", "bcd", "aaaa", "cbc"}
	x := []byte("a")
	findWordsContaining(array, x[0])
}

func findWordsContaining(words []string, x byte) []int {
	var result = []int{}
	for i, value := range words {
		checkResult := checkByte([]byte(value), x)
		if checkResult {
			result = append(result, i)
		}
	}
	return result
}

func checkByte(byteArray []byte, target byte) bool {
	for _, b := range byteArray {
		if b == target {
			return true
		}
	}
	return false
}
