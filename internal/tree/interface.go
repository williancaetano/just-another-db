package tree

import "cmp"

type Tree[T cmp.Ordered] interface {
	Search(data T) (*Node[T], error)
	Insert(data T) error
	Delete(data T) error
}
