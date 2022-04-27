package two_sum

//[2,7,11,15] 9
// 0,1
func TwoSum(nums []int, target int) []int {

	data := make(map[int]int)

	for i, num := range nums {
		op := target - num
		if _, ok := data[op]; ok {
			return []int{data[op], i}
		}

		data[num] = i
	}
	return nil
}