package main

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for idxFirst, firstNum := range nums {
		secondNum := target - firstNum
		if idxSecond, exist := hashMap[secondNum]; exist {
			return []int{idxSecond, idxFirst}
		}
		hashMap[firstNum] = idxFirst
	}
	return []int{}
}
