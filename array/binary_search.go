package array

// 二分查找，迭代方式，O(logn)
func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for high >= low {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func binarySearchRecursive(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	return searchRecursive(nums, low, high, target)
}

func searchRecursive(nums []int, low, high int, target int) int {
	index := -1
	if low >= high {
		return index
	}
	mid := low + (high-low)/2
	
	if nums[mid] == target {
		index = mid
	} else if target < nums[mid] {
		index = searchRecursive(nums, low, mid-1, target)
	} else {
		index = searchRecursive(nums, mid+1, high, target)
	}

	return index
}
