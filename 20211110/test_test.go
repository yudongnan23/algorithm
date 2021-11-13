package _0211110

import "testing"

type findNumberIn2DArrayMethod2Param struct {
	name string
	arg  struct {
		m [][]int
		t int
	}
	want bool
}

func Test_findNumberIn2DArrayMethod2(t *testing.T) {
	params := []findNumberIn2DArrayMethod2Param{
		{
			name: "1",
			arg: struct {
				m [][]int
				t int
			}{
				m: [][]int{
					{1}, {3}, {5},
					//{1,   4,  7, 11, 15},
					//{2,   5,  8, 12, 19},
					//{3,   6,  9, 16, 22},
					//{10, 13, 14, 17, 24},
					//{18, 21, 23, 26, 30},
				},
				t: 2,
			},
			want: false,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := findNumberIn2DArrayMethod2(p.arg.m, p.arg.t)

			if rlt != p.want {
				t.Errorf("rlt: %v not want: %v", rlt, p.want)
			}
		})
	}
}
