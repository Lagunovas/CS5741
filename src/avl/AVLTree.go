package avlTree

import (
	"fmt"
)

var Unlinked int32 = 1
var Growing int32 = 2
var GrowCountIncrement int32 = 1 << 3
var GrowCountMask int32 = 0xFF << 3
var Shrinking int32 = 4
var ShrinkCountIncrement int32 = 1 << 11
var IgnoreGrow int32 = ^(Growing | GrowCountMask)

// CUSTOM
const RETRY = 0
const SUCCESS = 1
const NIL = -1
const EXISTING_VALUE = -2
const NOT_FOUND = -3

// UTILS START

func Compare(v0, v1 int) int {
	if v0 < v1 {
		return -1
	} else if v0 > v1 {
		return 1
	}

	return 0
}

func Abs(value int64) int64 {
	if value < 0 {
		return -value
	}

	return value
}

// UTILS END

type AVLTree struct {
	Root *AVLNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{NewAVLNode(-100)}
}

// =====================

func attemptGet(value, worker int, avlNode *AVLNode, direction int, nodeVersion int32) (int, int) { // -1 null, 0 retry, 1 found, value
	for {
		child := avlNode.Child(direction)

		if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 { // XOR
			return RETRY, 0
		}

		if child == nil {
			fmt.Println("worker", worker, "Value not found")
			return NIL, 0
		}

		nextDirection := Compare(value, child.Value)

		if nextDirection == 0 {
			//fmt.Println("Value found")
			return SUCCESS, value
		}

		childVersion := child.LoadVersion()

		if childVersion != Unlinked && child == avlNode.Child(direction) {
			if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
				return RETRY, 0
			} else {
				fmt.Printf("worker %d recursion of attempGet- next direction: %v\n", worker, nextDirection)
				status, value := attemptGet(value, worker, child, nextDirection, childVersion)

				if status != RETRY {
					//fmt.Println("Not retrying, return result")
					return status, value
				}
			}
		}

	}
}

func (avlTree *AVLTree) Get(value, worker int) int {
	fmt.Println("worker", worker, "called attemptGet")
	_, value = attemptGet(value, worker, avlTree.Root, 1, 0)
	return value
}

func (avlTree *AVLTree) attemptPut(value, worker int, avlNode *AVLNode, direction int, nodeVersion int32) (int, int) {
	p := 0 // -1 null, 0 retry, 1 found, value
	returnValue := value

	for {
		child := avlNode.Child(direction)

		if child != nil {
			if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
				return RETRY, returnValue
			}

			nextDirection := Compare(value, child.Value)

			if nextDirection == 0 {
				fmt.Println("worker ", worker, " attemptPut: value exists", value)
				return EXISTING_VALUE, returnValue
			} else {
				childVersion := child.LoadVersion()

				if childVersion != Unlinked && child == avlNode.Child(direction) {
					if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
						fmt.Println("retry 2")
						return RETRY, returnValue
					}
					p, returnValue = avlTree.attemptPut(value, worker, child, nextDirection, childVersion)
				}
			}
		} else {
			p = avlTree.attemptInsert(value, worker, avlNode, direction, nodeVersion)
		}

		if p != RETRY {
			break
		}
	}

	return p, returnValue
}

func (avlTree *AVLTree) Put(value, worker int) (int, int) {
	fmt.Println("worker ", worker, " attemptPut called")
	return avlTree.attemptPut(value, worker, avlTree.Root, 1, 0)
}

func Balance(n *AVLNode) int64 {
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

func NodeStatus(n *AVLNode) int64 {
	if n == nil {
		return NOCHANGE
	}

	if (n.Left == nil || n.Right == nil) && n.Value == -1 {
		return UNLINK
	}

	if n.Right != nil && n.Left != nil {
		if Abs(n.Right.LoadHeight()-n.Left.LoadHeight()) > 1 {
			return ROTATE
		}

		return NOCHANGE
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

func (avlTree *AVLTree) FixHeightAndRotate(node *AVLNode, worker int) {
	status := NodeStatus(node)

	switch status {
	case UNLINK:
		fmt.Println("worker", worker, "called FixHeightAndRotate => Unlink")
		avlTree.UnlinkNode(node)
	case ROTATE:
		fmt.Println("worker", worker, "called FixHeightAndRotate => Rotate")

		if node == nil {
			return
		}

		node.Mutex.Lock()
		defer node.Mutex.Unlock()

		balance := Balance(node)

		if balance >= 2 {
			if Balance(node.Left) < 0 {
				RotateLeft(node.Left, worker)
			}
			RotateRight(node, worker)
		} else if Balance(node) <= -2 {
			if Balance(node.Right) > 0 {
				RotateRight(node.Right, worker)
			}
			RotateLeft(node, worker)
		}
	}
}

func (avlTree *AVLTree) attemptInsert(value, worker int, avlNode *AVLNode, direction int, nodeVersion int32) int {
	avlNode.Mutex.Lock()

	firstCond := ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0

	if firstCond || avlNode.Child(direction) != nil {
		avlNode.Mutex.Unlock()
		return RETRY
	} else {
		newChild := NewAVLNode(value)
		avlNode.SetChild(direction, newChild)
		avlNode.Mutex.Unlock()
	}

	avlTree.Root.StoreHeight(avlTree.Root.Height())
	avlTree.FixHeightAndRotate(avlTree.Root.Right, worker)
	return SUCCESS
}

func RotateRight(n *AVLNode, worker int) {
	fmt.Println("worker", worker, "called RotateRight")
	nP := n.Parent
	nL := n.Left
	nLR := nL.Right

	nP.Mutex.Lock()
	defer nP.Mutex.Unlock()
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	nL.Mutex.Lock()
	defer nL.Mutex.Unlock()

	nV := n.LoadVersion()
	nLV := nL.LoadVersion()

	n.StoreVersion(nV | Shrinking)
	nL.StoreVersion(nLV | Growing)

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

	height := 1 + Max(nLR.Height(), n.Right.Height())
	n.StoreHeight(height)
	nL.StoreHeight(1 + Max(nL.Left.Height(), height))

	nL.StoreVersion(nL.LoadVersion() + GrowCountIncrement)
	n.StoreVersion(n.LoadVersion() + ShrinkCountIncrement)
}

func RotateLeft(n *AVLNode, worker int) {
	fmt.Println("worker", worker, "called RotateLeft ")
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

	height := 1 + Max(nRL.Height(), n.Left.Height())
	n.StoreHeight(height)
	nR.StoreHeight(Max(height, nR.Right.Height()))

	nR.StoreVersion(nR.LoadVersion() + GrowCountIncrement)
	n.StoreVersion(n.LoadVersion() + ShrinkCountIncrement)
}

func (avlTree *AVLTree) UnlinkNode(node *AVLNode) {
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
	avlTree.Root.StoreHeight(avlTree.Root.Right.Height() + 1)
}

func (avlTree *AVLTree) attemptRemoveNode(par *AVLNode, n *AVLNode, worker int) (int, int) {
	if n.Value == -1 {
		return NIL, -1
	}

	var prev int

	if !n.CanUnlink() {
		n.Mutex.Lock()

		if n.LoadVersion() == Unlinked || n.CanUnlink() {
			n.Mutex.Unlock()
			return RETRY, 0
		}
		prev = n.Value
		n.Value = -1
		n.Mutex.Unlock()
	} else {
		par.Mutex.Lock()

		if par.LoadVersion() == Unlinked || n.Parent != par || n.LoadVersion() == Unlinked {
			par.Mutex.Unlock()
			return RETRY, 0
		}

		n.Mutex.Lock()

		prev = n.Value
		n.Value = -1

		if n.CanUnlink() {
			var c *AVLNode

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
		par.Mutex.Unlock()

		avlTree.Root.StoreHeight(avlTree.Root.Height())
		avlTree.FixHeightAndRotate(avlTree.Root.Right, worker)
	}

	return SUCCESS, prev
}

func (avlTree *AVLTree) attemptRemove(value, worker int, node *AVLNode, direction int, nodeVersion int32) int {
	p := RETRY

	for {
		child := node.Child(direction)

		if ((node.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
			return RETRY
		}

		if child == nil {
			return NOT_FOUND
		} else {
			nextDirection := Compare(value, child.Value)

			if nextDirection == 0 {
				fmt.Println("worker", worker, " attemptRemove: found node, try to remove")
				fmt.Println("worker", worker, "attemptRemoveNode")
				_, p = avlTree.attemptRemoveNode(node, child, worker)
			} else {
				childVersion := child.LoadVersion()

				// if (childVersion & Shrinking) != 0 {
				// 	WaitUntilNotChanging(child)
				// } else
				if childVersion != Unlinked && node.Child(direction) == child {
					if ((node.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 {
						return RETRY
					}

					p = avlTree.attemptRemove(value, worker, child, nextDirection, childVersion)
				}
			}
		}

		if p != RETRY {
			break
		}
	}

	return p
}

func (avlTree *AVLTree) Remove(value, worker int) int {
	fmt.Println("Worker", worker, "called attemptRemove")
	return avlTree.attemptRemove(value, worker, avlTree.Root, 1, 0)
}

func (avlTree *AVLTree) PrintTree(avlNode *AVLNode) {
	if avlNode == nil {
		return
	} else {
		avlTree.PrintTree(avlNode.Left)
		fmt.Printf("Node => value: %v, height: %v\n", avlNode.Value, avlNode.LoadHeight())
		avlTree.PrintTree(avlNode.Right)
	}
}
