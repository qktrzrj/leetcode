package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//小A参加了一个n人的活动，每个人都有一个唯一编号i(i>=0 & i<n)，其中m对相互认识，在活动中两个人可以通过互相都认识的一个人介绍认识。现在问活动结束后，小A最多会认识多少人？
//
//
//输入描述:
//第一行聚会的人数：n（n>=3 & n<10000）；
//第二行小A的编号: ai（ai >= 0 & ai < n)；
//第三互相认识的数目: m（m>=1 & m
//< n(n-1)/2）；
//第4到m+3行为互相认识的对，以','分割的编号。
//
//输出描述:
//输出小A最多会新认识的多少人？
//
//输入例子1:
//7
//5
//6
//1,0
//3,1
//4,1
//5,3
//6,1
//6,5
//
//输出例子1:
//3
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	n, _ := strconv.Atoi(string(line))
	line, _, _ = reader.ReadLine()
	ai, _ := strconv.Atoi(string(line))
	line, _, _ = reader.ReadLine()
	m, _ := strconv.Atoi(string(line))
	type node struct {
		i     int
		f     bool
		nodes []*node
	}
	g := make([]*node, n)
	for i := 0; i < n; i++ {
		g[i] = &node{i, false, nil}
	}
	for i := 0; i < m; i++ {
		line, _, _ = reader.ReadLine()
		split := strings.Split(string(line), ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		g[x].nodes = append(g[x].nodes, g[y])
		g[y].nodes = append(g[y].nodes, g[x])
	}
	var friend int
	g[ai].f = true
	queue := []*node{g[ai]}
	for len(queue) > 0 {
		p := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for i := 0; i < len(p.nodes); i++ {
			node := p.nodes[i]
			if node.f {
				continue
			}
			node.f = true
			friend++
			queue = append(queue, node)
		}
	}
	fmt.Println(friend - len(g[ai].nodes))
}
