package tree

import (
	"cmp"
	"fmt"
)

// TODO: Consider implementing a trait Comparable
type RBTree[T NodeData] struct {
	root  *Node[T]
	count uint32
}

var _ Tree[int] = (*RBTree[int])(nil)

func NewRBTree[T cmp.Ordered]() *RBTree[T] {
	return &RBTree[T]{}
}

func (t *RBTree[T]) Search(data T) (*Node[T], error) {
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
	var prev *Node[T]

	node, err := NewNode(data, WithColor[T](BlackNode))
	if err != nil {
		return err
	}
	if t.root == nil {
		t.root = node
		t.count++
		return nil
	}
	curr := t.root
	for curr != nil {
		prev = curr
		if data < curr.data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	node.parent = prev
	if data > prev.data {
		prev.right = node
	} else {
		prev.left = node
	}

	t.count++
	return nil
}

func (t *RBTree[T]) Delete(data T) error {
	return nil
}

func (t *RBTree[T]) String() string {
	fmt.Println(t.root.String())

	return "unimplemented"
}

func (t *RBTree[T]) PrintTree() {
	t.root.walk()
}

// func (t *RBTree[T]) Iter(order Ordering) func(yield func(T) bool) {

// }

func (t *RBTree[T]) rebalance(*Node[T]) error {
	//TODO: Implement Red black tree rebalancing after insertion
	return nil
}

func leftRotation[T cmp.Ordered](x *Node[T]) *Node[T], error {

}
