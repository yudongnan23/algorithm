package _022_03_19

import (
	"github.com/go-playground/assert"
	"testing"
)

type maxProfitParam struct {
	name string
	args struct {
		prices []int
	}
	want int
}

func Test_maxProfit(t *testing.T) {
	pp := []maxProfitParam{
		{
			name: "1",
			args: struct{ prices []int }{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "2",
			args: struct{ prices []int }{prices: []int{7, 6, 5, 4, 3, 2, 1}},
			want: 0,
		},
	}

	for _, p := range pp {
		assert.Equal(t, maxProfit(p.args.prices), p.want)
	}
}
