package _0211108

import "testing"

type replaceSpaceParam struct {
	name string
	s    string
	want string
}

func Test_replaceSpace(t *testing.T) {
	var params = []replaceSpaceParam{
		{
			name: "1",
			s:    "ss c",
			want: "ss%20c",
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := replaceSpace(p.s)
			if rlt != p.want {
				t.Errorf("rlt: %s not want: %s", rlt, p.want)
			}
		})
	}
}

type reverseLeftWordsParam struct {
	name string
	arg  struct {
		s string
		k int
	}
	want string
}

func Test_reverseLeftWords(t *testing.T) {
	var params = []reverseLeftWordsParam{
		{
			name: "1",
			arg: struct {
				s string
				k int
			}{
				s: "ccsca",
				k: 1,
			},
			want: "cscac",
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := reverseLeftWords(p.arg.s, p.arg.k)
			if rlt != p.want{
				t.Errorf("rlt: %s not want: %s", rlt, p.want)
			}
		})
	}
}
