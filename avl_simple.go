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

// NewKeySimple returns a new [KeyAVL] ordered with the standard function [cmp.Compare].
//
// If V is string, use [NewKeyString] instead.
func NewKeySimple[K cmp.Ordered, V any]() *KeyAVL[K, V] {
	return NewKey[K, V](cmp.Compare)
}

// NewKeySimpleImmutable returns a new [KeyAVL] storing immutable data ordered with the standard function [cmp.Compare].
func NewKeySimpleImmutable[K cmp.Ordered, V any](clone CloneFunc[V]) *KeyAVL[K, V] {
	return NewKeyImmutable[K](cmp.Compare, clone)
}

// NewKeySimpleMutable returns a new [KeyAVL] storing mutable data ordered with the standard function [cmp.Compare].
func NewKeySimpleMutable[K cmp.Ordered, V any]() *KeyAVL[K, V] {
	return NewKeyMutable[K, V](cmp.Compare)
}

// NewKeyString returns a new [KeyAVL] using strings as keys ordered with the standard function [strings.Compare].
func NewKeyString[V any]() *KeyAVL[string, V] {
	return NewKey[string, V](strings.Compare)
}
