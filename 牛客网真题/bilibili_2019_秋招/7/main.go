package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 猛兽侠中精灵鼠在利剑飞船的追逐下逃到一个n*n的建筑群中，精灵鼠从（0,0）的位置进入建筑群，
// 建筑群的出口位置为（n-1,n-1），建筑群的每个位置都有阻碍，每个位置上都会相当于给了精灵鼠一个固定值减速，
// 因为精灵鼠正在逃命所以不能回头只能向前或者向下逃跑，现在问精灵鼠最少在减速多少的情况下逃出迷宫？
//
//
//输入描述:
//第一行迷宫的大小: n >=2 & n <= 10000；
//第2到n+1行，每行输入为以','分割的该位置的减速,减速f >=1 & f < 10。
//
//输出描述:
//精灵鼠从入口到出口的最少减少速度？
//
//输入例子1:
//3
//5,5,7
//6,7,8
//2,2,4
//
//输出例子1:
//19
type point struct {
	x    int
	y    int
	dist int
	w    int
	in   bool
	prev *point
}

type queue []*point

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].w > q[j].w
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	n, _ := strconv.Atoi(string(line))
	dir := [][2]int{{1, 0}, {0, 1}}
	set := make([][]*point, n)
	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		line, _, _ = reader.ReadLine()
		split := strings.Split(string(line), ",")
		one, _ := strconv.Atoi(split[0])
		two, _ := strconv.Atoi(split[1])
		three, _ := strconv.Atoi(split[2])
		set[i] = append(set[i], &point{x: i, y: 0, w: one, dist: 10 * n * n}, &point{x: i, y: 1, w: two, dist: 10 * n * n}, &point{x: i, y: 2, w: three, dist: 10 * n * n})
	}
	set[0][0].dist = 0
	queue := queue{set[0][0]}
	for len(queue) > 0 {
		sort.Sort(queue)
		min := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if min.x == n-1 && min.y == n-1 {
			break
		}
		for i := 0; i < 2; i++ {
			dx := min.x + dir[i][0]
			dy := min.y + dir[i][1]
			if dx < 0 || dx >= n || dy < 0 || dy >= n {
				continue
			}
			next := set[dx][dy]
			if next.dist > min.dist+next.w {
				next.dist = min.dist + next.w
				next.prev = min
				if !next.in {
					queue = append(queue, next)
					next.in = true
				}
			}
		}
	}
}
