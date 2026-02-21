package avl

import (
	"cmp"
	"strings"
)

// SimpleAVL represents an [AVL] that stores simple values.
type SimpleAVL[T any] struct{ *AVL[T] }

// NewSimple returns a new [AVL] ordered with the standard function [cmp.Compare].
//
// If T is string, use [NewString] instead.
func NewSimple[T cmp.Ordered]() *SimpleAVL[T] {
	return &SimpleAVL[T]{New(cmp.Compare[T])}
}

// NewSimpleImmutable returns a new [AVL] storing immutable data ordered with the standard function [cmp.Compare].
func NewSimpleImmutable[T cmp.Ordered](clone CloneFunc[T]) *SimpleAVL[T] {
	return &SimpleAVL[T]{NewImmutable(cmp.Compare[T], clone)}
}

// NewSimpleMutable returns a new [AVL] storing mutable data ordered with the standard function [cmp.Compare].
func NewSimpleMutable[T cmp.Ordered]() *SimpleAVL[T] {
	return &SimpleAVL[T]{NewMutable(cmp.Compare[T])}
}

// NewString returns a new [AVL] storing strings ordered with the standard function [strings.Compare].
func NewString() *SimpleAVL[string] {
	return &SimpleAVL[string]{New(strings.Compare)}
}

// Has returns true if the value is in the [SimpleAVL].
func (a *SimpleAVL[T]) Has(s T) bool {
	return a.Get(func(v T) int { return a.compare(v, s) }) != nil
}
