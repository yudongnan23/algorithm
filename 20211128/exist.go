package _0211128

import (
	"fmt"
	"github.com/aws/smithy-go"
	"github.com/go-playground/locales/currency"
	"strings"
)

/**
 * 矩阵中的路径：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
 */

type coordinate struct {
	x int
	y int
}

func exist(b [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	if len(b) == 0 || len(b[0]) == 0 {
		return false
	}

	q := make([]coordinate, 0)
	q[0] = coordinate{0, 0}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if b[cur.x][cur.y] == word[0] && search(b, word, cur) {
			return true
		}

		if cur.x < len(b) {
			q = append(q, coordinate{x: cur.x + 1, y: cur.y})
		}

		if cur.y < len(b[0]) {
			q = append(q, coordinate{x: cur.x, y: cur.y + 1})
		}
	}

	return false
}

func search(b [][]byte, word string, c coordinate) bool {
	q := []coordinate{c}
	dump := make(map[string]int, 0)
	dump[makeKey(c.x, c.y)] = 0

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if b[cur.x][cur.y] != word[dump[makeKey(cur.x, cur.y)]]{
			continue
		}

		if cur.x > 0 {
			if !dump[makeKey(c.x -1, c.y)] {
				q = append()
			}
		}
	}
}

func makeKey(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}
