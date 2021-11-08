package _0211106

import (
	"fmt"
	"testing"
)

type queueParam struct {
	name   string
	args   []string
	values []int
	want   []int
}

func TestCQueue_Queue(t *testing.T) {
	var parmas = []queueParam{
		{
			name: "1",
			args: []string{"deleteHead","appendTail","deleteHead","appendTail","appendTail","deleteHead","deleteHead","deleteHead","appendTail","deleteHead","appendTail","appendTail","appendTail","appendTail","appendTail","appendTail","deleteHead","deleteHead","deleteHead","deleteHead"},
			values: []int{0,12,0,10,9,0,0,0,20,0,1,8,20,1,11,2,0,0,0,0},
			want: []int{-1,0,12,0,0,10,9,-1,0,20,0,0,0,0,0,0,1,8,20,1},
		},
	}

	q := Constructor()

	for _, p := range parmas {
		t.Run(p.name, func(t *testing.T) {
			for i := range p.args {
				fmt.Println("=========")
				fmt.Println(p.args[i])

				switch p.args[i] {
				case "deleteHead":
					rlt := q.DeleteHead()
					if rlt != p.want[i] {
						t.Errorf("arg: %d  rlt: %d not want: %d", i, rlt, p.want[i])
					}
					fmt.Println(rlt)
				case "appendTail":
					fmt.Println(p.values[i])
					q.AppendTail(p.values[i])
				}
			}
		})
	}
}
