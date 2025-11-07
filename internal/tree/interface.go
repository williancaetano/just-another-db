package tree

import "cmp"

type Tree[T cmp.Ordered] interface {
	Search(data T) (T, error)
	Insert(data T) error
	Delete(data T) error
}