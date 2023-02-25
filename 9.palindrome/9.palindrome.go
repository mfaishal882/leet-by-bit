package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)

	for i := 1; i < len(strX); i++ {
		// fmt.Println(string(strX[i-1]))
		// fmt.Println(string(strX[len(strX)-i]))
		// fmt.Println("==========")
		if strX[i-1] != strX[len(strX)-i] {
			// fmt.Println(" str x[i]", strX[i], " strx[-1]", strX[len(strX)-1])
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(10))    // false
	fmt.Println(isPalindrome(121))   // true
	fmt.Println(isPalindrome(12233)) // false
}

// fastest 3ms
/*
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x < 9 {
        return true
    }
    x_copy := x
    tmp := 0
    for x > 0 {
        tmp = (tmp * 10) + (x % 10)
        x = x / 10
    }
    return x_copy == tmp
}
*/
