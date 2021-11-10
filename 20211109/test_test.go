package _0211109

import (
	"testing"
)

type searchParam struct {
	name string
	arg  struct {
		n []int
		t int
	}
	want int
}

func Test_search(t *testing.T) {
	var params = []searchParam{
		{
			name: "1",
			arg: struct {
				n []int
				t int
			}{
				n: []int{2, 2},
				t: 3,
			},
			want: 0,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := search(p.arg.n, p.arg.t)
			if rlt != p.want {
				t.Errorf("rlt: %d not want: %d", rlt, p.want)
			}
		})
	}
}

type sumParam struct {
	name string
	arg int
	want int
}

func Test_sum(t *testing.T){
	params := []sumParam{{
		name: "1",
		arg: 10,
		want: 55,
	}}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := sum(p.arg)
			if rlt != p.want{
				t.Errorf("rlt: %d not want: %d", rlt, p.want)
			}
		})
	}
}
