package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNodeWithValidParams(t *testing.T) {
	_, err := NewNode(10, WithColor[int](RedNode))
	assert.Nil(t, err, "node shouln't have failed")

	_, err = NewNode(10, WithColor[int](BlackNode))
	assert.Nil(t, err, "node shouln't have failed")

	_, err = NewNode(10,
		WithColor[int](BlackNode),
		WithLeft[int](nil),
		WithRight[int](nil),
		WithParent[int](nil),
	)
	assert.Nil(t, err, "node shouln't have failed")
}

func TestCreateNodeWithInvalidColor(t *testing.T) {
	_, err := NewNode(10, WithColor[int](RedNode|BlackNode))
	assert.NotNil(t, err, "node should have failed")
}
