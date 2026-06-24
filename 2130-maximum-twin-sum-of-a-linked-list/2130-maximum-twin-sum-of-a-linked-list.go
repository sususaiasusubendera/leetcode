/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	arr := []int{}
	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}

	maxSum := 0
	for i := 0; i < len(arr)/2; i++ {
		maxSum = max(maxSum, arr[i]+arr[len(arr)-1-i])
	}

    return maxSum
}

// array, linked list
// time: O(n)
// space: O(n)