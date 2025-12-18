package tree

import (
	"cmp"
	"fmt"
)

type NodeData interface {
	cmp.Ordered
}

type Node[T NodeData] struct {
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
	color  Color
	data   T
}

func (n *Node[T]) walk() {
	if n == nil {
		return
	}

	fmt.Println(n.String())
	n.left.walk()
	n.right.walk()
}

func (n *Node[T]) String() string {
	if n == nil {
		return "nill"
	}
	return fmt.Sprintf("Current: [%v] Parent [%p] Left: [%p] Right: [%p] Color: [%s]", n.data, n.parent, n.left, n.right, n.color.String())
}

type nodeOptions[T cmp.Ordered] func(*Node[T]) error

func WithParent[T cmp.Ordered](parent *Node[T]) nodeOptions[T] {
	return func(n *Node[T]) error {
		n.parent = parent
		return nil
	}
}

func WithLeft[T cmp.Ordered](left *Node[T]) nodeOptions[T] {
	return func(n *Node[T]) error {
		n.left = left
		return nil
	}
}

func WithRight[T cmp.Ordered](right *Node[T]) nodeOptions[T] {
	return func(n *Node[T]) error {
		n.right = right
		return nil
	}
}

func WithColor[T cmp.Ordered](color Color) nodeOptions[T] {
	return func(n *Node[T]) error {
		if color != RedNode && color != BlackNode {
			return ErrInvalidColor
		}
		n.color = color
		return nil
	}
}

func NewNode[T cmp.Ordered](data T, options ...nodeOptions[T]) (*Node[T], error) {
	node := &Node[T]{
		data: data,
	}

	for _, option := range options {
		if err := option(node); err != nil {
			return nil, err
		}
	}

	return node, nil
}
