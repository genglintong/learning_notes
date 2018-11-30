package main
import (
    "fmt"
)
func findKthLargest(nums []int, k int) int {
    len := len(nums)
    
    start := 0
    end := len-1
    t := 0
    k = k - 1
    
    t = quickSort(nums, start, end)
    for t != k {
        if t < k {
            start = t + 1
        }else {
            end = t - 1
        }
        t = quickSort(nums, start, end)
    }
    return nums[t]
}

func quickSort(nums []int, start int, end int) int {
    target := nums[end]
    t := end
    for start < end {
        for start < end && nums[start] > target {
            start++
        }
        for start < end && nums[end] <= target {
            end--
        }
        if start < end {
            nums[start], nums[end] = nums[end], nums[start]
        }
    }
    nums[start] , nums[t] = nums[t], nums[start]
    return start
}
