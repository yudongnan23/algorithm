package _0211127

import (
	"reflect"
	"testing"
)

type exchangeParam struct {
	name string
	arg  []int
	want [][]int
}

func Test_exchange(t *testing.T) {
	params := []exchangeParam{
		{
			name: "1",
			arg:  []int{1, 2, 3, 4},
			want: [][]int{{1, 3, 2, 4}, {3, 1, 2, 4}},
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := exchange(p.arg)
			for _, w := range p.want {
				if reflect.DeepEqual(rlt, w){
					return
				}
			}
			t.Errorf("rlt: %v not one of want: %v", rlt, p.want)
		})
	}
}
