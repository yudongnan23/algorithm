package _0211114

import (
	"reflect"
	"testing"
)

type levelOrderParam struct {
	name string
	arg *TreeNode
	want []int
}

func Test_levelOrder(t *testing.T) {
	params := []levelOrderParam {
		{
			name: "1",
			arg: &TreeNode{v: 1, l: &TreeNode{v: 2}, r: &TreeNode{v: 3}},
			want: []int{1, 2, 3},
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := levelOrder(p.arg)
			if !reflect.DeepEqual(rlt, p.want) {
				t.Errorf("rlt: %v not want: %v", rlt, p.want)
			}
		})
	}
}