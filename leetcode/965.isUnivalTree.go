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
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return true
    } else {
        if root.Left != nil && root.Left.Val != root.Val {
            return false
        }
        if root.Right != nil && root.Right.Val != root.Val {
            return false
        }
        if isUnivalTree(root.Left) == true && isUnivalTree(root.Right) == true {
            return true
        } else {
            return false
        }
    }
