package binarysearch

func BinarySearch(numbers []int, guess int) int {
	min := 0
	max := len(numbers) - 1

	for min <= max {
		middle := (min + max) / 2
		item := numbers[middle]

		if item == guess {
			return middle
		} else if item > guess {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}

	return -1
}
