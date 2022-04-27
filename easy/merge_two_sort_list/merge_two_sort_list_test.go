package merge_two_sort_list

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	n1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	n2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	want := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}

	got := mergeTwoLists(&n1, &n2)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got!= want: got=%v, want=%v", got, want)
	}
}
