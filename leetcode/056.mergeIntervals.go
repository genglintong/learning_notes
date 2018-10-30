/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
import(
    "sort"
)

func merge(intervals []Interval) []Interval {
    // 返回本身
    if len(intervals) <= 1 {
        return intervals
    }
    
    // 首先按照第一个元素排序
    sort.Slice(intervals, func(i int, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    
    res := make([]Interval, 0, len(intervals))
    temp := intervals[0]
    
    for i := 1 ; i < len(intervals) ; i++ {
        // 存在交集
        if intervals[i].Start <= temp.End {
            temp.End = max(temp.End, intervals[i].End)
        } else {
            res = append(res, temp)
            temp = intervals[i]
        }
    }
    res = append(res, temp)
    
    return res
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
