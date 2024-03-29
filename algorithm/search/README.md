[TOC]

# 二分查找

### 可以用递归实现
子问题就是，在一定范围内查找目标，范围越缩越小

## 二分查找模板

```golang
int binarySearch(int[] nums, int target) {
    int left = 0; 
    int right = nums.length - 1; // 注意

    while(left <= right) {
        int mid = left + (right - left) / 2;
        if(nums[mid] == target)
            return mid; 
        else if (nums[mid] < target)
            left = mid + 1; // 注意
        else if (nums[mid] > target)
            right = mid - 1; // 注意
    }
    return -1;
}
```

# 二叉搜索树

#####平衡因子
每个节点都具有一个平衡因子（balance factor），它表示该节点的`左子树的高度 - 右子树的高度`

二叉搜索树（binary sorted tree），是一种特殊的二叉树结构，其中每个节点的左子树中的所有节点的值都小于该节点的值，而右子树中的所有节点的值都大于该节点的值。（默认 bst 内部没有重复节点）
> 左小右大

因此二叉查找树，中序遍历是递增的

# AVL 树

这个要掌握，它是最基本的平衡查找树，但可能会退化

二叉搜索树，AVL 树的平衡因子只能是 -1、0 或 1

AVL 树的优点在于它可以保证在最坏情况下的查找、插入和删除的时间复杂度为 O(log n)，因此它常常用于对时间复杂度有要求的场景中，比如数据库索引等。

> 相比红黑树，查询性能差不多，但是 AVL 自平衡的开销比较大，适合读多写少





# 红黑树

更多的是了解概念，一般不考察代码
把最难的搞懂，简单的东西理解起来可能就更容易

* 一种平衡的二叉搜索树，保证了查找、插入和删除操作的时间复杂度为O(log n)。
* 适合写多读少
* 红黑树广泛应用于各种数据结构和算法的实现中，比如 Linux 的内核调度、C++ STL 的 set 和 map 等。

### 5 大特性
- 每个节点是红色或黑色。
- 根节点是黑色。
- 每个叶子节点都是黑色的空节点（NIL节点）。
- 如果一个节点是红色的，则它的子节点必须是黑色的。
- （最重要）从任意一个节点到其每个叶子节点的所有路径都包含相同数目的黑色节点。

> 红黑两色，根黑，叶子黑
红节点的儿子黑
到达每个叶子节点的路径，包含的黑节点数相同

###crud
- 插入节点
仍然是按二叉搜索树的特性，先插入到有序的位置，再调整
插入的这个节点默认红色
调整有 5 种情况
通过对新节点的父节点、祖父节点和叔叔节点的颜色和位置进行不同的变换来恢复平衡

- 删除节点
按二叉搜索树删除节点的步骤先执行
调整，有 4 种情况
通过对待删除节点的兄弟节点和侄子节点的颜色和位置进行不同的变换来实现，使得红黑树恢复平衡

# 跳表

> 红黑树的平替

一种用于实现有序集合的数据结构，它基于链表的基础上增加了多层索引结构，以提高查询效率

空间换时间