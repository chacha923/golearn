package binary_tree

// 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
func levelOrderNTree(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*Node{root}
	var level int
	// 队列清空时操作结束
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{}) // 每层一个数组
		for i := 0; i < counter; i++ {
			if queue[i] != nil {
				res[level] = append(res[level], queue[i].Val)
				for _, n := range queue[i].Children {
					queue = append(queue, n) // 每个孩子进队列
				}
			}
		}
		queue = queue[counter:] // 这一层节点扔掉
		level++
	}
	return res
}
