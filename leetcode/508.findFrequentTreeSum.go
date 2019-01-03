package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    rec := make(map[int] int)
    sum(root, rec)
    res := make([]int, 0, len(rec))
    max := 0

    for i, v := range rec {
        if max <= v {
            if max < v {
                 max = v
                 res = res[0:0]
             }
            res = append(res, i)
        }
    }
    return res
}

// 遍历求和
func sum(root *TreeNode, rec map[int]int) int {
    if root == nil{
        return 0
    }

    result := root.Val + sum(root.Left, rec) + sum(root.Right, rec);

    rec[result] ++

    return result
}
