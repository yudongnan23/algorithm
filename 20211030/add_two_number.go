package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * http link: https://leetcode-cn.com/problems/add-two-numbers/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) getValSlice() []int {
	v := make([]int, 0)
	for l != nil {
		v = append(v, l.Val)
		l = l.Next
	}
	return v
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return makeLinkList(getNumber(l1) + getNumber(l2))
}

func makeLinkList(number int) *ListNode {
	fmt.Println(number)

	l := &ListNode{
		Val: number % 10,
	}
	head := l

	length := len(fmt.Sprintf("%d", number))
	for i := 1; i < length; i++ {
		l.Next = &ListNode{
			Val: (number / int(math.Pow(10, float64(i)))) % 10,
		}
		l = l.Next
	}

	return head
}

func getNumber(l *ListNode) int {
	var numberStr string
	for l != nil {
		numberStr = fmt.Sprintf("%d", l.Val) + numberStr
		l = l.Next
	}
	if len(numberStr) == 0 {
		return 0
	}
	number, _ := strconv.Atoi(numberStr)
	fmt.Println(number)
	return number
}

func makeLinkListBySlice(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	l := &ListNode{Val: s[0]}
	head := l
	for i := 1; i < len(s); i++ {
		l.Next = &ListNode{Val: s[i]}
		l = l.Next
	}
	return head
}

func addTwoNumbersMethod2(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head        *ListNode
		point       *ListNode
		isCarryOver bool
	)

	for i := 0; l1 != nil || l2 != nil; i++ {
		var v1, v2, count int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		count, isCarryOver = getCount(v1+v2, isCarryOver)
		node := ListNode{Val: count}
		if i == 0 {
			head = &node
			point = head
		} else {
			point.Next = &node
			point = point.Next
		}
	}

	if isCarryOver {
		point.Next = &ListNode{Val: 1}
	}

	return head
}

func getCount(count int, isCarryOver bool) (int, bool) {
	if isCarryOver {
		count = count + 1
	}
	rlt := count
	continueCarryOver := false
	if count >= 10 {
		rlt = count % 10
		continueCarryOver = true
	}
	return rlt, continueCarryOver
}

func main() {
	var params = []struct {
		name  string
		l1    *ListNode
		l2    *ListNode
		wants []int
	}{
		{
			name:  "1",
			l1:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			l2:    &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}},
			wants: []int{5, 7, 9},
		},
		{
			name:  "2",
			l1:    &ListNode{Val: 2, Next: &ListNode{Val: 9, Next: &ListNode{Val: 8}}},
			l2:    &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 7}}},
			wants: []int{3, 0, 6, 1},
		},
		{
			name:  "3",
			l1:    makeLinkListBySlice([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			l2:    makeLinkListBySlice([]int{5, 6, 4}),
			wants: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			name:  "4",
			l1:    &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:    &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			wants: []int{7, 0, 8},
		},
	}

	success := true
	for _, p := range params {
		l := addTwoNumbersMethod2(p.l1, p.l2)
		if !reflect.DeepEqual(l.getValSlice(), p.wants) {
			success = false
			fmt.Printf("Fail result: %+v not wants: %+v/%s\n", l.getValSlice(), p.wants, p.name)
			continue
		}
		fmt.Printf("success/%s\n", p.name)
	}
	if success {
		fmt.Println("Success!")
	} else {
		fmt.Println("Fail!")
	}
}
