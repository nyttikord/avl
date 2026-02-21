package avl

import (
	"cmp"
	"strings"
)

// NewOrdered returns a new AVL ordered with the standard function [cmp.Compare].
//
// If T is string, use [NewString] instead.
func NewOrdered[T cmp.Ordered]() *AVL[T] {
	return New(cmp.Compare[T])
}

// NewOrderedImmutable returns a new AVL storing immutable data ordered with the standard function [cmp.Compare].
func NewOrderedImmutable[T cmp.Ordered](clone CloneFunc[T]) *AVL[T] {
	return NewImmutable(cmp.Compare[T], clone)
}

// NewOrderedMutable returns a new AVL storing mutable data ordered with the standard function [cmp.Compare].
func NewOrderedMutable[T cmp.Ordered]() *AVL[T] {
	return NewMutable(cmp.Compare[T])
}

// NewString returns a new AVL storing strings ordered with the standard function [strings.Compare].
func NewString() *AVL[string] {
	return New(strings.Compare)
}
