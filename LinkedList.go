/***** Implementing STACK using Linked list *******/

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

// Pushing nodes to the TOP of stack
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

// Popping nodes from Top of stack
func (l *LList) popStack() {
	if l.length == 0 {
		return
	} else {
		fmt.Println(l.head.data)
		l.head = l.head.next
		l.length--
		return
	}
}

func main() {
	mylist := LList{}

	fmt.Println("Creating 5 nodes")

	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	node5 := &Node{data: 50}

	fmt.Println("Pushing the 5 nodes into stack")

	mylist.pushStack(node1)
	mylist.pushStack(node2)
	mylist.pushStack(node3)
	mylist.pushStack(node4)
	mylist.pushStack(node5)

	fmt.Println("Displaying current stack:")

	mylist.displayStack()

	fmt.Println("Popping a node from stack:")

	mylist.popStack()

	fmt.Println("Displaying current stack: ")

	mylist.displayStack()

}
