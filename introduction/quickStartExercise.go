package introduction

//[strStr](https://leetcode-cn.com/problems/implement-strstr/)
//
//> 给定一个  haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回  -1。
//
//思路：核心点遍历给定字符串字符，判断以当前字符开头字符串是否等于目标字符串

//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//1.9 MB
//, 在所有 Go 提交中击败了
//18.38%
//的用户
//通过测试用例：
//79 / 79
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen > haystackLen {
		return -1
	}

	getRes := -1
	for startIndex := 0; startIndex <= haystackLen-needleLen; startIndex++ {
		if haystack[startIndex:startIndex+needleLen] == needle {
			getRes = startIndex
			break
		}

	}

	return getRes
}

//> 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//思路：这是一个典型的应用回溯法的题目，简单来说就是穷尽所有可能性，算法模板如下
//func subsets(nums []int) [][]int {
//	res := make([][]int,0)
//
//	numLen := len(nums)
//	if numLen == 1 {
//		return [][]int{{nums[0]}}
//	}
//
//	for i := 0; i < numLen; i++ {
//		count := 1
//		res = append(res, nums[i])
//
//	}
//
//
//}