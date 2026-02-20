package avl

import "fmt"

// Node is contained in an AVL.
//
// T should be a pointer to return nil if the value is not found.
type Node[T any] struct {
	Value  T
	heigth uint
	left   *Node[T]
	right  *Node[T]
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

func (n *Node[T]) rotate(cmp CompareFunc[T]) *Node[T] {
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

func (n *Node[T]) insert(v T, cmp CompareFunc[T]) *Node[T] {
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

func (n *Node[T]) delete(key T, cmp CompareFunc[T]) *Node[T] {
	comp := cmp(n.Value, key)
	res := n
	if comp == 0 {
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

// Clone the Node.
func (n *Node[T]) Clone(clone func(T) T) *Node[T] {
	node := newNode(clone(n.Value))
	if n.left != nil {
		node.left = n.left.Clone(clone)
	}
	if n.right != nil {
		node.right = n.right.Clone(clone)
	}
	return node
}
