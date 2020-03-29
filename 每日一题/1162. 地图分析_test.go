package test17_16

import (
	"fmt"
	"testing"
)

//你现在手里有一份大小为 N x N 的『地图』（网格） grid，上面的每个『区域』（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，你知道距离陆地区域最远的海洋区域是是哪一个吗？请返回该海洋区域到离它最近的陆地区域的距离。
//
//我们这里说的距离是『曼哈顿距离』（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 - x1| + |y0 - y1| 。
//
//如果我们的地图上只有陆地或者海洋，请返回 -1。
//
//
//
//示例 1：
//
//
//
//输入：[[1,0,1],[0,0,0],[1,0,1]]
//输出：2
//解释：
//海洋区域 (1, 1) 和所有陆地区域之间的距离都达到最大，最大距离为 2。
//示例 2：
//
//
//
//输入：[[1,0,0],[0,0,0],[0,0,0]]
//输出：4
//解释：
//海洋区域 (2, 2) 和所有陆地区域之间的距离都达到最大，最大距离为 4。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/as-far-from-land-as-possible
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//

type Point struct {
	X int
	Y int
}

func maxDistance(grid [][]int) int {
	var queue []*Point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, &Point{i, j})
			}
		}
	}
	if len(queue) == 0 || len(queue) == len(grid)*len(grid) {
		return -1
	}

	ans := 0
	d := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) > 0 {
		ans++
		tmpQueue := queue
		queue = nil
		for len(tmpQueue) > 0 {
			p := tmpQueue[0]
			tmpQueue = tmpQueue[1:]
			for i := 0; i < 4; i++ {
				x := p.X + d[i][0]
				y := p.Y + d[i][1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != 0 {
					continue
				}
				queue = append(queue, &Point{x, y})
				grid[x][y] = 1
			}
		}
	}

	return ans - 1
}

func Test_maxDistance(t *testing.T) {
	fmt.Println(maxDistance([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}
