package avl

import(
	"fmt"
	"sync"
)

type AvlNode struct {
	height int 
	version int
	key int 
	value int 
	parent *AvlNode
	right *AvlNode
	left *AvlNode
}

type AvlTree struct {
	root *AvlNode
}
Retry := 0
Unlinked := 1
func (avl *AvlTree)remove(key int) (bool, int) {


	return false, 0
}

func canUnlink(node *AvlNode) bool {
	return node.left == nil || node.right == nil
}

func attemptRmNode(par *AvlNode, node *AvlNode) (bool, int) {
	var prev int 
	if !canUnlink(node){
		var l sync.Mutex
		l.Lock()
		if node.version == Unlinked || canUnlink(node) {
			return false, Retry 
		}
		prev  = node.value
		node.value  = nil
		l.Unlock()
	} else {
		var l sync.Mutex
		var ll sync.Mutex
		l.Lock()
			if par.version == Unlinked || node.parent != par || node.version == Unlinked {
				fmt.Println(par.version)
				fmt.Println(node.parent != par)
				fmt.Println(node.version)
				return false, Retry
			}
			
			ll.Lock()
				prev = node.value
				var c *AvlNode
				node.value = -1
				if canUnlink(node) {
					if node.left == nil {
						c = node.right 
					} else {
						c = node.left
					}
					if par.left == node {
						par.left = c 
					} else {
						par.right = c
					}

					if c != nil {
						c.parent = par
						node.version = Unlinked
					}
				}
			ll.Unlock()
		l.Unlock()
	}
	root.height = height(root)
	fixHeightAndRotate(root.right)
	return false, Retry

}
func fixHeightAndRotate(node *AvlAvlNode) {

	nodeParent := node.parent
	nodeLeft := node.left
	nodeLeftRight := nodeLeft.right  	
	var l sync.Mutex
	var ll sync.Mutex
	var lll sync.Mutex
	//is this really necessary 
	l.Lock()
		ll.Lock()
			lll.Lock()
			node.version |= Shrinking
			nodeLeft.version |= Growing

			node.left = nodeLeftRight
			nodeLeft,right = node 

			if nodeParent.left == node {
				nodeParent.left  = nodeLeft
			} else {
				nodeParent.right = left 
			}

			nodeLeft.Parent = nodeParent
			node.Parent = nodeLeft
			if nodeLeftRight != nil {
				nodeLeftRight.parent = node
			}
			h := 1 + Max(height(nodeLfetRight), height(node.right)) 
			node.height = h
			nodeLeft.height  = 1 + Max(height(nodeLeft), h)

			nodeLeft.version  += GrowCounterIncr
			node.version += ShrinkCountIncr

			lll.Unlock()
		ll.Unlock()
	l.Unlock()
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y 
}