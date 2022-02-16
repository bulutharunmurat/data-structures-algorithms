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

func (binarySearchTree *BinarySearchTree) search(data int) *TreeNode {

	p := binarySearchTree.root

	for p != nil {
		if p.data == data {
			return p
		} else if p.data < data {
			p = p.right
		} else if p.data > data {
			p = p.left
		}
	}

	return nil
}

func (binarySearchTree *BinarySearchTree) searchParent(data int) *TreeNode {

	parent := binarySearchTree.root
	p := binarySearchTree.root

	if binarySearchTree.root.data == data {
		return nil
	}

	for p != nil {
		if p.data == data {
			return parent
		} else if p.data < data {
			parent = p
			p = p.right
		} else if p.data > data {
			parent = p
			p = p.left
		}
	}
	return nil
}

func (binarySearchTree *BinarySearchTree) printInOrder() {
	fmt.Print("IN-ORDER TRAVERSE: ")
	printInOrderRecursive(binarySearchTree, binarySearchTree.root)
	fmt.Println()
}

func printInOrderRecursive(binarySearchTree *BinarySearchTree, treeNode *TreeNode) {
	if treeNode != nil {
		printInOrderRecursive(binarySearchTree, treeNode.left)
		fmt.Print(treeNode.data, " ")
		printInOrderRecursive(binarySearchTree, treeNode.right)
	}
}

func (binarySearchTree *BinarySearchTree) printPreOrder() {
	fmt.Print("PRE-ORDER TRAVERSE: ")
	printPreOrderRecursive(binarySearchTree, binarySearchTree.root)
	fmt.Println()
}

func printPreOrderRecursive(binarySearchTree *BinarySearchTree, treeNode *TreeNode) {
	if treeNode != nil {
		fmt.Print(treeNode.data, " ")
		printPreOrderRecursive(binarySearchTree, treeNode.left)
		printPreOrderRecursive(binarySearchTree, treeNode.right)
	}
}

func (binarySearchTree *BinarySearchTree) printPostOrder() {
	fmt.Print("POST-ORDER TRAVERSE: ")
	printPostOrderRecursive(binarySearchTree, binarySearchTree.root)
	fmt.Println()
}

func printPostOrderRecursive(binarySearchTree *BinarySearchTree, treeNode *TreeNode) {
	if treeNode != nil {
		printPostOrderRecursive(binarySearchTree, treeNode.left)
		printPostOrderRecursive(binarySearchTree, treeNode.right)
		fmt.Print(treeNode.data, " ")
	}
}
