package main

import "fmt"

func isPalindrome(x int) bool {
	tmp, rev := x, 0
	for tmp > 0 {
		rev = rev*10 + tmp%10
		tmp /= 10
	}
	return x == rev && x >= 0
}

func main() {
	fmt.Print(isPalindrome(11))
}
