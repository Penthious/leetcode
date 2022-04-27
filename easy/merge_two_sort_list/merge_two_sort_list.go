package merge_two_sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	base := &ListNode{}

	if list1 == nil && list2 == nil {
		return base
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	recurse(base, list1, list2)
	return base
}

func recurse(base *ListNode, l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		base.Next = l2
		return recurse(base, l1, l2.Next)
	}
	if l2 == nil && l1 != nil {
		base.Next = l1
		return recurse(base, l1.Next, l2)
	}
	if base.Val == 0 {
		if l1.Val >= l2.Val {
			base.Val = l2.Val
			return recurse(base, l1, l2.Next)
		}

		base.Val = l1.Val
		return recurse(base, l1.Next, l2)
	}
	if l1.Next == nil && l2.Next == nil {
		if l1.Val > l2.Val {
			base.Next = l1
			base.Next.Next = l2
			return base
		}
		base.Next = l2
		base.Next.Next = l1
		return base
	}

	if l1.Val <= l2.Val {
		if l1.Next != nil {
			base.Next = l1
		}

		return recurse(base.Next, l1.Next, l2)
	}

	if l2.Next != nil {
		base.Next = l2
		return recurse(base.Next, l1, l2.Next)
	}
	return base
}
