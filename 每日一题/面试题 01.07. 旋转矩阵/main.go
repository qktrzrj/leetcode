package main

import "fmt"

//给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
//
//不占用额外内存空间能否做到？
//
//
//
//示例 1:
//
//给定 matrix =
//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
//示例 2:
//
//给定 matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-matrix-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// i->j,j->n-1-i
// 逐层转换,循环连续赋值
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	if n <= 0 {
		return
	}
	var nexti, nextj, sw int
	level := (n + 1) / 2
	for i := 0; i < level; i++ {
		for j := i; j < n-i; j++ {
			nexti, nextj = j, n-i
			var over bool
			sw = matrix[i][j]
			for !over {
				matrix[nexti][nextj], sw = sw, matrix[nexti][nextj]
				nexti, nextj = nextj, n-nexti
				if nexti == i && nextj == j {
					matrix[i][j] = sw
					over = true
				}
			}
		}
	}
}

func main() {
	m := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(m)
	fmt.Println(m)
}
