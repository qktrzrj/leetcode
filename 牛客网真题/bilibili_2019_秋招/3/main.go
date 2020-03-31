package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

// 给定一个合法的表达式字符串，其中只包含非负整数、加法、减法以及乘法符号（不会有括号），例如7+3*4*5+2+4-3-1，请写程序计算该表达式的结果并输出；
//
//
//输入描述:
//输入有多行，每行是一个表达式，输入以END作为结束
//
//输出描述:
//每行表达式的计算结果
//
//输入例子1:
//7+3*4*5+2+4-3-1
//2-3*1
//END
//
//输出例子1:
//69
//-1
func main() {
	var s []string
	var stack []interface{}
	var input string
	reader := bufio.NewReader(os.Stdin)
	for {
		readLine, _, _ := reader.ReadLine()
		line := string(readLine)
		input += line + "\n"
		if line == "END" {
			break
		}
	}
	scan := &scanner.Scanner{
		Mode: scanner.ScanInts | scanner.ScanStrings,
	}
	scan.Init(strings.NewReader(input))
	line := 1
	scanFun := func() {
		for {
			scan.Scan()
			if scan.Line != line {
				var result int
				for i := 0; i < len(stack); {
					switch item := stack[i].(type) {
					case int:
						result = item
						i++
					case string:
						if item == "+" {
							result += stack[i+1].(int)
						} else {
							result -= stack[i+1].(int)
						}
						i += 2
					}
				}
				s = append(s, fmt.Sprintf("%d", result))
				line = scan.Line
			}
			break
		}
	}
	scanFun()
	for text := scan.TokenText(); text != "END"; text = scan.TokenText() {
		switch text {
		case "+", "-":
			stack = append(stack, text)
		case "*":
			pos := len(stack) - 1
			prev := stack[pos].(int)
			scanFun()
			next, _ := strconv.Atoi(scan.TokenText())
			stack[pos] = prev * next
		default:
			num, _ := strconv.Atoi(text)
			stack = append(stack, num)
		}
		scanFun()
	}
	fmt.Println(strings.Join(s, "\n"))
}
