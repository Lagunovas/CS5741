package binaryTree

import (
	"bytes"
	"strconv"
)

type BinaryTree struct {
	count int
	root  *BinaryTreeNode
}

type BinaryTreeNode struct {
	order int
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (binaryTree *BinaryTree) Clear() {
	binaryTree.root = NewBinaryTreeNode()
}

func (binaryTreeNode *BinaryTreeNode) IsRoot() bool {
	return binaryTreeNode.order == 1
}

func (binaryTreeNode *BinaryTreeNode) IsLeaf() bool {
	return binaryTreeNode.left == binaryTreeNode.right && binaryTreeNode.right == nil
}

func (binaryTreeNode *BinaryTreeNode) Value() int {
	return binaryTreeNode.value
}

func NewBinaryTreeNode() *BinaryTreeNode {
	return &BinaryTreeNode{0, 0, nil, nil}
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{0, nil}
}

func (binaryTree *BinaryTree) Empty() bool {
	return binaryTree.root == nil
}

func (binaryTree *BinaryTree) Size() int {
	return binaryTree.count
}

func (binaryTree *BinaryTree) Push(value int) bool {
	return binaryTree.PushOrder(value, binaryTree.count+1)
}

func (binaryTree *BinaryTree) PushOrder(value, order int) bool {
	var binaryTreeNode *BinaryTreeNode = binaryTree.root

	if binaryTree.Empty() {
		binaryTree.root = NewBinaryTreeNode()
		binaryTreeNode = binaryTree.root
	} else {
		for {
			if value < binaryTreeNode.value {
				if binaryTreeNode.left == nil {
					binaryTreeNode.left = NewBinaryTreeNode()
					binaryTreeNode = binaryTreeNode.left
					break
				} else {
					binaryTreeNode = binaryTreeNode.left
				}
			} else if value > binaryTreeNode.value {
				if binaryTreeNode.right == nil {
					binaryTreeNode.right = NewBinaryTreeNode()
					binaryTreeNode = binaryTreeNode.right
					break
				} else {
					binaryTreeNode = binaryTreeNode.right
				}
			} else {
				return false
			}
		}
	}

	binaryTreeNode.value = value

	binaryTree.count++

	binaryTreeNode.order = order
	return true
}

// In Order Traversal
func (binaryTree *BinaryTree) Remove(order int) (bool, int) {
	if !binaryTree.Empty() {
		var stack []*BinaryTreeNode
		currentBinaryTreeNode := binaryTree.root
		var stackSize int = len(stack)

		var currentParentBinaryTreeNode *BinaryTreeNode

		for {
			if currentBinaryTreeNode == nil {
				if stackSize > 0 {
					var poppedElement = stack[stackSize-1]
					stack = stack[:stackSize-1]
					currentBinaryTreeNode = poppedElement.right
					currentParentBinaryTreeNode = poppedElement
				} else {
					break
				}
			} else {
				if currentBinaryTreeNode.order == order {
					removedValue := currentBinaryTreeNode.value
					if currentBinaryTreeNode.IsLeaf() {
						if currentBinaryTreeNode.IsRoot() {
							binaryTree.root = nil
						} else {
							if currentParentBinaryTreeNode.left == currentBinaryTreeNode {
								currentParentBinaryTreeNode.left = nil
							} else {
								currentParentBinaryTreeNode.right = nil
							}
						}

						binaryTree.count--
						return true, removedValue
					} else if currentBinaryTreeNode.left != nil && currentBinaryTreeNode.right != nil {
						minimum := currentBinaryTreeNode.right

						for {
							if minimum.left != nil {
								currentParentBinaryTreeNode = minimum
								minimum = minimum.left
							} else {
								break
							}
						}

						currentBinaryTreeNode.value = minimum.value
						currentBinaryTreeNode.order = minimum.order
						currentParentBinaryTreeNode.left = nil

						binaryTree.count--
						return true, removedValue
					} else {
						toMove := currentBinaryTreeNode.left

						if toMove == nil {
							toMove = currentBinaryTreeNode.right
						}

						if currentParentBinaryTreeNode == nil {
							binaryTree.root = toMove
						} else {
							if currentParentBinaryTreeNode.left == currentBinaryTreeNode {
								currentParentBinaryTreeNode.left = toMove
							} else if currentParentBinaryTreeNode.right == currentBinaryTreeNode {
								currentParentBinaryTreeNode.right = toMove
							}
						}

						binaryTree.count--
						return true, removedValue
					}
				}

				stack = append(stack, currentBinaryTreeNode)
				currentParentBinaryTreeNode = currentBinaryTreeNode
				currentBinaryTreeNode = currentBinaryTreeNode.left
			}

			stackSize = len(stack)
		}
	}
	return false, 0
}

// Level Order Traversal (BFS)
func (binaryTree *BinaryTree) ToString() string {
	var buffer bytes.Buffer

	if !binaryTree.Empty() {
		var queue []*BinaryTreeNode = []*BinaryTreeNode{binaryTree.root}

		for len(queue) > 0 {
			var node *BinaryTreeNode = queue[0]
			buffer.WriteString(strconv.Itoa(node.value) + " ")

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}

			queue = queue[1:]
		}
	}

	return buffer.String()
}

func (binaryTree *BinaryTree) Get(order int) (bool, int) {
	if !binaryTree.Empty() {
		var queue []*BinaryTreeNode = []*BinaryTreeNode{binaryTree.root}

		for len(queue) > 0 {
			var node *BinaryTreeNode = queue[0]

			if node.order == order {
				return true, node.value
			} else {
				if node.left != nil {
					queue = append(queue, node.left)
				}

				if node.right != nil {
					queue = append(queue, node.right)
				}
			}

			queue = queue[1:]
		}
	}

	return false, 0
}

func (binaryTree *BinaryTree) Tail() (bool, int) {
	if !binaryTree.Empty() {
		return binaryTree.Get(binaryTree.count)
	}

	return false, 0
}
