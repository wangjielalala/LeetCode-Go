package leetcode

import (
	"sort"
)

// 解法一 最优解，双指针 + 排序
func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		// 这里采用两个头节点与一个尾节点相加，两个最小的节点与一个最大的节点
		start, end = 0, length-1
		// 这里因为index是顺序递增的，所以如果说本次的index 和 index-1 想等，则意味着index-1 前面的枚举场景已经遍历完成，无需再遍历一次，防止出现重复遍历的情况，影响性能以及重复值
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			// start相同的话，则略过
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			// end相同的话，则略过
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			// 三值相加
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				// 想等，则同时缩小尾脂针、增加头指针
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				// 缩小尾脂针
				end--
			} else {
				// 增加头指针
				start++
			}
		}
	}
	return result
}

// 解法二
func threeSum1(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}
