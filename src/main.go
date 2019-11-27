package main

import (
	"sync/atomic"
)

type AVLNode struct {
	height  *int64
	version *int32
	value   int
	parent  *AVLNode
	left    *AVLNode
	right   *AVLNode
}

func (avlNode *AVLNode) LoadHeight() int64 {
	return atomic.LoadInt64(avlNode.height)
}

func (avlNode *AVLNode) StoreHeight(value int64) {
	atomic.StoreInt64(avlNode.height, value)
}

func (avlNode *AVLNode) LoadVersion() int32 {
	return atomic.LoadInt32(avlNode.version)
}

func (avlNode *AVLNode) StoreVersion(value int32) {
	atomic.StoreInt32(avlNode.version, value)
}

func (avlNode *AVLNode) Child(direction int) *AVLNode { // -1, 1
	switch direction {
	case -1:
		return avlNode.left
	case 1:
		return avlNode.right
	default:
		return nil
	}
}

// AVLNode END

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

func AttemptGet(value int, avlNode *AVLNode, direction int, nodeVersion int32) (int, int) { // -1 null, 0 retry, 1 found, value
	for {
		child := avlNode.Child(direction)

		if ((avlNode.LoadVersion() ^ nodeVersion) & IgnoreGrow) != 0 { // XOR
			return RETRY, 0
		}

		if child == nil {
			return NIL, 0
		}

		nextDirection := Compare(child.value, value)

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

var SpinCount int = 100

func WaitUntilNotChanging(avlNode *AVLNode) {
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

func AttemptPut(value int, avlNode *AVLNode, direction int, nodeVersion int32) (int, int) {
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
			nextDirection := Compare(child.value, value)

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

// func AttemptInsert(value int, avlNode *AVLNode, direction bool, nodeVersion int32) int {
// 	// critical section start

// 	if ((avlNode.LoadVersion()^nodeVersion)&IgnoreGrow) != 0 || avlNode.Child(direction) != nil {
// 		return RETRY
// 	} else {
// 		//
// 	}

// 	// critical section end
// }

func main() {

}
