package main

func uniquePaths(m int, n int) int {
    // C(m + n - 2)(n-1)
    m = m-1
    n = n-1

    // total 总步数
    // n 较少步数
    total := m + n
    if m < n {
        n = m
    }
    num := 1
    m = 1

    for i := 1 ; i <= n ; i++ {
        m = m * i
        num = num * (total + 1 - i)
    }

    return num/m
}
