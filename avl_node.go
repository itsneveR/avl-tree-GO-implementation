package avl

type Node[T any] struct {
	key   uint
	value *T
	//bf    int // balance factor -1 , 0 , 1 == balanced
	h int //height = the maximal length of a path from the root to a leaf. a tree with one node has the height of 1
	L *Node[T]
	R *Node[T]
}

//we can define height -recursively- by saying that the
//empty tree has height 0, and the height of any node is
//one greater than the maximal height of its two children. == height invarient

// bf retruns balance factor of a node
func (n *Node[T]) bfactor() int {
	//[height of left subtree] - [height of right subtree]
	if n == nil {
		return -1
	}
	return n.L.Height() - n.R.Height()
}

func (n *Node[T]) Height() int {

	if n == nil {
		return 0
	} else {

		lh := n.L.Height()
		rh := n.R.Height()

		if rh > lh {
			return (rh + 1)
		} else {
			return (lh + 1)
		}

	}

}

func (n *Node[T]) IsLeaf() bool {
	return n.R == nil && n.L == nil
}

func (n *Node[T]) rightRotate() *Node[T] {
	x := n.L
	y := x.R

	x.R = n
	n.L = y

	return x
}

func (n *Node[T]) leftRotate() *Node[T] {
	x := n.R
	y := x.L

	x.L = n

	n.R = y

	return x
}

func (n *Node[T]) balance(k uint) *Node[T] {
	bf := n.bfactor()

	if (bf > 1) && (k < n.L.key) {
		return n.rightRotate()
	}

	if (bf > 1) && (k > n.L.key) {
		n.L = n.L.leftRotate()
		return n.rightRotate()
	}

	if (bf < -1) && (k > n.R.key) {
		return n.leftRotate()
	}

	if (bf < -1) && (k < n.R.key) {
		n.L = n.L.rightRotate()
		return n.leftRotate()
	}

	return n

}

func (n *Node[T]) rmBalance() *Node[T] {
	bf := n.bfactor()

	if (bf == 2) && (n.L.bfactor() >= 0) {
		return n.rightRotate()
	} else if (bf == 2) && (n.L.bfactor() == -1) {
		n.L = n.leftRotate()
		return n.rightRotate()
	} else if (bf == -2) && (n.R.bfactor() <= 0) {
		return n.leftRotate()
	} else if (bf == -2) && (n.R.bfactor() <= 1) {
		n.L = n.rightRotate()
		return n.leftRotate()
	}
	return n
}

func (n *Node[T]) Min() uint {
	if n.L == nil {
		return n.key
	}

	return n.L.Min()

}
