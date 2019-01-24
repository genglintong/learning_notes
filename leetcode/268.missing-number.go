package main

import "fmt"

/**
* 与041 解题思路一致
*/
func missingNumber(nums []int) int {
    for i := 0 ; i < len(nums) ; {
        if nums[i] < len(nums) && nums[i] != i {
            nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
        }else {
            i++
        }
    }

    for i:=0 ; i < len(nums) ; i++ {
        if (nums[i] != i){
            return i
        }
    }

    return len(nums)
}
