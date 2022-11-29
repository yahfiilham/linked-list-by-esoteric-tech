package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepand(v int) {
	newNode := node{data: v}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}

func (l *linkedList) printList() {
	if l.head == nil {
		return
	}

	currentNode := l.head
	i := 1
	for currentNode != nil {
		fmt.Printf("list %d is %d\n", i, currentNode.data)
		currentNode = currentNode.next
		i++
	}
}

func (l *linkedList) deleteWithValue(v int) {
	if l.length == 0 {
		return
	}

	if l.head.data == v {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != v {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	var myList linkedList
	myList.prepand(10)
	myList.prepand(20)
	myList.prepand(30)
	myList.prepand(40)
	myList.deleteWithValue(30)
	myList.printList()
}
