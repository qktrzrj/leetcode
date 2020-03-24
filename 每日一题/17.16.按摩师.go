package main

import (
	"fmt"
)

//一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。
// 在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。
// 给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
//
// 注意：本题相对原题稍作改动
//
// 示例 1：
//
// 输入： [1,2,3,1]
// 输出： 4
// 解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
//
// 示例 2：
//
// 输入： [2,7,9,3,1]
// 输出： 12
// 解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。
//
// 示例 3：
//
// 输入： [2,1,4,5,3,1,1,3]
// 输出： 12
// 解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/the-masseuse-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 求解: 动态规划求最优解问题

// 递归解法+制表法
//
// 假设有n个客户，nums[i]代表第i+1个客户的预约时长，由于不能接收相邻的预约，那么当按摩师接收第k个客户的预约，则不能接受第k+1个客户的预约。
// 推导出公式为：sum(k) = nums[k+1] + max(sum(k+2),sum(k+3))
// 按照最优解要求，按摩师必定接受第一个客户或者第二个客户的预约
// 最终公式为：s = max(sum(0),sum(1))
func massage1(nums []int) int {
	table := make([]int, len(nums))
	for i := range table {
		table[i] = -1
	}
	var sum func(int) int
	sum = func(pos int) int {
		if pos >= len(nums) {
			return 0
		}
		if table[pos] != -1 {
			return table[pos]
		}
		table[pos] = nums[pos] + max(sum(pos+2), sum(pos+3))
		return table[pos]
	}

	return max(sum(0), sum(1))
}

// 迭代法
//
// 在此题中，按摩师有两个选择，工作或者休息，同时按摩师不能连续工作。
// 那么可以定义work,rest两个变量表示这一次选择工作或者休息的结果
// 如果第i次选择工作，那么上一次的结果肯定是休息，所以 work = rest + i
// 如果第i次选择休息，那么取决于上一次是否休息,如果上一次工作则 rest = work,如果上一次休息则 rest = rest
// 当第i次选择休息时,rest为第i次的最优解,故而rest=max(work,rest)
// 得出推导公式rest(k) =max(work(k-1),rest(k-1)), work(k-1) = rest(k-2) + nums[k-1]
// 那么这个问题就可以从末尾开始迭代，求解n的最优解，可以理解为n+1为休息,那么最优解 = max(work(n),rest(n))
func massage2(nums []int) int {
	var work, rest int
	for _, v := range nums {
		work, rest = rest+v, max(work, rest)
	}
	return max(work, rest)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("no.1: %d\n", massage2([]int{1, 2, 3, 1}))
	fmt.Printf("no.2: %d\n", massage2([]int{2, 7, 9, 3, 1}))
	fmt.Printf("no.3: %d\n", massage1([]int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124}))
}
