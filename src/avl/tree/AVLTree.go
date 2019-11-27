// package avlTree
package main

import (
	avlNode "github.com/CS5741/src/avl/node"
)

var Unlinked int32 = 1
var Growing int32 = 2
var GrowCountIncr int32 = 1 << 3
var GrowCountMask int32 = 0xFF << 3
var Shrinking int32 = 4
var ShrinkCountIncr int32 = 1 << 11
var IgnoreGrow int32 = ^(Growing | GrowCountMask)

// CUSTOM
const RETRY = 0
const SUCCESS = 1
const NIL = -1

// UTILS START

func Compare(v0, v1 int) int {
	switch {
	case v0 == v1:
		return 0
	case v0 < v1:
		return -1
	default:
		return 1
	}
}

func Abs(value int64) int64 {
	if value < 0 {
		return -value
	}

	return value
}

// UTILS END

type AVLTree struct {
	root *avlNode.AVLNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{avlNode.NewAVLNode()}
}

// =====================

func AttemptGet(value int, avlNode *avlNode.AVLNode, direction int, nodeVersion int32) (int, int) { // -1 null, 0 retry, 1 found, value
	for {
		child := avlNode.Child(direction)

		if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 { // XOR
			return RETRY, 0
		}

		if child == nil {
			return NIL, 0
		}

		nextDirection := Compare(child.Value, value)

		if nextDirection == 0 {
			return SUCCESS, value
		}

		childVersion := child.LoadVersion()

		if (childVersion & Shrinking) != 0 {
			WaitUntilNotChanging(child)
		} else if childVersion != Unlinked && child == avlNode.Child(direction) {
			if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
				return RETRY, 0
			} else {
				status, value := AttemptGet(value, child, nextDirection, childVersion)

				if status != RETRY {
					return status, value
				}
			}
		}

	}
}

func (avlTree *AVLTree) Get(value int) int {
	_, value = AttemptGet(value, avlTree.root, -1, 0)
	return value
}

var SpinCount int = 100

func WaitUntilNotChanging(avlNode *avlNode.AVLNode) {
	version := avlNode.LoadVersion()

	if (version & (Growing | Shrinking)) != 0 {
		i := 0

		for avlNode.LoadVersion() == version && i < SpinCount {
			if i == SpinCount {
				// synchronized (n) {}
			}
			i++
		}

	}
}

func AttemptPut(value int, avlNode *avlNode.AVLNode, direction int, nodeVersion int32) (int, int) {
	p := 0 // -1 null, 0 retry, 1 found, value
	returnValue := 0

	for {
		child := avlNode.Child(direction)

		if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
			return RETRY, 0
		}

		if child == nil {
			// p = AttemptInsert(value, avlNode, direction, nodeVersion)
			p = 5 // to avoid error/warning etc
		} else {
			nextDirection := Compare(child.Value, value)

			if nextDirection == 0 {
				// p = AttemptUpdate(child, value)
				p = 5 // to avoid error/warning etc
			} else {
				childVersion := child.LoadVersion()

				if (childVersion & Shrinking) != 0 {
					WaitUntilNotChanging(child)
				} else if childVersion != Unlinked && child == avlNode.Child(direction) {
					if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
						return RETRY, 0
					}
					p, returnValue = AttemptPut(value, child, nextDirection, childVersion)
				}
			}
		}

		if p != RETRY {
			break
		}
	}

	return p, returnValue
}

func (avlTree *AVLTree) Put(value int) (int, int) {
	return AttemptPut(value, avlTree.root, -1, 0)
}

func Balance(n *avlNode.AVLNode) int64 {
	var left int64 = 0
	var right int64 = 0

	if n.Left != nil {
		left = n.Left.LoadHeight()
	}

	if n.Right != nil {
		right = n.Right.LoadHeight()
	}

	return left - right
}

const UNLINK = -1
const ROTATE = -2
const NOCHANGE = -3

func NodeStatus(n *avlNode.AVLNode) int64 {
	if (n.Left == nil || n.Right == nil) && n.Value == -1 {
		return UNLINK
	}

	if n.Right != nil && n.Left != nil {
		if Abs(n.Right.LoadHeight()-n.Left.LoadHeight()) > 1 {
			return ROTATE // rotate
		}

		return NOCHANGE // noChange
	} else if n.Right == nil {
		if n.Left == nil || n.Left.LoadHeight() <= 1 {
			return NOCHANGE
		}

		return ROTATE
	} else {
		if n.Right.LoadHeight() <= 1 {
			return NOCHANGE
		}

		return ROTATE
	}
}

func (avlTree *AVLTree) FixHeightAndRotate(node *avlNode.AVLNode) {
	status := NodeStatus(node)

	switch status {
	case UNLINK:
		avlTree.UnlinkNode(node)
	case ROTATE:
		if node == nil {
			return
		}

		node.Mutex.Lock()
		defer node.Mutex.Unlock()

		balance := Balance(node)

		if balance >= 2 {
			if Balance(node.Left) < 0 {
				RotateLeft(node.Left)
			}
			RotateRight(node)
		} else if balance <= -2 {
			if Balance(node.Right) > 0 {
				RotateRight(node.Right)
			}
			RotateLeft(node)
		}
	}
}

func (avlTree *AVLTree) AttemptInsert(value int, avlNode *avlNode.AVLNode, direction int, nodeVersion int32) int {
	avlNode.Mutex.Lock()

	if ((avlNode.LoadVersion()^nodeVersion)&IgnoreGrow) != 0 || avlNode.Child(direction) != nil {
		avlNode.Mutex.Unlock()
		return RETRY
	} else {
		avlNode.SetChild(direction, nil)
		avlNode.Mutex.Unlock()
	}

	avlTree.FixHeightAndRotate(avlNode)
	return SUCCESS
}

func RotateRight(n *avlNode.AVLNode) {
	nP := n.Parent
	nL := n.Left
	nLR := nL.Right

	nP.Mutex.Lock()
	defer nP.Mutex.Unlock()
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	nL.Mutex.Lock()
	defer nL.Mutex.Unlock()

	n.StoreVersion(n.LoadVersion() | Shrinking)
	nL.StoreVersion(nL.LoadVersion() | Growing)

	n.Left = nLR
	nL.Right = n

	if nP.Left == n {
		nP.Left = nL
	} else {
		nP.Right = nL
	}

	nL.Parent = nP
	n.Parent = nL

	if nLR != nil {
		nLR.Parent = n
	}

	height := 1 + avlNode.Max(nLR.Height(), n.Height())
	n.StoreHeight(height)
	nL.StoreHeight(1 + avlNode.Max(nL.Left.Height(), height))

	nL.StoreVersion(nL.LoadVersion() + GrowCountIncr)
	n.StoreVersion(n.LoadVersion() + ShrinkCountIncr)
}

func AttemptUpdate(avlNode *avlNode.AVLNode, value int) (int, int) {
	avlNode.Mutex.Lock()

	if avlNode.LoadVersion() == Unlinked {
		avlNode.Mutex.Unlock()
		return RETRY, 0
	}

	previousValue := avlNode.Value
	avlNode.Value = value
	avlNode.Mutex.Unlock()
	return SUCCESS, previousValue
}

func RotateLeft(n *avlNode.AVLNode) {
	nP := n.Parent
	nR := n.Right
	nRL := nR.Left

	nP.Mutex.Lock()
	defer nP.Mutex.Unlock()
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	nR.Mutex.Lock()
	defer nR.Mutex.Unlock()

	n.StoreVersion(n.LoadVersion() | Shrinking)
	nR.StoreVersion(nR.LoadVersion() | Growing)

	n.Right = nRL
	nR.Left = n

	if nP.Left == n {
		nP.Left = nR
	} else {
		nP.Right = nR
	}

	nR.Parent = nP
	n.Parent = nR

	if nRL != nil {
		nRL.Parent = n
	}

	height := 1 + avlNode.Max(nRL.Height(), n.Left.Height())
	n.StoreHeight(height)
	nR.StoreHeight(avlNode.Max(height, nR.Right.Height()))

	nR.StoreVersion(nR.LoadVersion() + GrowCountIncr)
	n.StoreVersion(n.LoadVersion() + ShrinkCountIncr)
}

func (avlTree *AVLTree) UnlinkNode(node *avlNode.AVLNode) {
	newNode := node

	if node.Right != nil && node.Left == nil {
		node.Right.Parent = node.Parent
		newNode = node.Right
	} else if node.Left != nil && node.Right == nil {
		node.Left.Parent = node.Parent
		newNode = node.Left
	}

	if node.Parent.Right == node {
		node.Parent.Right = newNode
	} else if node.Parent.Left == node {
		node.Parent.Left = newNode
	}

	node.Parent.StoreHeight(node.Parent.LoadHeight() - 1)
	avlTree.root.StoreHeight(avlTree.root.Right.Height() + 1)
}

func (avlTree *AVLTree) AttemptRemoveNode(par *avlNode.AVLNode, n *avlNode.AVLNode) (int, int) {
	if n == nil { // should be value???
		return NIL, 0
	}

	var prev int

	if n.CanUnlink() == 0 {
		n.Mutex.Lock()

		if n.LoadVersion() == Unlinked || n.CanUnlink() != 0 {
			n.Mutex.Unlock()
			return RETRY, 0
		}
		prev = n.Value
		n.Value = -1
	} else {
		par.Mutex.Lock()

		if par.LoadVersion() == Unlinked || n.Parent != par || n.LoadVersion() == Unlinked {
			par.Mutex.Unlock()
			return RETRY, 0
		}

		n.Mutex.Lock()

		prev = n.Value
		n.Value = -1

		if n.CanUnlink() != 0 {
			var c *avlNode.AVLNode

			if n.Left == nil {
				c = n.Right
			} else {
				c = n.Left
			}

			if par.Left == n {
				par.Left = c
			} else {
				par.Right = c
			}

			if c != nil {
				c.Parent = par
				n.StoreVersion(Unlinked)
			}
		}

		n.Mutex.Unlock()

		avlTree.root.StoreHeight(avlTree.root.Height())
		avlTree.FixHeightAndRotate(avlTree.root.Right)
	}

	return SUCCESS, prev
}

func (avlTree *AVLTree) AttemptRemove(value int, node *avlNode.AVLNode, direction int, nodeVersion int32) int {
	p := RETRY

	for {
		child := node.Child(direction)

		if ((node.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
			return RETRY
		}

		if child != nil {
			return NIL
		} else {
			nextDirection := Compare(value, child.Value)

			if nextDirection == 0 {
				_, p = avlTree.AttemptRemoveNode(node, child)
			} else {
				childVersion := child.LoadVersion()

				//	int a1 = isShrinking(chV);
				//	if (a1 != 0){
				//		waitUntilNotChanging(child);
				//	}
				//	else

				if childVersion != Unlinked && node.Child(direction) == child {
					if ((node.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
						return RETRY
					}

					p = avlTree.AttemptRemove(value, child, nextDirection, childVersion)
				}
			}
		}

		if p != RETRY {
			break
		}
	}

	return p
}

func (avlTree *AVLTree) Remove(value int) int {
	return avlTree.AttemptRemove(value, avlTree.root, -1, 0)
}

func main() {

}
