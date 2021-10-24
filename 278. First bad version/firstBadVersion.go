package main

import "fmt"

const badVersion = 3

func isBadVersion(version int) bool {
	if badVersion > version {
		return false
	} else {
		return true
	}
}

func firstBadVersion(n int) int {
	for i := 0; i < n; {
		m := (i + n + 1) / 2
		if isBadVersion(m) {
			if n == m {
				return n
			}
			n = m
		} else {
			i = m
		}
	}
	return n
}

func main() {
	fmt.Print(firstBadVersion(10))
}
