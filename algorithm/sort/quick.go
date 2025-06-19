package sort

// quick sort
// 以数组中的某个元素为基准（比如中间元素），将数组分成两个子数组，左边的子数组的所有元素都小于基准元素，右边的子数组的所有元素都大于基准元素，然后递归地重复这个过程直到排序完成。
// 时间复杂度：O(nlogn)
func QuickSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	// 基准元素
	pivot := nums[0]

	leftNums := []int{}
	rightNums := []int{}
	for i := 1; i < l; i++ {
		if nums[i] < pivot {
			leftNums = append(leftNums, nums[i])
		} else {
			rightNums = append(rightNums, nums[i])
		}
	}
	leftNums = QuickSort(leftNums)
	leftNums = append(leftNums, pivot)
	return append(leftNums, QuickSort(rightNums)...)
}
