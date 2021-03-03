// Leetcode 15. 三数之和
func threeSum(nums []int) [][]int {
    n := len(nums)
    if n < 3 {
        return nil
    }
    var ans [][]int
    quickSort(nums, 0, n - 1)

    l, r := 0, 0
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            return ans
        }
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        l = i + 1
        r = n - 1

        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0 {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                l++
                r--
            } else if sum < 0 {
                l++
            } else {                
                r--
            }
        } 
    }
    return ans
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

    for ; j < end; j++ {
        if nums[j] < nums[end] {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }

    nums[i], nums[end] = nums[end], nums[i]
    return i
}