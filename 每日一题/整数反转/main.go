package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	minus := x < 0
	if minus {
		x = -x
	}
	var ans int
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
		if minus && -ans < math.MinInt32 {
			return 0
		} else if !minus && ans > math.MaxInt32 {
			return 0
		}
	}
	if minus {
		ans = -ans
	}
	return ans
}

func main() {
	fmt.Println(reverse(321))
}
