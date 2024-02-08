package avl

import (
	"sync/atomic"
)

type Tree[T any] struct {
	id   uint64
	root *Node[T]
}

func New[T any]() *Tree[T] {
	return &Tree[T]{
		id:   idgen(),
		root: new(Node[T]),
	}
}

func (tr *Tree[T]) Insert(k uint, v T) *Node[T] {
	return tr.insert(tr.root, k, v)
}

func (tr *Tree[T]) insert(n *Node[T], k uint, v T) *Node[T] {

	if n == nil {
		return &Node[T]{
			key:   k,
			value: &v,
		}
	}

	if k > n.key {
		n.R = tr.insert(n.R, k, v)
	} else if k < n.key {
		n.L = tr.insert(n.L, k, v)
	} else {
		return n
	}

	return n.balance(k)
}

func (tr *Tree[T]) Search(k uint) *Node[T] {
	return tr.search(tr.root, k)
}

func (tr *Tree[T]) search(n *Node[T], k uint) *Node[T] {
	//exit con
	if n == nil {
		return nil
	}

	if k > n.key {
		return tr.search(n.R, k)
	} else if k < n.key {
		return tr.search(n.L, k)
	}

	return n
}

func (tr *Tree[T]) Remove(k uint) *Node[T] {
	return tr.remove(tr.root, k)
}

func (tr *Tree[T]) remove(n *Node[T], k uint) *Node[T] {
	if n == nil {
		return n
	} //base/exit con

	if k > n.key {
		n.R = tr.remove(n.R, k)
	} else if k < n.key {
		n.L = tr.remove(n.L, k)
	}

	if k == n.key {

		if n.R == nil {
			return n.L
		} else if n.L == nil {
			return n.R
		} else if n.IsLeaf() {
			return nil
		} else {
			min := n.R.Min()
			n.key = min
			temp := tr.Search(min)
			n.value = temp.value
			n.R = tr.remove(n.R, min)
		}

	}
	return n.rmBalance()

}

func idgen() uint64 {
	var i uint64
	return atomic.AddUint64(&i, 1)
}
