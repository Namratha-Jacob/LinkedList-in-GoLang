//implementing stack using linked list

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LList struct {
	head   *Node
	length int
}

func (l *LList) pushStack(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LList) displayStack() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *LList) popStack() {
	if l.length == 0 {
		return
	} else {
		l.head = l.head.next
		l.length--
		return
	}
}

func main() {
	mylist := LList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	node5 := &Node{data: 50}

	mylist.pushStack(node1)
	mylist.pushStack(node2)
	mylist.pushStack(node3)
	mylist.pushStack(node4)
	mylist.pushStack(node5)

	mylist.displayStack()

	mylist.popStack()

	mylist.displayStack()

}
