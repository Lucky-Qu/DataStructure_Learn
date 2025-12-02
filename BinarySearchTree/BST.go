package BinarySearchTree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BinarySearchTree[K constraints.Ordered, V any] struct {
	rootNode *TreeNode[K, V]
}
type TreeNode[K constraints.Ordered, V any] struct {
	key       K
	value     V
	leftNode  *TreeNode[K, V]
	rightNode *TreeNode[K, V]
}

func NewBinarySearchTree[K constraints.Ordered, V any]() *BinarySearchTree[K, V] {
	return &BinarySearchTree[K, V]{}
}

func (tree *BinarySearchTree[K, V]) Insert(key K, value V) {
	// 创建节点
	newNode := &TreeNode[K, V]{key: key, value: value}
	// 查看树是否为空
	if tree.rootNode == nil {
		tree.rootNode = newNode
		return
	}
	// 寻找插入点
	for currNode := tree.rootNode; currNode != nil; {
		// 重复Key检查
		if currNode.key == key {
			fmt.Println("禁止使用相同的Key")
			return
		}
		// 插入检查
		if newNode.key > currNode.key {
			// 检查右节点是否为空
			if currNode.rightNode == nil {
				currNode.rightNode = newNode
				return
			}
			// 进入右节点
			currNode = currNode.rightNode
			continue
		} else {
			// 检查左节点是否为空
			if currNode.leftNode == nil {
				currNode.leftNode = newNode
				return
			}
			currNode = currNode.leftNode
			continue
		}
	}
}
