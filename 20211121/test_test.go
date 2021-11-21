package _0211121

import (
	"github.com/yudongnan23/algorithm"
	"testing"
)

type translateNumParam struct {
	name string
	arg int
	want int
}

func Test_translateNum(t *testing.T){
	params := []translateNumParam{
		{
			name: "1",
			arg: 12221,
			want: 8                                                                                                                    ,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			algorithm.New(translateNum(p.arg), p.want, t).Valid()
		})
	}
}