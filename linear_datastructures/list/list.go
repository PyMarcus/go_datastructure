package main

import (
	"fmt"
)

// Lists are a collection of elements and unlikely array
// these can expand dinamically
type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

// addToHead adds an integer property into starts of the linkedlist
func (linkedList *LinkedList) addToHead(property int) {
	node := Node{
		property: property,
	}
	linkedList.headNode = &node
}

// addToTail adds an integer property into ends of the linkedlist
func (linkedList *LinkedList) addToTail(property int) {
	node := &Node{
		property: property,
	}

	var n *Node
	var last *Node

	for n = linkedList.headNode; n != nil; n = n.nextNode {
		if n.nextNode == nil {
			last = n
		}
	}

	if last != nil {
		last.nextNode = node
	}
}

// print display results in stdout
func (linkedlist LinkedList) print() {
	for l := linkedlist.headNode; l != nil; l = l.nextNode {
		fmt.Printf("%d ", l.property)
	}
}

// addNodeWithValue
func (linkedlist *LinkedList) addNodeWithValue(node *Node) {

	for l := linkedlist.headNode; l != nil; l = l.nextNode {
		if l.nextNode == nil {
			l.nextNode = node
			break
		}
	}

}

func main() {
	n := Node{property: 88}
	l := LinkedList{}
	l.addToHead(2)
	//l.addToHead(3)
	//l.addToHead(4)
	l.addToTail(5)
	l.addToTail(0)
	l.addNodeWithValue(&n)
	l.print()
}
