package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: A Tree
 * @return: Level order a list of lists of integer
 */
type localQueue struct {
	lTreeNode *TreeNode
	next      *localQueue
	length    int
}

func (q *localQueue) enQueue(node *TreeNode) {
	temp := &localQueue{}
	temp.lTreeNode = new(TreeNode)
	temp.lTreeNode = node
	temp.next = nil
	for q.next != nil {
		q = q.next
	}
	q.next = new(localQueue)
	q.next = temp
}

func deQueue(q *localQueue) (*TreeNode, *localQueue) {
	if q == nil {
		return nil, nil
	}
	temp := q.next
	return q.lTreeNode, temp
}

func qLength(q *localQueue) int {
	temp := q
	if q == nil {
		return 0
	}
	count := 0
	for temp != nil {
		temp = temp.next
		count += 1
	}
	return count
}

func levelOrder(root *TreeNode) [][]int {
	retVal := make([][]int, 0)
	// write your code here
	if root == nil {
		return retVal
	}
	flagTree := &TreeNode{}
	flagTree.Val = -1
	flagTree.Left = nil
	flagTree.Right = nil

	localVal := make([]int, 0)
	q := new(localQueue)
	q.lTreeNode = new(TreeNode)
	q.lTreeNode = root
	q.next = nil
	q.enQueue(flagTree)
	var child = &TreeNode{}
	for qLength(q) > 0 {
		child, q = deQueue(q)
		if nil != child {
			if child.Val == -1 {
				retVal = append(retVal, localVal)
				if qLength(q) == 0 {
					return retVal
				}
				q.enQueue(flagTree)
				localVal = make([]int, 0)
				continue
			}
			localVal = append(localVal, child.Val)
			if child.Left != nil {
				q.enQueue(child.Left)
			}
			if child.Right != nil {
				q.enQueue(child.Right)
			}
		}

	}
	return retVal
}
