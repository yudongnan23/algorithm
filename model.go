package algorithm

// 单向链表节点
type LinkedListNode struct {
	V    int
	Next *LinkedListNode
}

// 双向链表节点
type DoubleLinkedListNode struct {
	V    int
	Pre  *DoubleLinkedListNode
	Next *DoubleLinkedListNode
}
