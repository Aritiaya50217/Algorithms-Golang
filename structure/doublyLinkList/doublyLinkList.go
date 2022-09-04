package main

import "fmt"

type node struct {
	data string
	prev *node
	next *node
}

type doublyLinkList struct {
	len  int
	tail *node
	head *node
}

func initDoublyList() *doublyLinkList {
	return &doublyLinkList{}
}

func (d *doublyLinkList) AddFrontNodeDLL(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}

func (d *doublyLinkList) AddEndNodeDLL(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
	return
}

func (d *doublyLinkList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is Empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("value = %v , prev = %v , next= %v\n", temp.data, temp.prev, temp.next)
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *doublyLinkList) TraverseReverse() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v , prev = %v , next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.prev
	}
	fmt.Println()
	return nil
}

func (d *doublyLinkList) Size() int {
	return d.len
}

func main() {
	doublyList := initDoublyList()
	fmt.Printf("Add Front Node: C\n")
	doublyList.AddFrontNodeDLL("C")
	fmt.Printf("Add Front Node: B\n")
	doublyList.AddFrontNodeDLL("B")
	fmt.Printf("Add Front Node: A\n")
	doublyList.AddFrontNodeDLL("A")
	fmt.Printf("Add End Node: D\n")
	doublyList.AddEndNodeDLL("D")
	fmt.Printf("Add End Node: E\n")
	doublyList.AddEndNodeDLL("E")

	fmt.Printf("Size of doubly linked ist: %d\n", doublyList.Size())

	err := doublyList.TraverseForward()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = doublyList.TraverseReverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	// [] => add front node : C => [C]
	// add front node : B => [B,C]
	// add front node : A => [A,B,C]
	// add End node : D => [A,B,C,D]
	// add End node : C => [A,B,C,D,E]

	// TraverseForward
	// value = A , prev = [] , next = [B]
	// value = B , prev = [A] , next = [C]
	// value = C , prev = [B] , next = [D]
	// value = D , prev = [C] , next = [E]
	// value = E , prev = [D] , next = []

	// TraverseReverse (ย้อนกลับ)
	// value = E , prev = [D] , next = []
	// value = D , prev = [C] , next = [E]
	// value = C , prev = [B] , next = [D]
	// value = B , prev = [A] , next = [C]
	// value = A , prev = [] , next = [B]

}
