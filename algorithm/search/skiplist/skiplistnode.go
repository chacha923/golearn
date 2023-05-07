package skiplist

type SkipListNode struct {
	key                   string
	value                 string
	up, down, left, right *SkipListNode

	next []SkipListNode		//保存node节点所在level, 排在node之后的节点
}

func NewSkipListNode(key, value string) *SkipListNode {
	sln := &SkipListNode{}
	sln.key = key
	sln.value = value
	sln.up = nil
	sln.down = nil
	sln.left = nil
	sln.right = nil

	return sln
}

func (sln *SkipListNode) SetUp(up *SkipListNode) {
	sln.up = up
}

func (sln *SkipListNode) GetUp() *SkipListNode {
	return sln.up
}

func (sln *SkipListNode) SetDown(down *SkipListNode) {
	sln.down = down
}

func (sln *SkipListNode) GetDown() *SkipListNode {
	return sln.down
}

func (sln *SkipListNode) SetLeft(left *SkipListNode) {
	sln.left = left
}

func (sln *SkipListNode) GetLeft() *SkipListNode {
	return sln.left
}

func (sln *SkipListNode) SetRight(right *SkipListNode) {
	sln.right = right
}

func (sln *SkipListNode) GetRight() *SkipListNode {
	return sln.right
}

func (sln *SkipListNode) SetKey(key string) {
	sln.key = key
}

func (sln *SkipListNode) GetKey() string {
	return sln.key
}

func (sln *SkipListNode) SetValue(value string) {
	sln.value = value
}

func (sln *SkipListNode) GetValue() string {
	return sln.value
}
