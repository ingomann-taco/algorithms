package util

// BinarySearch Finds a target in a slice using binary search
func BinarySearch(slice []int, target int) (int, bool) {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == target {
			return mid, true
		} else if slice[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, false
}
