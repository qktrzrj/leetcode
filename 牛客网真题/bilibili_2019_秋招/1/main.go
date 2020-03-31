package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 22娘和33娘接到了小电视君的扭蛋任务：
//一共有两台扭蛋机，编号分别为扭蛋机2号和扭蛋机3号，22娘使用扭蛋机2号，33娘使用扭蛋机3号。
//扭蛋机都不需要投币，但有一项特殊能力：
//扭蛋机2号：如果塞x（x范围为>=0正整数）个扭蛋进去，然后就可以扭到2x+1个
//扭蛋机3号：如果塞x（x范围为>=0正整数）个扭蛋进去，然后就可以扭到2x+2个
//22娘和33娘手中没有扭蛋，需要你帮她们设计一个方案，两人“轮流扭”（谁先开始不限，扭到的蛋可以交给对方使用），用“最少”的次数，使她们能够最后恰好扭到N个交给小电视君。

func min(n int) string {
	if n == 0 {
		return ""
	}
	var s string
	for n != 0 {
		if n%2 == 0 {
			s, n = "3"+s, (n-2)/2
		} else {
			s, n = "2"+s, (n-1)/2
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	read, _, _ := reader.ReadLine()
	atoi, _ := strconv.Atoi(string(read))
	fmt.Println(min(atoi))
}
