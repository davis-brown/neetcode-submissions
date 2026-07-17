func hasDuplicate(nums []int) bool {
    count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if(count[nums[i]] == 1){
			return true;
		}

		count[nums[i]] = 1;
	}

	return false
}
