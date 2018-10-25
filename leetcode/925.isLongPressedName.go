package main
import (
    "fmt"
)
func isLongPressedName(name string, typed string) bool {
    n := []byte(name)
    t := []byte(typed)

    i := 0
    for j := 0 ; j < len(t) && i < len(n) ; j++ {
        if n[i] == t[j] {
            i++;
        }
    }

    if i == len(n) {
        return true
    }else {
        return false
    }
}
