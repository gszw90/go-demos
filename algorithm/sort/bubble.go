package sort

// BubbleSort 冒泡排序
// 遍历所有的数据，每次对相邻元素进行两两比较，
// 如果顺序和预先规定的顺序不一致，则进行位置交换；
// 这样一次遍历会将最大或最小的数据上浮到顶端，
// 之后再重复同样的操作，直到所有的数据有序。
//
// 时间复杂度：O(n^2)
func BubbleSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	for i := 0; i < l; i++ {
		// 每次循环都将最大的数放到最后面
		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
