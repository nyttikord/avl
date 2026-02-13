package avl

// Comparable represents a comparable type.
type Comparable interface {
	// Compare returns 0 when current and c are the same.
	// Returns > 0 when current is bigger than c.
	// Returns < 0 when current is smaller than c.
	Compare(c Comparable) int
}

// Node is contained in an AVL.
//
// K is a comparable used to order nodes.
// V is the type value stored in the node.
type Node[V Comparable] struct {
	Value  V
	heigth uint
	left   *Node[V]
	right  *Node[V]
}

// AVL is a standard AVL tree containing Node.
type AVL[V Comparable] struct {
	root *Node[V]
}

// NewAVL creates a new AVL.
func NewAVL[V Comparable]() *AVL[V] {
	return &AVL[V]{}
}

// NewNode creates a new Node.
func NewNode[V Comparable]() *Node[V] {
	return &Node[V]{}
}

// Get returns the value associated with the key provided.
//
// If key is not found, it returns the default value of V.
// In this case, it returns nil if V is a pointer.
func (a *AVL[V]) Get(key Comparable) V {
	var v V
	node := a.root
	for node != nil {
		res := node.Value.Compare(key)
		if res == 0 {
			return node.Value
		} else if res > 0 {
			node = node.right
		} else {
			node = node.left
		}
	}
	return v
}

// Min returns the min contained in AVL.
func (a *AVL[V]) Min() V {
	node := a.root
	for node.left != nil {
		node = node.left
	}
	return node.Value
}

// Max returns the max contained in AVL.
func (a *AVL[V]) Max() V {
	node := a.root
	for node.right != nil {
		node = node.right
	}
	return node.Value
}

// Insert nodes in the AVL.
func (a *AVL[V]) Insert(nodes ...*Node[V]) *AVL[V] {
	for _, node := range nodes {
		if a.root == nil {
			a.root = node
		} else {
			a.root.insert(node)
		}
	}
	return a
}

// Delete nodes in the AVL.
func (a *AVL[V]) Delete(nodes ...*Node[V]) *AVL[V] {
	for _, node := range nodes {
		if a.root != nil {
			a.root = a.root.delete(node.Key)
		}
	}
	return a
}

func (n *Node[V]) updateHeight() {
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

func (n *Node[V]) rotateLeft() *Node[V] {
	newNode := n.left
	n.left = newNode.right
	newNode.right = n
	return newNode
}

func (n *Node[V]) rotateRight() *Node[V] {
	newNode := n.right
	n.right = newNode.left
	newNode.left = n
	return newNode
}

func (n *Node[V]) rotate() *Node[V] {
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
		if l.Value.Compare(n.Value) < 0 {
			n.left = l.rotateLeft()
		}
		return n.rotateRight()
	case -2:
		r := n.right
		if r.Value.Compare(n.Value) < 0 {
			n.left = r.rotateRight()
		}
		return n.rotateLeft()
	default:
		return n
	}
}

func (n *Node[V]) insert(node *Node[V]) *Node[V] {
	comp := n.Value.Compare(node.Value)
	if comp == 0 {
		n.Value = node.Value
		return n
	} else if comp > 0 {
		if node.left == nil {
			node.left = node
			return node
		}
		node.left = node.left.insert(node.left)
	} else {
		if node.right == nil {
			node.right = node
			return node
		}
		node.right = node.right.insert(node.right)
	}
	n.updateHeight()
	return n.rotate()
}

func (n *Node[V]) delete(key Comparable) *Node[V] {
	comp := n.Value.Compare(key)
	res := n
	if comp == 0 {
		if n.right == nil {
			if n.left == nil {
				return nil
			}
			res = n.left.delete(key)
		} else if n.left == nil {
			if n.right == nil {
				return nil
			}
			res = n.right.delete(key)
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
		return res.rotate()
	}
	return res
}
