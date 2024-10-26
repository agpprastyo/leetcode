/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{} // Create a dummy head for the result linked list.
	current := dummyHead      // Pointer to build the new linked list.
	carry := 0                // Variable to hold the carry over value.

	// Loop through both linked lists until both are exhausted.
	for l1 != nil || l2 != nil || carry != 0 {
		var x, y int // x for l1's value, y for l2's value.

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry // Calculate sum of the two digits and the carry.
		carry = sum / 10     // Update carry for the next iteration.
		current.Next = &ListNode{Val: sum % 10} // Create a new node for the result.
		current = current.Next // Move to the next node in the result list.
	}

	return dummyHead.Next // Return the next of the dummy head to exclude the dummy node.
}