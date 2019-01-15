package main

import "fmt"

func longestConsecutive(nums []int) int {
    numsMap := make(map[int]int)

    for _, v := range nums {
        // 前一个数 是否在map 中
        pre, preOk := numsMap[v-1]
        if preOk {
            numsMap[v] = pre + 1
        }else {
            numsMap[v] = 1
        }

        // 后一个数 是否在map 中
        t := v + 1
        _, nextOk := numsMap[t]
        for nextOk {
            numsMap[t] = numsMap[t-1] + 1
            t = t + 1
            _, nextOk = numsMap[t]
        }

        _, ok := numsMap[v]
        // 存在在map
        if !ok {
            numsMap[v] = 1
        }
        /*
        for k, v := range numsMap {
            fmt.Printf("%d %d\n",k , v)
        }*/
    }

    max := 0

    for _, v := range numsMap {
        if v > max {
            max = v
        }
    }

    return max
}
