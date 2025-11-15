package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInEmptyTree(t *testing.T) {
	tree := RBTree[int]{}
	_, err := tree.Search(10)

	assert.NotNil(t, err, "should've returned a ErrNotFound")
}

func TestInsertSearchOnRoot(t *testing.T) {
	tree := RBTree[int]{}

	err := tree.Insert(10)
	assert.Nil(t, err, "insert should have not errored")

	found, err := tree.Search(10)
	assert.Nil(t, err, "search should have not errored")
	assert.NotNil(t, found, "search should have not found something")
	assert.Equal(t, found, tree.root, "what was found is the tree's root")

}
