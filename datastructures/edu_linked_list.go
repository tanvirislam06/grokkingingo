package datastructures

import "fmt"

type EduLinkedList struct {
	head *EduLinkedListNode
}

func (l *EduLinkedList) InsertNodeAtHead(node *EduLinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.Next = l.head
		l.head = node
	}
}

func (l *EduLinkedList) PrintLinkedList() {
	if l.head == nil {
		fmt.Println("The list is empty")
	} else {
		fmt.Print("[")
		// if head is not nil set temp to head
		temp := l.head
		for temp != nil {
			// print the data
			fmt.Print(temp.Data)
			// point temp to next node and if next node is not empty print comma
			temp = temp.Next
			if temp != nil {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
	}
}

func (l *EduLinkedList) CreateLinkedList(nums []int) {
	for i := 0; i < len(nums); i++ {
		newNode := InitLinkedListNode(nums[i])
		l.InsertNodeAtHead(newNode)
	}
}

func (l *EduLinkedList) RemoveNthLastNode(n int) *EduLinkedListNode {

	// set two pointer left and right at the head of the node
	left := l.head
	right := l.head

	// move right pointer n nodes forward so right and left always have a distance of n
	for i := 0; i < n; i++ {
		right = right.Next
	}

	// If the right pointer reaches NULL while moving it n steps forward, the head node should be removed. We return the head’s Next node.
	if right == nil {
		return l.head.Next
	}

	// move both pointers forward until the right pointer reaches the end of the last node

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next
	return l.head
}

func (l *EduLinkedList) GetNthNodeFromHead(nthPosition int) *EduLinkedListNode {
	//if head is not set or node is just head
	// return nil
	if l.head == nil || l.head.Next == nil {
		return nil
	} else { // else traverse through node.next til the nth node and return node
		nthNode := l.head
		nodeCount := 1
		for nthNode.Next != nil {
			if nodeCount == nthPosition {
				break
			}
			nthNode = nthNode.Next
			nodeCount++
		}
		return nthNode
	}
}
