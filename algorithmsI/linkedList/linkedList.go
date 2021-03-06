package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	Len  int
}

func (L *List) Insert(key int) {
	ln := ListNode{}
	ln.Val = key

	if L.Len == 0 {
		L.Head = &ln
		L.Len++
		return
	}

	p := L.Head

	for i := 0; i < L.Len; i++ {
		if p.Next == nil {
			p.Next = &ln
			L.Len++
			return
		}
		p = p.Next
	}
}

func (L *List) Init(keys []int) *ListNode {
	L.Len = 0
	for _, v := range keys {
		L.Insert(v)
	}
	return L.Head
}

func (L *List) ToSlice() []int {
	s := []int{}
	n := L.Head
	for n != nil {
		s = append(s, n.Val)
		n = n.Next
	}
	return s
}

//Given the head of a linked list, remove the nth node from the end of the list and return its head.
// Time complexity : O(L).
// The algorithm makes one traversal of the list of LLL nodes. Therefore time complexity is O(L)O(L)O(L).
// Space complexity : O(1).
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	var d, f, s *ListNode

	d = &ListNode{}
	d.Next = head
	f = d
	s = d

	for i := 0; i <= n; i++ {
		f = f.Next
	}

	for f != nil {
		f = f.Next
		s = s.Next
	}

	s.Next = s.Next.Next

	return d.Next
}

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
// Time Complexity: O(N), where NNN is the number of nodes in the given list.
// Space Complexity: O(1), the space used by slow and fast.
func MiddleNode(head *ListNode) *ListNode {

	f, s := head, head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	return s
}

// Given the head of a singly linked list, reverse the list, and return the reversed list.
// Time complexity : O(n). Assume that nnn is the list's length, the time complexity is O(n).
// Space complexity : O(1).
func ReverseList(head *ListNode) *ListNode {
	var p, c *ListNode

	if head == nil {
		return head
	}

	p, c = nil, head

	for c != nil {
		t := c.Next
		c.Next = p
		p = c
		c = t
	}
	return p
}

// Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val >= list2.Val {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = MergeTwoLists(list2, list1.Next)
		return list1
	}
}
