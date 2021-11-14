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
			arg: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
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

type levelOrderIIParam struct {
	name string
	arg *TreeNode
	want [][]int
}

func Test_levelOrderII(t *testing.T) {
	params := []levelOrderIIParam{
		{
			name: "1",
			arg: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := levelOrderII(p.arg)
			if !reflect.DeepEqual(rlt, p.want) {
				t.Errorf("rlt: %v not want: %v", rlt, p.want)
			}
		})
	}
}

func Test_levelOrderIII(t *testing.T) {
	params := []levelOrderIIParam{
		{
			name: "1",
			arg: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
			want: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := levelOrderIII(p.arg)
			if !reflect.DeepEqual(rlt, p.want) {
				t.Errorf("rlt: %v not want: %v", rlt, p.want)
			}
		})


	}
}