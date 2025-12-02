package OrderedNodeList

import (
	"golang.org/x/exp/constraints"
)

// NodeList 有序链表
type NodeList[K constraints.Ordered, V any] struct {
	firstNode *Node[K, V]
}

// Node 单节点，有Key和Value和指向下一个节点的指针
type Node[K constraints.Ordered, V any] struct {
	key   K
	value V
	next  *Node[K, V]
}

// NewNodeList 获取一个空链表
func NewNodeList[K constraints.Ordered, V any]() *NodeList[K, V] {
	return &NodeList[K, V]{
		firstNode: nil,
	}
}

// AddNode 向链表中添加一个节点
func (NodeList *NodeList[K, V]) AddNode(key K, value V) {
	// 新建Node
	newNode := &Node[K, V]{key: key, value: value, next: nil}
	// 检查是否为空链表
	if NodeList.firstNode == nil {
		NodeList.firstNode = newNode
		return
	}
	// 寻找插入位置
	// 记录首节点
	prevNode := NodeList.firstNode
	// 判断是否比首节点小
	if newNode.key < prevNode.key {
		newNode.next = prevNode
		NodeList.firstNode = newNode
		return
	}
	// 遍历NodeList
	for currNode := prevNode.next; currNode != nil; currNode = currNode.next {
		if newNode.key < currNode.key {
			newNode.next = currNode
			prevNode.next = newNode
			return
		}
		prevNode = currNode
	}
	prevNode.next = newNode
	return
}
