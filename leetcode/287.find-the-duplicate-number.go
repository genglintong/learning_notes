package main

import "fmt"

/**
* 看评论 使用快慢指针思路
* 2. 使用 遍历 变负数 或者 异或都是可以的 但是需要修改数组
*/
func findDuplicate(nums []int) int {
    fast := 0
    slow := 0

    for true {
        // 快慢遍历
        fast = nums[nums[fast]]
        slow = nums[slow]

        if (slow == fast) {
            fast = 0
            for nums[fast] != nums[slow] {
                fast = nums[fast]
                slow = nums[slow]
            }
            return nums[slow]
        }
    }

    return 0
}
