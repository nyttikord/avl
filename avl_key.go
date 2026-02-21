package avl

import "fmt"

// KeyAVL is an [AVL] where the comparison is made on keys (K).
type KeyAVL[K, V any] struct {
	avl     *AVL[wrapper[K, V]]
	compare CompareFunc[K]
	clone   CloneFunc[V]
}

type wrapper[K, V any] struct {
	key   K
	data  V
	clone CloneFunc[V]
}

func unwrap[K, V any](v *wrapper[K, V]) *V {
	if v == nil {
		return nil
	}
	return &v.data
}

// Clone doesn't clone the key.
func (w wrapper[K, V]) Clone() wrapper[K, V] {
	return wrapper[K, V]{
		data:  w.clone(w.data),
		key:   w.key,
		clone: w.clone,
	}
}

func (w wrapper[K, V]) String() string {
	return fmt.Sprintf("%v:%v", w.key, w.data)
}

// NewKeyImmutable creates a new [AVL] storing immutable data.
//
// clone is the function used to clone data to avoid side effects.
func NewKeyImmutable[K, V any](cmp CompareFunc[K], clone CloneFunc[V]) *KeyAVL[K, V] {
	return &KeyAVL[K, V]{
		avl:     New(func(w1, w2 wrapper[K, V]) int { return cmp(w1.key, w2.key) }),
		compare: cmp,
		clone:   clone,
	}
}

// NewKey returns a new [KeyAVL].
//
// If V implements [Clonable], the inserted data becomes immutable.
// See [NewKeyMutable] to avoid this behavior.
// See [NewKeyImmutable] to use immutable data for types that don't implements Clonable.
func NewKey[K, V any](cmp CompareFunc[K]) *KeyAVL[K, V] {
	var v V
	return NewKeyImmutable(cmp, getClone[V](v))
}

// NewKeyMutable creates a new [AVL] storing mutable data.
func NewKeyMutable[T any](cmp CompareFunc[T]) *AVL[T] {
	return NewImmutable(cmp, defaultCloneFunc)
}

func (a *KeyAVL[K, V]) String() string {
	return a.avl.String()
}

// Get returns the value associated with the key provided.
//
// If key is not found, it returns nil.
func (a *KeyAVL[K, V]) Get(k K) *V {
	return unwrap(a.avl.Get(func(v wrapper[K, V]) int { return a.compare(v.key, k) }))
}

// Has returns true if the value is in the [KeyAVL].
func (a *KeyAVL[K, V]) Has(k K) bool {
	return a.Get(k) != nil
}

// Min returns the min contained in the [KeyAVL].
func (a *KeyAVL[K, V]) Min() *V {
	return unwrap(a.avl.Min())
}

// Max returns the max contained in the [KeyAVL].
func (a *KeyAVL[K, V]) Max() *V {
	return unwrap(a.avl.Max())
}

// Insert one node in the [KeyAVL].
func (a *KeyAVL[K, V]) Insert(k K, v V) *KeyAVL[K, V] {
	a.avl.Insert(wrapper[K, V]{k, v, a.clone})
	return a
}

// Delete nodes in the [KeyAVL].
func (a *KeyAVL[K, V]) Delete(keys ...K) *KeyAVL[K, V] {
	for _, k := range keys {
		a.avl.Delete(wrapper[K, V]{key: k})
	}
	return a
}

// Sort returns the sorted array of values contained in the [KeyAVL].
func (a *KeyAVL[K, V]) Sort() []V {
	s := a.avl.Sort()
	sorted := make([]V, len(s))
	for i, v := range s {
		sorted[i] = v.data
	}
	return sorted
}

// Clone the [KeyAVL].
func (a *KeyAVL[K, V]) Clone() *KeyAVL[K, V] {
	tree := NewKey[K, V](a.compare)
	tree.avl = a.avl.Clone()
	return tree
}

// Size the returns the number of nodes in the [KeyAVL].
func (a *KeyAVL[K, V]) Size() uint {
	return a.avl.Size()
}
