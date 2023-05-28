# 链表 数组 线性表

## 秒杀技巧，6 大





## 遍历框架
```golang
// 基本的单链表节点
type ListNode struct {
    val int
    next *ListNode
}

func traverse(head *ListNode) {
    for p := head; p != nil; p = p.next {
        // 迭代访问 p.val
    }
}

func traverseRecursively(head *ListNode) {
    // 递归访问 head.val
    traverseRecursively(head.next)
}
```