package tree

import "errors"

var ErrNodeNotFound = errors.New("value not found")
var ErrUnexpected = errors.New("we fucked up, sorry :<")

var ErrInvalidColor = errors.New("invalid color for node, must be RedNode or BlackNode")
