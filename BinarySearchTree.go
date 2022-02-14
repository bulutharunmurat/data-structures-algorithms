package main

type BinarySearchTree struct {
	root *TreeNode
}

func (binarySearchTree *BinarySearchTree) add(data int) {

	node := TreeNode{data: data}

	if binarySearchTree.root == nil {
		binarySearchTree.root = &node
	}

}
