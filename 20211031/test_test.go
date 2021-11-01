package main

import (
	"testing"
)

type isPalindromeParam struct {
	name string
	s    string
	want bool
}

func Test_isPalindrome(t *testing.T) {
	params := []isPalindromeParam{
		{
			name: "1",
			s:    "aca",
			want: true,
		},
		{
			name: "2",
			s:    "ab",
			want: false,
		},
		{
			name: "3",
			s:    "accd",
			want: false,
		},
		{
			name: "4",
			s:    "acca",
			want: true,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := isPalindrome(p.s)
			if p.want != rlt {
				t.Errorf("rlt: %v not want: %v", rlt, p.want)
			}
		})
	}
}

type longestPalindromeParam struct {
	name string
	s    string
	want string
}

func Test_longestPalindrome(t *testing.T) {
	params := []longestPalindromeParam{
		{
			name: "1",
			s: "sssssssssssssssssssssssss",
			want: "sssssssssssssssssssssssss",
		},
		{
			name: "2",
			s: "asdsdfdfdfd",
			want: "dfdfdfd",
		},
		{
			name: "3",
			s: "babad",
			want: "bab",
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			rlt := longestPalindromeMethod2(p.s)
			if p.want != rlt {
				t.Errorf("rlt: %s not want: %s", rlt, p.want)
			}
		})
	}
}
