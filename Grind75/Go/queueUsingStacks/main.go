package main

type MyQueue struct {
	nodes []int
}

func Constructor() MyQueue {
	return MyQueue{
		[]int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.nodes = append(this.nodes, x)
}

func (this *MyQueue) Pop() int {
	if len(this.nodes) == 0 {
		return 0
	}
	val := this.nodes[0]
	if len(this.nodes) > 1 {
		this.nodes = this.nodes[1:]
	} else {
		this.nodes = []int{}
	}
	return val
}

func (this *MyQueue) Peek() int {
	return this.nodes[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.nodes) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
