package tree

import (
	"cmp"
)

// TODO: Consider implementing a trait Comparable
type RBTree[T cmp.Ordered] struct {
	root *Node[T]
	count uint32
}

func NewRBTree[T cmp.Ordered]() *RBTree[T] {
	return &RBTree[T]{}
}

func (t *RBTree[T]) Search(data T) (T, error) {
	if (t.root.data == data) {
		return t.root.data, nil
	}
	
	
	var zero T
	return zero, ErrNodeNotFound
}

func (t *RBTree[T]) Insert(data T) error {
	node := NewNode(data)

	if (t.root == nil) {
		t.root = node
		return nil 
	}

	return nil
}

func (t *RBTree[T]) Delete(data T) error {
	return nil
}

func (t *RBTree[T]) String() string {
	return ""
}


