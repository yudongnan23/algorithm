package _0211120

import (
	"github.com/yudongnan23/algorithm"
	"testing"
)

type maxSubArrayParam struct {
	name string
	arg  []int
	want interface{}
	rlt  interface{}
}

func Test_maxSubArray(t *testing.T) {
	params := []maxSubArrayParam{
		{
			name: "1",
			arg:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
	}
	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			algorithm.New(maxSubArray(p.arg), p.want, t).Valid()
		})
	}

}
