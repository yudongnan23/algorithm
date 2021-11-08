package _0211107

import (
	"reflect"
	"testing"
)

type reverseListParam struct{
	name string
	arg  *ListNode
	want []int
}

func Test_reverseList (t *testing.T) {
	var params = []reverseListParam{
		{
			name: "1",
			arg: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			want: []int{3, 2, 1},
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := reverseList(p.arg)
			if !reflect.DeepEqual(rlt.getValueAsList(), p.want) {
				t.Errorf("rlt: %v not want: %v", rlt, p.want)
			}
		})
	}
}