package main

import "testing"

type lengthOfLongestSubstringParam struct {
	name string
	s    string
	want int
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	var params = []lengthOfLongestSubstringParam{
		{
			name: "1",
			s:    "sssd",
			want: 2,
		},
		{
			name: "2",
			s: "acacacacacacacacac",
			want: 2,
		},
		{
			name: "3",
			s: "dsssss",
			want: 2,
		},
		{
			name: "4",
			s: "abcdefghijk",
			want: 11,
		},
		{
			name: "5",
			s: "qabcdefghijkq",
			want: 12,
		},
		{
			name: "6",
			s: "dvdf",
			want: 3,
		},
		{
			name: "7",
			s: "tmmzuxt",
			want: 5,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := lengthOfLongestSubstring(p.s)
			if rlt != p.want {
				t.Errorf("rlt: %d not want: %d", rlt, p.want)
			}
		})
	}
}
