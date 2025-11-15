package tree

import (
	"cmp"
)

// TODO: Consider implementing a trait Comparable
type RBTree[T cmp.Ordered] struct {
	root  *Node[T]
	count uint32
}

var _ Tree[int] = (*RBTree[int])(nil)

func NewRBTree[T cmp.Ordered]() *RBTree[T] {
	return &RBTree[T]{}
}

func (t *RBTree[T]) Search(data T) (*Node[T], error) {
	if t.root.data == data {
		return t.root, nil
	}

	curr := t.root
	for curr != nil {
		if curr.data == data {
			return curr, nil
		} else if data < curr.data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return nil, ErrNodeNotFound
}

func (t *RBTree[T]) Insert(data T) error {
	var err error

	if t.root == nil {
		t.root, err = NewNode(data, WithColor[T](BlackNode))
		if err != nil {
			return err
		}
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
