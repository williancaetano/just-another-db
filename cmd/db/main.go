package main

import (
	"github.com/williancaetano/just-another-db/internal/tree"
)

func main() {
	testTree := tree.NewRBTree[int]()
	testTree.Insert(10)
	testTree.Insert(20)
	testTree.Insert(5)
	testTree.Insert(6)
	testTree.Insert(12)
	testTree.Insert(30)
	testTree.Insert(7)
	testTree.Insert(17)

	testTree.PrintTree()
}
