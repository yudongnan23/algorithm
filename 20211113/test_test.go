package _0211113

import "testing"


type minArrayParam struct {
	name string
	arg []int
	want int
}
func Test_minArray(t *testing.T) {
	params := []minArrayParam{
		{
			name: "1",
			arg: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "2",
			arg: []int{4, 5, 1, 2, 3},
			want: 1,
		},
		{
			name: "3",
			arg: []int{1, 1, 1},
			want: 1,
		},
		{
			name: "4",
			arg: []int{1, 3, 3},
			want: 1,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := minArray(p.arg)
			if rlt != p.want {
				t.Errorf("rlt: %d not want: %d", rlt, p.want)
			}
		})
	}
}

type firstUniqCharParam struct {
	name string
	arg string
	want byte
}

func Test_firstUniqChar(t *testing.T) {
	params := []firstUniqCharParam{
		{
			name: "1",
			arg: "abaccdeff",
			want: "b"[0],
		},
		{
			name: "2",
			arg: "aaccffd",
			want: "d"[0],
		},
		{
			name: "3",
			arg: "defccffea",
			want: "d"[0],
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := firstUniqChar(p.arg)
			if rlt != p.want{
				t.Errorf("rlt: %s not want: %s", string(rlt), string(p.want))
			}
		})
	}
}