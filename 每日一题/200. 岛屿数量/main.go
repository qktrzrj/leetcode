package main

import "fmt"

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
//示例 1:
//
//输入:
//11110
//11010
//11000
//00000
//输出: 1
//示例 2:
//
//输入:
//11000
//11000
//00100
//00011
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-of-islands
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// BFS: 遍历网格，搜索到1，将其加入队列，开始进行BFS搜索，直到队列为空

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	n, m := len(grid), len(grid[0]) // 行数，列数
	var num int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				num++
				grid[i][j] = '0'
				queue := []int{i*m + j}
				for len(queue) > 0 {
					row, col := queue[len(queue)-1]/m, queue[len(queue)-1]%m
					queue = queue[:len(queue)-1]
					// 竖直
					if row-1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, (row-1)*m+col)
						grid[row-1][col] = '0'
					}
					if row+1 < n && grid[row+1][col] == '1' {
						queue = append(queue, (row+1)*m+col)
						grid[row+1][col] = '0'
					}
					// 水平
					if col-1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, row*m+col-1)
						grid[row][col-1] = '0'
					}
					if col+1 < m && grid[row][col+1] == '1' {
						queue = append(queue, row*m+col+1)
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}

	return num
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}))
}
