// Package avl contains the definition of the [AVL] tree.
//
// Use [New] to create a new [AVL].
// The data is treated as immutable if the type implements [Clonable].
// See [NewImmutable], [NewMutable] and [NewSimple] family functions.
//
// Use [NewKey] to create a new [KeyAVL].
// You can use this struct as a custom map.
// The data is treated as immutable if the type implements [Clonable].
// See [NewKeyImmutable], [NewKeyMutable] and [NewSimple] family functions.
package avl

// CompareFunc is a function that compares two values.
// Returns 0 if a = b, > 0 if a > b and < 0 if a < b.
type CompareFunc[T any] func(T, T) int

// CloneFunc is a function that clone the value.
// It is used to avoid side effects.
type CloneFunc[T any] func(T) T

func defaultCloneFunc[T any](t T) T { return t }

// AVL is a standard AVL tree containing nodes.
type AVL[T any] struct {
	root    *Node[T]
	compare CompareFunc[T]
	clone   CloneFunc[T]
	n       uint
}

// Clonable represents a type that will be automatically cloned when used in an [AVL].
type Clonable[T any] interface {
	Clone() T
}

func getClone[T any](v any) CloneFunc[T] {
	if _, ok := v.(T); !ok {
		panic("invalid usage of getClone")
	}
	if _, ok := v.(Clonable[T]); !ok {
		return defaultCloneFunc
	}
	return func(t T) T {
		var r any = t
		return r.(Clonable[T]).Clone()
	}
}

// NewImmutable creates a new [AVL] storing immutable data.
//
// clone is the function used to clone data to avoid side effects.
func NewImmutable[T any](cmp CompareFunc[T], clone CloneFunc[T]) *AVL[T] {
	return &AVL[T]{compare: cmp, clone: clone, n: 0}
}

// New creates a new [AVL].
//
// If T implements [Clonable], the inserted data becomes immutable.
// See [NewMutable] to avoid this behavior.
// See [NewImmutable] to use immutable data for types that don't implements Clonable.
func New[T any](cmp CompareFunc[T]) *AVL[T] {
	var t T
	return NewImmutable(cmp, getClone[T](t))
}

// NewMutable creates a new [AVL] storing mutable data.
func NewMutable[T any](cmp CompareFunc[T]) *AVL[T] {
	return NewImmutable(cmp, defaultCloneFunc)
}

// newNode creates a new [Node].
func newNode[T any](v T) *Node[T] {
	return &Node[T]{Value: v, heigth: 1}
}

func (a *AVL[T]) String() string {
	if a.root == nil {
		return "."
	}
	return a.root.String()
}

// Get returns the value associated with the key provided.
//
// cmp takes the checked value and returns an integer representing the order.
// It returns 0 if it is the value, > 0 if the checked value is bigger than the expected one and < 0 else.
//
//	// returns value 5 if it is present
//	a.Get(func(v int) { return a - 5 })
//
// If key is not found, it returns nil.
func (a *AVL[T]) Get(cmp func(v T) int) *T {
	node := a.root
	for node != nil {
		t := a.clone(node.Value)
		res := cmp(t)
		if res == 0 {
			tv := a.clone(node.Value)
			return &tv
		} else if res < 0 {
			node = node.right
		} else {
			node = node.left
		}
	}
	return nil
}

// Min returns the min contained in the [AVL].
func (a *AVL[T]) Min() *T {
	node := a.root
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	tv := a.clone(node.Value)
	return &tv
}

// Max returns the max contained in the [AVL].
func (a *AVL[T]) Max() *T {
	node := a.root
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	tv := a.clone(node.Value)
	return &tv
}

// Insert nodes in the [AVL].
func (a *AVL[T]) Insert(vals ...T) *AVL[T] {
	for _, v := range vals {
		v = a.clone(v)
		if a.root == nil {
			a.root = newNode(v)
		} else {
			a.root = a.root.insert(v, a.compare)
		}
		a.n++
	}
	return a
}

// Delete nodes in the [AVL].
func (a *AVL[T]) Delete(vals ...T) *AVL[T] {
	for _, v := range vals {
		if a.root != nil {
			a.root = a.root.delete(v, a.compare)
			a.n--
		}
	}
	return a
}

// Sort returns the sorted array of values contained in the [AVL].
func (a *AVL[T]) Sort() []T {
	arr := make([]T, a.n)
	sort(a.root, arr, 0)
	return arr
}

func sort[T any](n *Node[T], arr []T, i uint) uint {
	if n == nil {
		return i
	}
	i = sort(n.left, arr, i)
	arr[i] = n.Value
	return sort(n.right, arr, i+1)
}

// Clone the [AVL].
func (a *AVL[T]) Clone() *AVL[T] {
	tree := NewImmutable(a.compare, a.clone)
	tree.n = a.n
	tree.root = a.root.Clone(tree.clone)
	return tree
}

// Size the returns the number of nodes in the [AVL].
func (a *AVL[T]) Size() uint {
	return a.n
}
