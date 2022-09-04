package main

import (
	"container/list"
	"fmt"
)

// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)

type customStack struct {
	stack *list.List
}

func (c *customStack) Push(value string) {
	// เพิ่ม value
	c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
	// ดึงค่า value ออกไปเรื่อย ๆ จนกว่า len = 0
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func (c *customStack) Front() (string, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Peep Error: Stack Datatype is incorrect")
	}
	return "", fmt.Errorf("Peep Error: Stack is empty")
}

func (c *customStack) Size() int {

	return c.stack.Len()
}

func (c *customStack) Empty() bool {
	return c.stack.Len() == 0
}

func main() {
	customStack := &customStack{
		stack: list.New(),
	}
	fmt.Printf("Push: A\n")
	customStack.Push("A")
	fmt.Printf("Push: B\n")
	customStack.Push("B")
	fmt.Printf("Size: %d\n", customStack.Size())
	for customStack.Size() > 0 {
		frontVal, _ := customStack.Front()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Pop: %s\n", frontVal)
		customStack.Pop()
	}
	fmt.Printf("Size: %d\n", customStack.Size())
	// [] => push A => [A] => push B => [B,A]  Size : 2
	// pop B => [A] => pop A => []  Size : 0
}
