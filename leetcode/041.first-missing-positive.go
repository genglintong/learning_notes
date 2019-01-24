package main

import "fmt"

/**
* 首先这个正整数 最大可能是 n+1 (n为数组长度)
*/
func firstMissingPositive(nums []int) int {
    for i := 0 ; i < len(nums) ; {
        if nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i] - 1] != nums[i] {
            nums[nums[i] - 1] , nums[i] = nums[i], nums[nums[i] - 1]
        }else {
            i++
        }
    }
    for i := 0 ; i < len(nums); i++ {
        if (nums[i] != i+1) {
            return i+1
        }
    }
    return len(nums)+1
}
