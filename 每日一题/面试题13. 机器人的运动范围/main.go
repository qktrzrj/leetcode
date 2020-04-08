package main

import "fmt"

//地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，
// 它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
// 例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。
// 请问该机器人能够到达多少个格子？
//
//
//
//示例 1：
//
//输入：m = 2, n = 3, k = 1
//输出：3
//示例 1：
//
//输入：m = 3, n = 1, k = 0
//输出：1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func movingCount(m int, n int, k int) int {
	visited := make([][]byte, m)
	for i := range visited {
		visited[i] = make([]byte, n)
	}
	queue := make([][2]int, 1)
	queue[0][0], queue[0][1] = 0, 0
	var ans int
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		x, y := pos[0], pos[1]
		if x >= m || y >= n || visited[x][y] != 0 || add(x, y) > k {
			continue
		}
		visited[x][y] = '1'
		ans++
		queue = append(queue, [2]int{x + 1, y}, [2]int{x, y + 1})
	}
	return ans
}

func add(a, b int) int {
	var sum int
	for ans := a % 10; a > 0; ans = a % 10 {
		sum += ans
		a /= 10
	}
	for ans := b % 10; b > 0; ans = b % 10 {
		sum += ans
		b /= 10
	}
	return sum
}

func main() {
	fmt.Println(movingCount(15, 38, 9))
}
