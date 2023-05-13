package graph

type Node struct {
	adj []*Node
	val int
}

func NewNode(value int) *Node {
	return &Node{val: value}
}

// 无向图节点，添加边
func (n *Node) AddEdge(node *Node) {
	n.adj = append(n.adj, node)
	node.adj = append(node.adj, n)
}
