package _022_03_27

import (
	"fmt"
)

// 矩阵中的路径
// leetcode链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/

/**
 * 思路：
 */

func exist(board [][]byte, word string) bool {
	var exist bool
	if len(board) == 0 {
		return exist
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] && newFinder(board, []byte(word), i, j).find() {
				exist = true
				break
			}
		}
		if exist {
			break
		}
	}
	return exist
}

func newFinder(nums [][]byte, word []byte, i, j int) *finder {
	return &finder{
		n: nums,
		w: word,
		i: i,
		j: j,
	}
}

type finder struct {
	n [][]byte
	w []byte
	i int
	j int
}

type index struct {
	i       int
	j       int
	wIndex  int
	mapping map[string]bool
}

func makeMappingField(i, j int) string {
	return fmt.Sprintf("%d_%d", i, j)
}

func (f *finder) find() bool {
	queue := []index{{i: f.i, j: f.j, wIndex: 0, mapping: map[string]bool{makeMappingField(f.i, f.j): true}}}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.wIndex == len(f.w)-1 {
			return true
		}

		if cur.i-1 >= 0 && !cur.mapping[makeMappingField(cur.i-1, cur.j)] && f.n[cur.i-1][cur.j] == f.w[cur.wIndex+1] {
			mapping := cur.mapping
			mapping[makeMappingField(cur.i-1, cur.j)] = true
			queue = append(queue, index{i: cur.i - 1, j: cur.j, wIndex: cur.wIndex + 1, mapping: mapping})
		}

		if cur.i+1 < len(f.n) && !cur.mapping[makeMappingField(cur.i+1, cur.j)] && f.n[cur.i+1][cur.j] == f.w[cur.wIndex+1] {
			mapping := cur.mapping
			mapping[makeMappingField(cur.i+1, cur.j)] = true
			queue = append(queue, index{i: cur.i + 1, j: cur.j, wIndex: cur.wIndex + 1, mapping: mapping})

		}

		if cur.j+1 < len(f.n[0]) && !cur.mapping[makeMappingField(cur.i, cur.j+1)] && f.n[cur.i][cur.j+1] == f.w[cur.wIndex+1] {
			mapping := cur.mapping
			mapping[makeMappingField(cur.i, cur.j+1)] = true
			queue = append(queue, index{i: cur.i, j: cur.j + 1, wIndex: cur.wIndex + 1})

		}

		if cur.j-1 >= 0 && !cur.mapping[makeMappingField(cur.i, cur.j-1)] && f.n[cur.i][cur.j-1] == f.w[cur.wIndex+1] {
			mapping := cur.mapping
			mapping[makeMappingField(cur.i, cur.j-1)] = true
			queue = append(queue, index{i: cur.i, j: cur.j - 1, wIndex: cur.wIndex + 1})

		}
	}

	return false
}
