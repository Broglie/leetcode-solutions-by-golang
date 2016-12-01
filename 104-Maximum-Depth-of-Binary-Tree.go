/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // 方法一：递归，虽然简单但效率教低
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // 我又忍不住吐槽Go的内置类型之间的转换，太坑爹，看下面这行神奇的代码，Max函数参数必须是float64类型
    return int(float64(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))) + 1
}

// 方法二：迭代（BFS），效率比上面那个不知道高到哪里去了
func maxDepth(root *TreeNode) int {
    ret := 0
    if root == nil {
        return 0
    }
    // 用数组充当队列
    var queue [](*TreeNode)
    queue = append(queue, root)
    for len(queue) > 0 {
        ret++
        // 下面这个for循环是层次遍历
        for range queue {
            p := queue[0]
            queue = queue[1:]
            if p.Left != nil {
                queue = append(queue, p.Left)
            }
            if p.Right != nil {
                queue = append(queue, p.Right)
            }
        }
    }
    return ret
}