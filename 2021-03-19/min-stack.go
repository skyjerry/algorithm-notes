// Leetcode 155. 最小栈

type MinStack struct {
    stack []int
    minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)

    if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack) - 1] {
        this.minStack = append(this.minStack, val)
    }
}


func (this *MinStack) Pop()  {
    if len(this.stack) == 0 {
        return
    }

    val := this.stack[len(this.stack) - 1]
    this.stack = this.stack[:len(this.stack) - 1]
    if val == this.minStack[len(this.minStack) - 1] {
        this.minStack = this.minStack[:len(this.minStack) - 1]
    }
}


func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        return -1
    }

    return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
    if len(this.stack) == 0 {
        return -1
    }
    return this.minStack[len(this.minStack) - 1]
}