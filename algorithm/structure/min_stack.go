package structure

// 最小栈, 使用辅助栈
type MinStack struct {
	dataStack []int
	minStack  []int //存储最小值,  栈顶总是最小值
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *Minstack) Push(x int) {
	if len(this.stack) == 0 {
		this.dataStack = append(this.dataStack, x)
		this.minStack = append(this.minStack, x)
		return
	}
	// 存储的时候同时计算最小值
	// 分别将x和min存储到两个栈，stack和minStack
	min := this.minStack[len(this.minStack)-1]
	if x < min {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, min)
	}
	this.dataStack = append(this.dataStack, x)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.dataStack[len(this.dataStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
