package algorithms

import (
	"fmt"
	. "grokkingingo/datastructures"
)

func RunRemoveNthLastNode() {
	inputs := [][]int{
		{23, 89, 10, 5, 67, 39, 70, 28},
		{34, 53, 6, 95, 38, 28, 17, 63, 16, 76},
		{288, 224, 275, 390, 4, 383, 330, 60, 193},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{69, 8, 49, 106, 116, 112, 104, 129, 39, 14, 27, 12},
	}

	n := []int{4, 1, 6, 9, 11}

	var newlinklist EduLinkedList
	for i := 0; i < len(inputs); i++ {
		newlinklist.CreateLinkedList(inputs[i])
		fmt.Print("Created Linked List : ")
		newlinklist.PrintLinkedList()
		newlinklist.RemoveNthLastNode(n[i])
		fmt.Printf("removed %d-th node\n", n[i])
		fmt.Print("Modified Linked List : ")
		newlinklist.PrintLinkedList()
		fmt.Println()
	}

}
