package main

func numJewelsInStones(J string, S string) int {
    // 初始化map
    isJewel := make(map[byte]bool , len(J))

    // 更新map
    for i := range J {
        isJewel[J[i]] = true
    }

    // 初始化
    res := 0
    for i := range S {
        if isJewel[S[i]] {
            res++
        }
    }

    return res
}
