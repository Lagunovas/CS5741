package avlTree

import (
	"sync/atomic"

	recmutex "github.com/CS5741/src/misc"
)

type AVLNode struct {
	height  *int64
	version *int32
	Value   int
	Parent  *AVLNode
	Left    *AVLNode
	Right   *AVLNode
	Mutex   recmutex.RecursiveMutex
}

func NewAVLNode(value int) *AVLNode {
	return &AVLNode{height: new(int64), version: new(int32), Value: value}
}

func (avlNode *AVLNode) LoadHeight() int64 {
	return atomic.LoadInt64(avlNode.height)
	// return *avlNode.height
}

func (avlNode *AVLNode) StoreHeight(value int64) {
	atomic.StoreInt64(avlNode.height, value)
	// *avlNode.height = value
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}

	return y
}

func (avlNode *AVLNode) Height() int64 {
	if avlNode == nil {
		return 0
	}

	// needs sync, reetrant locks???

	avlNode.Mutex.Lock()
	defer avlNode.Mutex.Unlock()

	avlNode.StoreHeight(1 + Max(avlNode.Left.Height(), avlNode.Right.Height()))
	return avlNode.LoadHeight()
}

func (avlNode *AVLNode) LoadVersion() int32 {
	return atomic.LoadInt32(avlNode.version)
	// return *avlNode.version
}

func (avlNode *AVLNode) StoreVersion(value int32) {
	atomic.StoreInt32(avlNode.version, value)
	// *avlNode.version = value
}

func (avlNode *AVLNode) Child(direction int) *AVLNode { // -1, 1
	switch direction {
	case -1:
		return avlNode.Left
	case 1:
		return avlNode.Right
	default:
		return nil
	}
}

func (avlNode *AVLNode) SetChild(direction int, child *AVLNode) {
	avlNode.Mutex.Lock()
	defer avlNode.Mutex.Unlock()

	child.Mutex.Lock()
	defer child.Mutex.Unlock()

	switch direction {
	case -1:
		avlNode.Left = child
	case 1:
		avlNode.Right = child
	}

	child.Parent = avlNode
}

func (avlNode *AVLNode) CanUnlink() bool {
	return avlNode.Left == nil || avlNode.Right == nil
	// return avlNode == nil || avlNode.Left == nil || avlNode.Right == nil
	// if avlNode != nil {
	// 	if avlNode.Right == nil || avlNode.Left == nil {
	// 		return 1
	// 	}
	// 	return 0
	// }
	// return 1
}

func (avlNode *AVLNode) Leaf() bool {
	return avlNode.Left == nil && avlNode.Right == nil
}
