package main

func trailingZeroes(n int) int {
    t := 5
    count := 0

    for t <= n {
        count += n / t
        t = t * 5
    }

    return count
}
