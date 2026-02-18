package avl

import "fmt"

// Compare is a function that compares two values.
// Returns 0 if a = b, > 0 if a > b and < 0 if a < b.
type Compare[T any] func(a, b T) int

// Node is contained in an AVL.
//
// T should be a pointer to return nil if the value is not found.
type Node[T any] struct {
	Value  T
	heigth uint
	left   *Node[T]
	right  *Node[T]
}

// AVL is a standard AVL tree containing Node.
type AVL[T any] struct {
	root    *Node[T]
	compare Compare[T]
}

// NewAVL creates a new AVL.
func NewAVL[T any](cmp Compare[T]) *AVL[T] {
	return &AVL[T]{compare: cmp}
}

// NewAVLInt creates a new AVL storing int.
func NewAVLInt() *AVL[int] {
	return &AVL[int]{
		compare: func(a, b int) int { return a - b },
	}
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
func (a *AVL[T]) Get(cmp func(v T) int) (v *T) {
	node := a.root
	for node != nil {
		res := cmp(node.Value)
		if res == 0 {
			return &node.Value
		} else if res < 0 {
			node = node.right
		} else {
			node = node.left
		}
	}
	return
}

// Min returns the min contained in AVL.
func (a *AVL[T]) Min() (v *T) {
	node := a.root
	if node == nil {
		return
	}
	for node.left != nil {
		node = node.left
	}
	return &node.Value
}

// Max returns the max contained in AVL.
func (a *AVL[T]) Max() (v *T) {
	node := a.root
	if node == nil {
		return
	}
	for node.right != nil {
		node = node.right
	}
	return &node.Value
}

// Insert nodes in the AVL.
func (a *AVL[T]) Insert(vals ...T) *AVL[T] {
	for _, v := range vals {
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

func (n *Node[T]) String() string {
	left := `""`
	if n.left != nil {
		left = n.left.String()
	}
	right := `""`
	if n.right != nil {
		right = n.right.String()
	}
	if n.right == n.left {
		return fmt.Sprintf("%v", n.Value)
	}
	return fmt.Sprintf("{%d}%v: [%s, %s]", n.heigth, n.Value, left, right)
}

func (n *Node[T]) updateHeight() {
	left := uint(0)
	if n.left != nil {
		left = n.left.heigth
	}
	right := uint(0)
	if n.right != nil {
		right = n.right.heigth
	}
	n.heigth = 1 + max(left, right)
}

func (n *Node[T]) rotateLeft() *Node[T] {
	next := n.right
	n.right = next.left
	next.left = n
	n.updateHeight()
	next.updateHeight()
	return next
}

func (n *Node[T]) rotateRight() *Node[T] {
	next := n.left
	n.left = next.right
	next.right = n
	n.updateHeight()
	next.updateHeight()
	return next
}

func (n *Node[T]) rotate(cmp Compare[T]) *Node[T] {
	left := 0
	if n.left != nil {
		left = int(n.left.heigth)
	}
	right := 0
	if n.right != nil {
		right = int(n.right.heigth)
	}
	switch left - right {
	case 2:
		l := n.left
		if cmp(l.Value, n.Value) > 0 {
			n.left = l.rotateLeft()
		}
		return n.rotateRight()
	case -2:
		r := n.right
		if cmp(r.Value, n.Value) < 0 {
			n.right = r.rotateRight()
		}
		return n.rotateLeft()
	default:
		return n
	}
}

func (n *Node[T]) insert(v T, cmp Compare[T]) *Node[T] {
	comp := cmp(n.Value, v)
	var next **Node[T]
	if comp == 0 {
		n.Value = v
		return n
	} else if comp > 0 {
		next = &n.left
	} else {
		next = &n.right
	}
	if *next == nil {
		*next = newNode(v)
	}
	*next = (*next).insert(v, cmp)
	n.updateHeight()
	return n.rotate(cmp)
}

func (n *Node[T]) delete(key T, cmp Compare[T]) *Node[T] {
	comp := cmp(n.Value, key)
	res := n
	if comp == 0 {
		println("found :D", n.Value)
		if n.right == nil {
			res = n.left
		} else if n.left == nil {
			res = n.right
		} else {
			cur := n
			for cur.left != nil {
				cur = cur.left
			}
			res = cur
		}
	} else if comp < 0 {
		if n.right == nil {
			return nil
		}
		n.right = n.right.delete(key, cmp)
	} else {
		if n.left == nil {
			return nil
		}
		n.left = n.left.delete(key, cmp)
	}
	if res != nil {
		res.updateHeight()
		return res.rotate(cmp)
	}
	return res
}
