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

// func (t *RBTree[T]) rebalance(n *Node[T]) error {
// 	if n.parent == n.parent.parent.left {
// 		uncle := n.parent.parent.right
// 	} else if n.parent == n.parent.parent.right {
// 		uncle := n.parent.parent.left
// 		if n.parent.color == RedNode && uncle.color == RedNode {
// 			n.parent.color = BlackNode
// 			uncle.color = BlackNode
// 			t.rebalance(n.parent.parent)
// 		}
// 	}
// 	// TODO: Implement Red black tree rebalancing after insertion
// 	return nil
// }

/*
Before Rotation:

    x
     \
      y
     / \
    a   b

After Left Rotation:

      y
     / \
    x   b
     \
      a
*/

func (t *RBTree[T]) rotateLeft(x *Node[T]) {
	// nomenclature references the 'before' (eh antes)
	y := x.right
	a := y.left

	// detach a and place it under x
	// 'a' will always be on the right of 'x', because 'y' was on the right of 'x' and 'a' was 'a' child of 'y'
	x.right = a
	if a != nil {
		a.parent = x
	}

	// rotate y upwards
	subTreeParent := x.parent
	if x.parent == nil {
		// x is root
		t.root = y
	} else if subTreeParent.right == x {
		// x is right child
		subTreeParent.right = y
	} else {
		// x is left child
		subTreeParent.left = y
	}
	y.parent = subTreeParent

	// rearrange 'x'
	y.left = x
	x.parent = y
}

/*
Befor Right Rotation:

      x
     /
    y
   / \
  a   b

After Right Rotation:

    y
   / \
  a   x
     /
    b
*/

func (t *RBTree[T]) rotateRight(x *Node[T]) {
	// nomenclature references the 'before' (eh antes)
	y := x.left
	b := y.right

	// detach a and place it under x
	// 'a' will always be on the right of 'x', because 'y' was on the right of 'x' and 'a' was 'a' child of 'y'
	x.left = b
	if b != nil {
		b.parent = x
	}

	// rotate y upwards
	subTreeParent := x.parent
	if x.parent == nil {
		// x is the root
		t.root = y
	} else if subTreeParent.right == x {
		// x is right child
		subTreeParent.right = y
	} else {
		// x is left child
		subTreeParent.left = y
	}
	y.parent = subTreeParent

	// rearrange 'x'
	y.right = x
	x.parent = y

}
