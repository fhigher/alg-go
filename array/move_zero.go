package array

func moveZero(arr []int) {
	zeroNums := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroNums++
		} else if zeroNums > 0 {
			arr[i-zeroNums], arr[i] = arr[i], arr[i-zeroNums]
		}
	}
} 

func removeElem(nums []int, val int) int {
	index := 0 
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}

	return index 
}

// nums increasing order
func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i - 1] {
			nums[index] = nums[i]
			index ++
		}
	}

	return index
}

func removeDuplicates2(nums []int) int {
	index := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i - 1] {
			count += 1
		}
		if nums[i] == nums[i - 1] && count == 2 {
			nums[index] = nums[i]
			index ++
			continue
		}
		if nums[i] > nums[i - 1]  {
			nums[index] = nums[i]
			index ++
			count = 1
		} 
	}

	return index
}

/* func sortColors(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			nums[i-1], nums[i] = nums[i], nums[i-1]
		}
	}
} */