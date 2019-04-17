func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, v := range nums {
		dict[v] = i
	}
	for i, v := range nums {
		if v1, ok := dict[target - v]; ok {
            if i == v1 {
				continue
			}
			return []int{i, v1}
		}
	}
	return nil
}