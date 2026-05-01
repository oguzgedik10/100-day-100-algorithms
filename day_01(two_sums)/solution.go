package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		diff := target - n
		if idx, ok := m[diff]; ok {
			return []int{idx, i}
		}
		m[n] = i
	}
	return nil
}
