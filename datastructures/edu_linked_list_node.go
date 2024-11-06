package datastructures

type EduLinkedListNode struct {
	Data int
	Next *EduLinkedListNode
}

func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	newNode := new(EduLinkedListNode)
	newNode.Data = data
	newNode.Next = next
	return newNode
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	initNode := new(EduLinkedListNode)
	initNode.Data = data
	initNode.Next = nil
	return initNode
}
