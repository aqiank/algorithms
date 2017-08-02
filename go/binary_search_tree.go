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

func (b *binaryNode) add(value int) {
	if b.value < value {
		if b.right != nil {
			b.right.add(value)
		} else {
			b.right = newBinaryNode(value)
		}
	} else {
		if b.left != nil {
			b.left.add(value)
		} else {
			b.left = newBinaryNode(value)
		}
	}
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
