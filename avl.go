package avl

// Compare is a function that compares two values.
// Returns 0 if a = b, > 0 if a > b and < 0 if a < b.
type Compare[T any] func(a, b T) int

// AVL is a standard AVL tree containing Node.
type AVL[T any] struct {
	root    *Node[T]
	compare Compare[T]
	clone   func(T) T
}

// NewClone creates a new AVL that clone inserted data to avoid side effects.
func NewClone[T any](cmp Compare[T], clone func(T) T) *AVL[T] {
	return &AVL[T]{compare: cmp, clone: clone}
}

// New creates a new AVL.
func New[T any](cmp Compare[T]) *AVL[T] {
	return NewClone(cmp, func(t T) T { return t })
}

// NewInt creates a new AVL storing int.
func NewInt() *AVL[int] {
	return New(func(a, b int) int { return a - b })
}

// newNode creates a new Node.
func newNode[T any](v T) *Node[T] {
	return &Node[T]{Value: v, heigth: 1}
}

func (a *AVL[T]) String() string {
	if a.root == nil {
		return ""
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

// Min returns the min contained in AVL.
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

// Max returns the max contained in AVL.
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

// Insert nodes in the AVL.
func (a *AVL[T]) Insert(vals ...T) *AVL[T] {
	for _, v := range vals {
		v = a.clone(v)
		if a.root == nil {
			a.root = newNode(v)
		} else {
			a.root = a.root.insert(v, a.compare)
		}
	}
	return a
}

// Delete nodes in the AVL.
func (a *AVL[T]) Delete(vals ...T) *AVL[T] {
	for _, v := range vals {
		if a.root != nil {
			a.root = a.root.delete(v, a.compare)
		}
	}
	return a
}

func (a *AVL[T]) Sort() []T {
	return sort(a.root)
}

func sort[T any](n *Node[T]) (v []T) {
	if n == nil {
		return
	}
	return append(append(sort(n.left), n.Value), sort(n.right)...)
}
