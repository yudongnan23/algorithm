package _022_02_26

import (
	"github.com/go-playground/assert"
	"testing"
)

type reverseLeftWordsParam struct {
	name string
	args struct {
		s string
		k int
	}
	want string
}

func Test_reverseLeftWords(t *testing.T) {
	pp := []reverseLeftWordsParam{
		{
			name: "1",
			args: struct {
				s string
				k int
			}{
				s: "abcdefg",
				k: 2,
			},
			want: "cdefgab",
		},
		{
			name: "2",
			args: struct {
				s string
				k int
			}{
				s: "lrloseumgh",
				k: 6,
			},
			want: "umghlrlose",
		},
	}

	for _, p := range pp {
		t.Run(p.name, func(t *testing.T) {
			assert.Equal(t, reverseLeftWords(p.args.s, p.args.k), p.want)
		})
	}
}
