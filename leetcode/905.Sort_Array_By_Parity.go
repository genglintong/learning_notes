package main

func sortArrayByParity(A []int) []int {
    j := len(A) - 1
    i := 0

    // 快排思路
    for i < j {
        // 判断奇偶
        for (A[i] & 1 == 0) && (i < j) {
            i++
        }
        for (A[j] & 1 == 1) && (i < j) {
            j--
        }
        if i < j {
            // 交换元素
            A[i] , A[j] = A[j] , A[i]
        }
    }

    return A
}
