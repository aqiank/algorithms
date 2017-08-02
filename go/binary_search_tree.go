package main

type binaryNode struct {
	left   *binaryNode
	right  *binaryNode
	value  int
	height int
}

func newBinaryNode(value int) *binaryNode {
	return &binaryNode{
		left:  nil,
		right: nil,
		value: value,
	}
}

func (b *binaryNode) add(value int) *binaryNode {
	newRoot := b
	if value > b.value {
		if b.right != nil {
			b.right = b.addToSubTree(b.right, value)
			if b.heightDifference() == 2 {
				if value > b.left.value {
					newRoot = b.rotateRight()
				} else {
					newRoot = b.rotateLeftRight()
				}
			}
		} else {
			b.right = newBinaryNode(value)
		}
	} else {
		if b.left != nil {
			b.left = b.addToSubTree(b.left, value)
		} else {
			b.left = newBinaryNode(value)
		}
	}
	return newRoot
}

func (b *binaryNode) addToSubTree(parent *binaryNode, value int) *binaryNode {
	if parent == nil {
		return newBinaryNode(value)
	}
	return parent.add(value)
}

func (b *binaryNode) contains(value int) bool {
	if b.value != value {
		if b.right != nil && value > b.value {
			return b.right.contains(value)
		} else if b.left != nil {
			return b.left.contains(value)
		} else {
			return false
		}
	}

	return true
}

func (b *binaryNode) computeHeight() {
	height := -1
	if b.left != nil && b.left.height > height {
		height = b.left.height
	}
	if b.right != nil && b.right.height > height {
		height = b.right.height
	}
	b.height = height + 1
}

func (b *binaryNode) heightDifference() int {
	leftTarget := 0
	rightTarget := 0
	if b.left != nil {
		leftTarget = 1 + b.left.height
	}
	if b.right != nil {
		rightTarget = 1 + b.right.height
	}
	return leftTarget - rightTarget
}

func (b *binaryNode) rotateRight() *binaryNode {
	newRoot := b.left
	if newRoot.right != nil {
		b.left = newRoot.right
	}
	newRoot.right = b
	return newRoot
}

func (b *binaryNode) rotateLeft() *binaryNode {
	newRoot := b.right
	if newRoot.left != nil {
		b.right = newRoot.left
	}
	newRoot.left = b
	return newRoot
}

func (b *binaryNode) rotateLeftRight() *binaryNode {
	newRoot := b.left.rotateLeft()
	newRoot = b.rotateRight()
	return newRoot
}

func (b *binaryNode) rotateRightLeft() *binaryNode {
	newRoot := b.right.rotateRight()
	newRoot = b.rotateLeft()
	return newRoot
}
