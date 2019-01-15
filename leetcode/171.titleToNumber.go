package main

import "fmt"

/**
* 进制转换 问题
*/
func titleToNumber(s string) int {
    res := 0
    
    for _, v := range s {
        temp := int(v - 'A' + 1)
        res = res*26 + temp
    }
    
    return res
}
