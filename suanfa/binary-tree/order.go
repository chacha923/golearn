package binary_tree

import (
	"fmt"
	"github.com/eapache/queue"
)

//遍历

//递归前序遍历
func PreOrder(root *TreeNode){
	if root == nil {
		return
	}
	fmt.Println(root.value)
	PreOrder(root.left)
	PreOrder(root.right)
}

//递归中序遍历
func InOrder(root *TreeNode){
	if root == nil {
		return
	}
	PreOrder(root.left)
	fmt.Println(root.value)
	PreOrder(root.right)
}

//递归后序遍历
func PostOrder(root *TreeNode){
	if root == nil {
		return
	}
	PreOrder(root.left)
	PreOrder(root.right)
	fmt.Println(root.value)
}


//分层遍历
func LevelOrder(root *TreeNode){
	if root != nil {
		return
	}
	queue := queue.New()
	queue.Add(root)
	for queue.Length() != 0 {
		length := queue.Length()
		for i := 0; i < length; i++{
			node := queue.Remove().(*TreeNode)
			fmt.Print(node.value," ")
			if node.left != nil {
				queue.Add(node.left)
			}
			if node.right != nil {
				queue.Add(node.right)
			}
		}
		fmt.Println()
	}
	return
}