package tree

import "cmp"

type Node[T cmp.Ordered] struct {
	data	T
	left	*Node[T]
	right	*Node[T]
}

func NewNode[T cmp.Ordered](data T) *Node[T] {
	return &Node[T]{
		data:  data,
		left:  nil,
		right: nil,
	}
}
