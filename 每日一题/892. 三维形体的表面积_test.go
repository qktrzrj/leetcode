package test17_16

import "testing"

// 在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。
//
//每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。
//
//请你返回最终形体的表面积。
//
//
//
//示例 1：
//
//输入：[[2]]
//输出：10
//示例 2：
//
//输入：[[1,2],[3,4]]
//输出：34
//示例 3：
//
//输入：[[1,0],[0,2]]
//输出：16
//示例 4：
//
//输入：[[1,1,1],[1,0,1],[1,1,1]]
//输出：32
//示例 5：
//
//输入：[[2,2,2],[2,1,2],[2,2,2]]
//输出：46
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/surface-area-of-3d-shapes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// 求解最终形体的表面积，可以分解为，求解所有网格立方体的表面积之和
// 网格表面积解法：
// 设要计算的网格位置为i,j,高度为v
// 上方格子位置为i,j-1,高度为vn
// 左侧格子位置为i-1,j,高度为vw
// 右侧格子位置为i+1,j,高度为ve
// 下方格子位置为i,j+1,高度为vs
// 当周边格子超出N*N范围时，高度为0
// 设函数count(a,b) = a-b>0?a-b:0
// 则，上方位表面积为count(v-vn)
// 以此类推,格子表面积为sum(count(v-vn|vw|ve|vs)) (未计算上下表面积)
func surfaceArea(grid [][]int) int {

	N := len(grid)

	gv := func(v, i, j int) int {
		if i < 0 || j < 0 || i > N-1 || j > N-1 {
			return v
		}
		vv := grid[i][j]
		c := v - vv
		if c < 0 {
			return 0
		}
		return c
	}

	var sum int

	for i, vi := range grid {
		for j, v := range vi {
			if v > 0 {
				sum += gv(v, i-1, j) + gv(v, i+1, j) + gv(v, i, j-1) + gv(v, i, j+1)
				sum += 2
			}
		}
	}
	return sum
}

func Test_surfaceArea(t *testing.T) {
	surfaceArea([][]int{{1, 2}, {3, 4}})
}
