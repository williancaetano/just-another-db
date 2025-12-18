package tree

import (
	"strconv"
	"strings"
)

type Color byte

const (
	RedNode Color = iota + 1
	BlackNode
)

func (c *Color) String() string {
	switch *c {
	case RedNode:
		return "red"
	case BlackNode:
		return "black"
	}
	return strings.Join([]string{"unmapped color with value [", strconv.Itoa(int(*c)), "]"}, "")
}

type Ordering byte

const (
	PreOrder Ordering = iota + 1
	PostOrder
)
