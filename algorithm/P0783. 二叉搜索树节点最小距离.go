/**
 * Author: Deean
 * Date: 2022/11/13 16:17
 * FileName: algorithm/P0783. 二叉搜索树节点最小距离.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func minDiffInBST(root *lib.TreeNode) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	values := []int{}
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		values = append(values, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	minDiff := 100001
	for i := 1; i < len(values); i++ {
		minDiff = min(minDiff, values[i]-values[i-1])
	}
	return minDiff
}

func main() {
	fmt.Println(minDiffInBST(lib.Str2Tree("[1,0,48,null,null,12,49]")))
}
