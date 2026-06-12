package base

func Mono(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	isInc, isDec := true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			isInc = false
		} else {
			isDec = false
		}

		if !isInc && !isDec {
			return false
		}
	}

	return true
}
