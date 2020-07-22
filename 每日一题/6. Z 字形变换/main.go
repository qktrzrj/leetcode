package main

import (
	"fmt"
	"math"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, int(math.Min(float64(numRows), float64(len(s)))))
	var curRow int
	var down bool
	for _, r := range s {
		rows[curRow].WriteRune(r)
		if curRow == 0 || curRow == numRows-1 {
			down = !down
		}
		if down {
			curRow++
		} else {
			curRow--
		}
	}

	var res strings.Builder
	for _, sb := range rows {
		res.WriteString(sb.String())
	}
	return res.String()
}

func main() {
	fmt.Println(convert("AB", 1))
}
