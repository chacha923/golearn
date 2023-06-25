package skiplist

import (
	"math/rand"
)

// https://mp.weixin.qq.com/s/fvfz6bdvsZJtGsdL0MPYoA  基于golang从零到一实现跳表

var p float64 = 0.5 //控制是否在上层level创建节点的概率, 如果rand < p , 那么创建节点

type SkipList struct {
	level int
	high  int
	size  int

	head *SkipListNode //level 0 头结点
	tail *SkipListNode //level 0 尾节点
}

// 初始化
func NewSkipList() *SkipList {
	sl := &SkipList{}
	sl.head = NewSkipListNode("head", "head") //顶层的头节点
	sl.tail = NewSkipListNode("tail", "tail") //顶层的尾节点
	sl.head.SetRight(sl.tail)
	sl.tail.SetLeft(sl.head)
	sl.high = 0
	sl.size = 0

	return sl
}

// 判断跳表是否空
func (sl *SkipList) IsEmpty() bool {
	if sl.size == 0 {
		return true
	}
	return false
}

func (sl *SkipList) FindFront(key string) *SkipListNode {
	tmp := sl.head
	for {
		for tmp.GetRight().key != "tail" && tmp.GetRight().key <= key {
			tmp = tmp.right
		}
		if tmp.GetDown() != nil {
			tmp = tmp.GetDown()
		} else {
			break
		}
	}

	return tmp
}

/*
*
添加一个新节点
*/
func (sl *SkipList) Add(k, v string) string {
	tmp := sl.FindFront(k)

	if k == tmp.GetKey() { //不允许添加重复的key
		println("对象属性完全相同无法添加!")
		a := tmp.value
		tmp.value = v
		return a
	}
	//tmp1 被插入 tmp 右边
	tmp1 := NewSkipListNode(k, v)
	tmp1.SetLeft(tmp)
	tmp1.SetRight(tmp.GetRight())
	tmp.GetRight().SetLeft(tmp1)
	tmp.SetRight(tmp1)

	i := 0
	for rand.Float64() < p { //随机数 < 0.5 , 拷贝到上一层
		if i >= sl.high { //超出当前层数
			sl.high += 1 //创建新的一层
			p1 := NewSkipListNode("head", "head")
			p2 := NewSkipListNode("tail", "tail")
			p1.SetRight(p2)
			p1.SetDown(sl.head)
			p2.SetLeft(p1)
			p2.SetDown(sl.tail)

			sl.head.SetUp(p1)
			sl.tail.SetUp(p2)

			sl.head = p1 //head指向新一层的头结点
			sl.tail = p2

		}
		for tmp.GetUp() == nil {
			tmp = tmp.GetLeft()
		}
		tmp = tmp.GetUp()
		node := NewSkipListNode(k, v)
		node.SetLeft(tmp)
		node.SetRight(tmp.right)
		node.SetDown(tmp1)

		tmp.GetRight().SetLeft(node)
		tmp.SetRight(node)
		tmp1.SetUp(node)

		tmp1 = node
		i += 1

	}

	sl.size += 1
	return "head"
}

/*
*
查找一个节点

	  从顶层头结点开始遍历, 当右边节点key > k,
		说明该层没有k, 移动到下一层, 继续向右查找
*/
func (sl *SkipList) Find(k string) *SkipListNode {
	tmp := sl.head
	node := tmp
	println("查找路线 ")

	for tmp != nil {
		for node.GetRight().key != "tail" && node.GetRight().GetKey() <= k {
			node = node.GetRight()
			println("----> " + node.GetKey())
		}

		if (node.GetDown()) != nil {
			node = node.GetDown()
			println("----> " + node.GetKey())
		} else {
			if node.key == k {
				println("----> " + node.GetKey())
				println("----> " + node.GetValue())
				return node
			}
			return nil
		}
	}
	return nil
}

/*
*
删除一个节点
调用查找函数，删除最底层的某个节点，并把其节点的左右相连，
和链表操作一样，只是其上方若有则都需要调整
*/
func (sl *SkipList) Del(k string) {
	tmp := sl.Find(k)
	for tmp != nil {
		tmp.GetLeft().SetRight(tmp.GetRight())
		tmp.GetRight().SetLeft(tmp.GetLeft())
		tmp = tmp.GetUp()
	}
}

/*
*
格式化输出跳表
*/
func (sl *SkipList) Print() {
	node1 := sl.head
	var node *SkipListNode
	for node1 != nil {
		k := 0
		node = node1
		for node != nil {
			println(node.GetKey() + "\t")
			k++
			node = node.GetRight()
		}
		println("\t")
		println("(" + string(k) + ")")
		println()
		node1 = node1.GetDown()
	}
}
