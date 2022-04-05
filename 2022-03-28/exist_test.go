package _022_03_27

import (
	"github.com/go-playground/assert"
	"testing"
)

type findParam struct {
	name string
	args struct {
		nums [][]byte
		word []byte
		i    int
		j    int
	}
	want bool
}

func Test_Find(t *testing.T) {
	pp := []findParam{
		{
			name: "1",
			args: struct {
				nums [][]byte
				word []byte
				i    int
				j    int
			}{
				nums: [][]byte{[]byte(`ABCE`), []byte(`SFCS`), []byte(`ADEE`)},
				word: []byte(`ABCCED`),
				i:    0,
				j:    0,
			},
			want: true,
		},
		{
			name: "2",
			args: struct {
				nums [][]byte
				word []byte
				i    int
				j    int
			}{
				nums: [][]byte{[]byte(`ab`)},
				word: []byte(`ba`),
				i:    0,
				j:    1,
			},
			want: true,
		},
	}

	for _, p := range pp {
		t.Run(p.name, func(t *testing.T) {
			assert.Equal(t, newFinder(p.args.nums, p.args.word, p.args.i, p.args.j).find(), p.want)
		})
	}
}
