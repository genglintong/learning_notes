package main

import "fmt"

func rob(nums []int) int {
    /**
    * 简单动态规划问题
    * dp[n] = max (dp[n-1], dp[n-2] + nums[n])
    */
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    
    // 用两个变量 代替数组 换时间
    dp0 := 0
    dp := nums[0]
    
    for i := 1 ; i < len(nums) ; i++ {
        dp0 = max(dp, dp0 + nums[i])
        dp , dp0 = dp0 , dp
    }
    
    return dp
}

func max(a , b int) int {
    if a < b {
        return b
    }
    return a
}
