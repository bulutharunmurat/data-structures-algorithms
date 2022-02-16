package main

import "fmt"

type BinarySearchTree struct {
	root *TreeNode
}

func (binarySearchTree *BinarySearchTree) add(data int) {

	newNode := TreeNode{data: data}

	if binarySearchTree.root == nil {
		binarySearchTree.root = &newNode
	} else {
		p := binarySearchTree.root
		for true {
			if p.data > newNode.data {
				if p.left == nil {
					p.left = &newNode
					return
				} else {
					p = p.left
				}
			} else if p.data < newNode.data {
				if p.right == nil {
					p.right = &newNode
					return
				} else {
					p = p.right
				}
			} else {
				return
			}
		}
	}
}

func (binarySearchTree *BinarySearchTree) printBF() {

	var queue []*TreeNode //TODO can be implemented with our own data structure
	queue = append(queue, binarySearchTree.root)
	fmt.Print("BREATH-FIRST TRAVERSE: ")
	for len(queue) != 0 {
		treeNode := queue[0]
		queue = queue[1:]

		fmt.Print(treeNode.data, " ")

		if treeNode.left != nil {
			queue = append(queue, treeNode.left)
		}
		if treeNode.right != nil {
			queue = append(queue, treeNode.right)
		}
	}
	fmt.Println()
}

func (binarySearchTree *BinarySearchTree) printInOrder() {
	printInOrderRecursive(binarySearchTree, binarySearchTree.root)
}

func printInOrderRecursive(binarySearchTree *BinarySearchTree, treeNode *TreeNode) {
	if treeNode != nil {
		printInOrderRecursive(binarySearchTree, treeNode.left)
		fmt.Print(treeNode.data, " ")
		printInOrderRecursive(binarySearchTree, treeNode.right)
	}

}
