package main

import (
    "fmt"
)
func threeEqualParts(A []int) []int {
    sum := 0
    result := make([]int, 2)
    right := make([]int, 2)
    result[0] = -1
    result[1] = -1

    // 找到1的个数
    for i := range A {
        if A[i] == 1 {
            sum ++
        }
    }

    if sum == 0 {
        result[0] = 0
        result[1] = len(A) - 1
        return result
    }
    // 1的个数不为三的倍数 错误
    if sum % 3 != 0 {
        return result
    }

    // 1的个数
    num := sum / 3
    // t 尾部连续0的个数
    t := 0

    for i := len(A) - 1 ; i > 0 ; i-- {
        if A[i] == 0 {
            t++
        }else{
            break
        }
    }

    p := 0
    q := 0
    z := 0

    // 状态
    flag := 0
    n := 0 // 暂存 1个数
    m := 0 // 暂存 0个数

    for i := 0 ; i < len(A) ; i++ {
        // 第一段
        if flag == 0 {
            if A[i] == 1 {
                n++;
                p = p * 2 + 1;
            }else{
                p = p * 2
                if n == num {
                    m++
                }
            }
            if n > num {
                return result
            }
            if n == num && m == t {
                //fmt.Println(i)
                right[0] = i
                flag = 1
                n = 0
                m = 0
            }
        }else if flag == 1 {
            if A[i] == 1 {
                n++;
                q = q * 2 + 1;
            }else{
                q = q * 2
                if n == num {
                    m++
                }
            }
            if n > num {
                return result
            }
            if n == num && m == t {
                //fmt.Println(i)
                right[1] = i + 1
                flag = 2
                n = 0
                m = 0
            }
        }else if flag == 2 {
            //fmt.Println(A[i],n,m)
            if A[i] == 1 {
                n++;
                z = z * 2 + 1;
            }else{
                z = z * 2
                if n == num {
                    m++
                }
            }
            if n > num {
                return result
            }
            if n == num && m == t {
                flag = 3
                n = 0
                m = 0
            }
        }else{
            return result
        }
    }
    //fmt.Println(p,q,z)
    if p == q && p == z {
        return right
    }else {
        return result
    }
}
