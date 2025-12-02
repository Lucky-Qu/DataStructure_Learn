package main

import "DataStructure_Learn/OrderedNodeList"

func main() {
	// 获取一个新有序链表
	NewOrderedNodeList := OrderedNodeList.NewNodeList[int, string]()
	NewOrderedNodeList.AddNode(2, "asd")
}
