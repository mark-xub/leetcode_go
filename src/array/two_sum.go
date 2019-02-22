package array

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		except := target - nums[i]
		if val, ok := numMap[except]; ok {
			return []int{val, i}
		}
		numMap[nums[i]] = i
	}

	return nil
}
