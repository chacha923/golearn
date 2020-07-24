package structure

// https://leetcode-cn.com/problems/implement-trie-prefix-tree/solution/trie-tree-de-shi-xian-gua-he-chu-xue-zhe-by-huwt/
// 字典树
type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Char    rune
	HasData bool
	Data    interface{} //存放得数据
	Child   map[rune]*TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(' '),
	}
}

func NewTrieNode(char rune) *TrieNode {
	node := TrieNode{
		Char:    char,
		HasData: false,
		Data:    "",
		Child:   make(map[rune]*TrieNode),
	}
	return &node
}

func (t *Trie) AddNode(key string, data interface{}) {
	parent := t.Root
	keyRune := []rune(key)
	for _, v := range keyRune {
		//判断是否存在此节点
		node, ok := parent.Child[v]
		if !ok {
			//如果不存在则构建节点
			node := &TrieNode{
				Char: v,
			}
			//加入到父节点得子节点
			parent.Child[v] = node
		}
		//将父节点改成当前找到（或者添加）得节点
		parent = node
	}
	//给最后得节点添加数据
	parent.Data = data
	parent.HasData = true
}

func (t *Trie) SearchNode(key string, limit int) (res []interface{}) {
	keyRune := []rune(key)
	parent := t.Root
	//查找到key所能找到得节点
	for _, v := range keyRune {
		node, ok := parent.Child[v]
		if !ok {
			return
		}
		parent = node
	}
	//将获取到得节点与其后代节点全部找出来
	var queue = make([]*TrieNode, 0, len(keyRune)) //定义一个当前找到得节点或者子节点得切片
	queue = append(queue, parent)
	for len(queue) > 0 {
		var childQueue []*TrieNode
		for _, v := range queue {
			if v.HasData {
				res = append(res, v.Data)
			}
			if len(res) == limit {
				return
			}
			for _, vi := range v.Child {
				childQueue = append(childQueue, vi)
			}
		}
	}
}
