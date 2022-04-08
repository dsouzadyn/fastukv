package datastructures

import "errors"

type Node struct {
	key   uint64
	value uint64
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func newNode(k uint64, v uint64) *Node {
	return &Node{key: k, value: v}
}

type LinkedListBehavior interface {
	Add(k uint64, v uint64)
	Search(k uint64) (uint64, error)
}

func (l *LinkedList) Add(k uint64, v uint64) {
	var n *Node = newNode(k, v)
	if l.head == nil {
		l.head = n
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}

}

func (l *LinkedList) Delete(k uint64) {
	if l.head != nil {
		curr := l.head
		prev := curr
		if curr.key == k {
			l.head = l.head.next
		} else {
			for curr != nil {
				if curr.key == k {
					prev.next = curr.next
					break
				}
				prev = curr
				curr = curr.next

			}
		}

	}
}

func (l *LinkedList) Search(k uint64) (uint64, error) {
	if l.head == nil {
		return uint64(0), errors.New("empty LinkedList")
	}
	curr := l.head
	for curr != nil {
		if curr.key == k {
			return curr.value, nil
		}
		curr = curr.next
	}
	return uint64(0), errors.New("no such key found")
}
