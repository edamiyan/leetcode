package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	runeNumber := []rune(s)
	result := romanMap[runeNumber[len(runeNumber)-1]]

	for i := len(runeNumber) - 2; i >= 0; i-- {
		if romanMap[runeNumber[i]] < romanMap[runeNumber[i+1]] {
			result -= romanMap[runeNumber[i]]
		} else {
			result += romanMap[runeNumber[i]]
		}
	}
	return result
}

func main() {
	romanNum := "MMXXI"
	fmt.Print(romanToInt(romanNum))
}
