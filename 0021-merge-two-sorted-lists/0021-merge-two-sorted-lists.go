
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to simplify the merge process
	dummy := &ListNode{}
	current := dummy

	// Merge the two lists
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// If there are remaining nodes in either list, append them
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Return the merged list, starting from the next of the dummy node
	return dummy.Next
}

// Helper function to print the linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// Helper function to create a linked list from a slice of integers
func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}