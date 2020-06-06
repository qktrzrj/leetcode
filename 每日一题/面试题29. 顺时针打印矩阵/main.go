package main

import "fmt"

// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
//
//
//示例 1：
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//示例 2：
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
