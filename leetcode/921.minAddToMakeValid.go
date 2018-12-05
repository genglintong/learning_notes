import (
    "container/list"
)

func minAddToMakeValid(S string) int {
    l := list.New()
    
    for _,v := range S {
        // 栈为空 直接入栈
        if l.Back() == nil {
            l.PushBack(v)
            continue
        }
        if v == ')' && l.Back().Value == '(' {
            l.Remove(l.Back())
        }else {
            l.PushBack(v)
        }
    }
    
    return l.Len()
}
