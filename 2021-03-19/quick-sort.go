// 快速排序(原地)

func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums) - 1)
    
    return nums
}

func quickSort(nums []int, start, end int) {
    if start >= end {
        return
    }

    q := partition(nums, start, end)
    quickSort(nums, start, q - 1)
    quickSort(nums, q + 1, end)
}

func partition(nums []int, start, end int) int {
    i, j := start, start

    for j < end {
        if nums[j] < nums[end] {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
        j++
    }

    nums[i], nums[end] = nums[end], nums[i]

    return i
}