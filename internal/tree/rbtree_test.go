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
