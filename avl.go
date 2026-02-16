package avl

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
//
// T should be a pointer to return nil if the value is not found.
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

// NewNode creates a new Node.
func NewNode[T any]() *Node[T] {
	return &Node[T]{}
}

// Get returns the value associated with the key provided.
//
// cmp takes the checked value and returns an integer representing the order.
// It returns 0 if it is the value, > 0 if the checked value is bigger than the expected one and < 0 else.
//
//	// returns value 5 if it is present
//	a.Get(func(v int) { return a - 5 })
//
// If key is not found, it returns the default value of V (nil if V is a pointer).
func (a *AVL[T]) Get(cmp func(v T) int) T {
	var v T
	node := a.root
	for node != nil {
		res := cmp(node.Value)
		if res == 0 {
			return node.Value
		} else if res < 0 {
			node = node.right
		} else {
			node = node.left
		}
	}
	return v
}

// Min returns the min contained in AVL.
func (a *AVL[T]) Min() T {
	node := a.root
	for node.left != nil {
		node = node.left
	}
	return node.Value
}

// Max returns the max contained in AVL.
func (a *AVL[T]) Max() T {
	node := a.root
	for node.right != nil {
		node = node.right
	}
	return node.Value
}

// Insert nodes in the AVL.
func (a *AVL[T]) Insert(nodes ...*Node[T]) *AVL[T] {
	for _, node := range nodes {
		if a.root == nil {
			a.root = node
		} else {
			a.root.insert(node, a.compare)
		}
	}
	return a
}

// Delete nodes in the AVL.
func (a *AVL[T]) Delete(nodes ...*Node[T]) *AVL[T] {
	for _, node := range nodes {
		if a.root != nil {
			a.root = a.root.delete(node.Value, a.compare)
		}
	}
	return a
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
	newNode := n.left
	n.left = newNode.right
	newNode.right = n
	return newNode
}

func (n *Node[T]) rotateRight() *Node[T] {
	newNode := n.right
	n.right = newNode.left
	newNode.left = n
	return newNode
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
	comp := left - right
	switch comp {
	case 2:
		l := n.left
		if cmp(l.Value, n.Value) < 0 {
			n.left = l.rotateLeft()
		}
		return n.rotateRight()
	case -2:
		r := n.right
		if cmp(r.Value, n.Value) < 0 {
			n.left = r.rotateRight()
		}
		return n.rotateLeft()
	default:
		return n
	}
}

func (n *Node[T]) insert(node *Node[T], cmp Compare[T]) *Node[T] {
	comp := cmp(n.Value, node.Value)
	if comp == 0 {
		n.Value = node.Value
		return n
	} else if comp > 0 {
		if node.left == nil {
			node.left = node
			return node
		}
		node.left = node.left.insert(node.left, cmp)
	} else {
		if node.right == nil {
			node.right = node
			return node
		}
		node.right = node.right.insert(node.right, cmp)
	}
	n.updateHeight()
	return n.rotate(cmp)
}

func (n *Node[T]) delete(key T, cmp Compare[T]) *Node[T] {
	comp := cmp(n.Value, key)
	res := n
	if comp == 0 {
		if n.right == nil {
			if n.left == nil {
				return nil
			}
			res = n.left.delete(key, cmp)
		} else if n.left == nil {
			if n.right == nil {
				return nil
			}
			res = n.right.delete(key, cmp)
		} else {
			cur := n
			for cur.left != nil {
				cur = cur.left
			}
			res = cur
		}
	}
	if res != nil {
		res.updateHeight()
		return res.rotate(cmp)
	}
	return res
}
