package structure

import "testing"

func TestStack(t *testing.T) {
	var stack = NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Pop() != 3 {
		t.Error("stack pop error")
	}
	if stack.Pop() != 2 {
		t.Error("stack pop error")
	}
	if stack.Pop() != 1 {
		t.Error("stack pop error")
	}
}

func TestQueue(t *testing.T) {
	var queue = NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if queue.Pop() != 1 {
		t.Error("queue pop error")
	}
	if queue.Pop() != 2 {
		t.Error("queue pop error")
	}
	if queue.Pop() != 3 {
		t.Error("queue pop error")
	}
}
