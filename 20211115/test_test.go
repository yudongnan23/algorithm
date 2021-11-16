package _0211115

import "testing"

type isSubStructureParam struct {
	name string
	arg struct{
		a *TreeNode
		b *TreeNode
	}
	want bool
}

func Test_isSubStructure(t *testing.T) {
	params := []isSubStructureParam{
		{
			name: "1",
			arg: struct {
				a *TreeNode
				b *TreeNode
			}{
				a: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
				b: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			},
			want: true,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := isSub(p.arg.a, p.arg.b)
			if rlt != p.want {
				t.Errorf("rlt: %v not want: %v", rlt, p.want)
			}
		})
	}
}
