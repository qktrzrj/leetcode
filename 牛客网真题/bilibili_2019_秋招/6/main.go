package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 如果version1 > version2 返回1，如果 version1 < version2 返回-1，不然返回0.
//
//输入的version字符串非空，只包含数字和字符.。
// .字符不代表通常意义上的小数点，只是用来区分数字序列。
// 例如字符串2.5并不代表二点五，只是代表版本是第一级版本号是2，第二级版本号是5.
//
//
//输入描述:
//两个字符串，用空格分割。
//每个字符串为一个version字符串，非空，只包含数字和字符.
//
//输出描述:
//只能输出1, -1，或0
//
//输入例子1:
//0.1 1.1
//
//输出例子1:
//-1
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	split := strings.Split(string(line), " ")
	var v1, v2 int
	for i, s := range split {
		ss := strings.Split(s, ".")
		i1, _ := strconv.Atoi(ss[0])
		i2, _ := strconv.Atoi(ss[1])
		pow := int(math.Pow10(i))
		v1 = v1*pow + i1
		v2 = v2*pow + i2
	}
	if v1 > v2 {
		fmt.Println(1)
	} else if v1 < v2 {
		fmt.Println(-1)
	} else {
		fmt.Println(0)
	}
}
