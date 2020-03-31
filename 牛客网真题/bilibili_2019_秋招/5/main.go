package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

//金闪闪死后，红A拿到了王之财宝，里面有n个武器，长度各不相同。红A发现，拿其中三件武器首尾相接，组成一个三角形，进行召唤仪式，
// 就可以召唤出一个山寨金闪闪。（例如，三件武器长度为10、15、20，可以召唤成功。若长度为10、11、30，首尾相接无法组成三角形，
// 召唤失败。）红A于是开了一个金闪闪专卖店。
// 他把王之财宝排成一排，每个客人会随机抽取到一个区间[l,r],客人可以选取区间里的三件武器进行召唤
// （客人都很聪慧，如果能找出来合适的武器，一定不会放过）。召唤结束后，客人要把武器原样放回去。m个客人光顾以后，
// 红A害怕过多的金闪闪愉悦太多男人，于是找到了你，希望你帮他统计出有多少山寨金闪闪被召唤出来。
//
//
//输入描述:
//第一行武器数量:n <= 1*10^7
//第二行空格分隔的n个int，表示每件武器的长度。
//第三行顾客数量：m <= 1*10^6
//后面m行，每行两个int l，r，表示每个客人被分配到的区间。（l<r）
//
//输出描述:
//山寨金闪闪数量。
//
//输入例子1:
//5
//1 10 100 95 101
//4
//1 3
//2 4
//2 5
//3 5
//
//输出例子1:
//3
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	n, _ := strconv.Atoi(string(line))
	line, _, _ = reader.ReadLine()
	split := strings.Split(string(line), " ")
	wq := make([]int, n)
	for i, e := range split {
		en, _ := strconv.Atoi(e)
		wq[i] = en
	}
	line, _, _ = reader.ReadLine()
	m, _ := strconv.Atoi(string(line))
	for i := 0; i < m; i++ {
		line, _, _ = reader.ReadLine()
		split := strings.Split(string(line), " ")
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		wqCp := wq[l:r]
		sort.Ints(wqCp)
		if wqCp[0]+wqCp[1] < wqCp[2] {

		}
	}
}
