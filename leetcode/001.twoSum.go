package main

func twoSum(nums []int, target int) []int {
    // index 负责保存整数map
    index := make(map[int]int, len(nums))

    for i , b :=  range nums {
        if j, ok := index[target - b] ; ok {
            // map存在 返回
            return []int{j , i}
        }
        // map 没有 更新map
        index[b] = i
    }

    // 不存在
    return nil
}
