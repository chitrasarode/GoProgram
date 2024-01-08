package bubbleSort

func MyBubbleSort(f func([]int, int) []int, arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			f(arr, j)
		}
	}
	return arr
}

func Ascending(arr []int, x int) []int {
	if arr[x] > arr[x+1] {
		arr[x], arr[x+1] = arr[x+1], arr[x]
	}
	return arr
}

func Descending(arr []int, x int) []int {
	if arr[x] < arr[x+1] {
		arr[x], arr[x+1] = arr[x+1], arr[x]
	}
	return arr
}
