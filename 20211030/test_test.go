package main

import "testing"

type lengthOfLongestSubstringParam struct {
	name string
	s string
	want int
}

func Test_lengthOfLongestSubstring(t *testing.T){
	var params = []lengthOfLongestSubstringParam{
		{
			name: "1",
			s: "sssd",
			want: 2,
		},
	}

	for _, p := range params{
		t.Run(p.name, func(t *testing.T) {
			rlt := lengthOfLongestSubstring(p.s)
			if rlt != p.want {
				t.Errorf("rlt: %d not want: %d", rlt, p.want)
			}
		})
	}
}
