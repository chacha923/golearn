package structure

type MyCircularDeque struct {
	cache    []int
	capacity int
	length   int
	front    int
	rear     int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cache:    make([]int, k),
		capacity: k,
		front:    1,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length == this.capacity {
		return false
	}
	this.length++
	this.front--
	if this.front == -1 {
		this.front = this.capacity - 1
	}
	this.cache[this.front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length == this.capacity {
		return false
	}
	this.length++
	this.rear++
	if this.rear == this.capacity {
		this.rear = 0
	}
	this.cache[this.rear] = value
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.length == 0 {
		return false
	}
	this.length--
	this.front++
	if this.front == this.capacity {
		this.front = 0
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.length == 0 {
		return false
	}
	this.length--
	this.rear--
	if this.rear == -1 {
		this.rear = this.capacity - 1
	}
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.length == 0 {
		return -1
	}
	return this.cache[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.length == 0 {
		return -1
	}
	return this.cache[this.rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.capacity
}
