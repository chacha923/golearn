package binarytree

// https://www.zhihu.com/question/41911660 遍历二叉树深度优先和广度优先的通俗解释？

// 层次遍历

// 二叉树最大宽度
// 给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。
// 每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	queue := []Item{{0, root}} // bfs使用的队列
	for len(queue) > 0 {
		l := queue[len(queue)-1].idx - queue[0].idx + 1
		if l > ans {
			ans = l
		}
		tmp := []Item{} // 临时队列, 使得每一次循环, queue内包含同一层的节点(分层遍历?)
		for _, q := range queue {
			if q.Left != nil {
				tmp = append(tmp, Item{q.idx * 2, q.Left})
			}
			if q.Right != nil {
				tmp = append(tmp, Item{q.idx*2 + 1, q.Right})
			}
		}
		queue = tmp
	}
	return ans
}

// 节点指针+下标构成的元组, golang没有元组用结构体代替
type Item struct {
	idx int
	*TreeNode
}
