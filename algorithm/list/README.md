# 链表 数组 线性表

## 双指针
在处理数组和链表相关问题时，双指针技巧是经常用到的，双指针技巧主要分为两类：左右指针和快慢指针。

### 秒杀技巧，6 大
1、合并两个有序链表

2、链表的分解

3、合并 k 个有序链表

4、寻找单链表的倒数第 k 个节点

5、寻找单链表的中点

6、判断单链表是否包含环并找出环起点

7、判断两个单链表是否相交并找出交点




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

## dummy
> 经常有读者问我，什么时候需要用虚拟头结点？我这里总结下：当你需要创造一条新链表的时候，可以使用虚拟头结点简化边界情况的处理。

比如说，让你把两条有序链表合并成一条新的有序链表，是不是要创造一条新链表？再比你想把一条链表分解成两条链表，是不是也在创造新链表？这些情况都可以使用虚拟头结点简化边界情况的处理。

