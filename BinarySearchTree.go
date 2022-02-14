package main

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
