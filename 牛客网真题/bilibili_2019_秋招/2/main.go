package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// av394281 中，充满威严的蕾米莉亚大小姐因为触犯某条禁忌，被隙间妖怪八云紫（紫m……èi）按住头在键盘上滚动。
//同样在弹幕里乱刷梗被紫姐姐做成罪袋的你被指派找到大小姐脸滚键盘打出的一行字中的第 `k` 个仅出现一次的字。
//(为简化问题，大小姐没有滚出 ascii 字符集以外的字)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s []string
	for line, _, err := reader.ReadLine(); err == nil; line, _, err = reader.ReadLine() {
		if len(line) == 0 {
			break
		}
		found := 0
		k, _ := strconv.Atoi(string(line[:1]))
		str := line[2:]
		for i := 0; i < len(str); i++ {
			if strings.LastIndexByte(string(str), str[i]) == strings.IndexByte(string(str), str[i]) {
				found++
				if found == k {
					s = append(s, "["+string(str[i])+"]")
					break
				}
			}
		}
		if found == 0 {
			s = append(s, "Myon~")
		}
	}
	fmt.Println(strings.Join(s, "\n"))
}
