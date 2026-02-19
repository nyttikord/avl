package avl

import "testing"

func TestNode_SimpleRotate(t *testing.T) {
	n := newNode(1)
	n.left = newNode(2)
	n.right = newNode(3)
	n.left.left = newNode(21)
	n.left.right = newNode(22)
	n.left.updateHeight()
	n.updateHeight()
	if n.heigth != 3 {
		t.Errorf("invalid heigth: got %v, wanted %v", n.heigth, 3)
	}
	n = n.rotateRight()
	if n.Value != 2 {
		t.Errorf("root is invalid: got %v, wanted %v", n.Value, 2)
	}
	if n.left.Value != 21 {
		t.Errorf("left is invalid: got %v, wanted %v", n.left.Value, 21)
	}
	if n.right.Value != 1 {
		t.Errorf("right is invalid: got %v, wanted %v", n.right.Value, 1)
	}
	if n.right.left.Value != 22 {
		t.Errorf("right left is invalid: got %v, wanted %v", n.right.left.Value, 22)
	}
	if n.right.right.Value != 3 {
		t.Errorf("right right is invalid: got %v, wanted %v", n.right.right.Value, 3)
	}
	n = n.rotateLeft()
	if n.Value != 1 {
		t.Errorf("root is invalid: got %v, wanted %v", n.Value, 1)
	}
	if n.left.Value != 2 {
		t.Errorf("left is invalid: got %v, wanted %v", n.left.Value, 2)
	}
	if n.right.Value != 3 {
		t.Errorf("right is invalid: got %v, wanted %v", n.right.Value, 3)
	}
	if n.left.left.Value != 21 {
		t.Errorf("right left is invalid: got %v, wanted %v", n.left.left.Value, 21)
	}
	if n.left.right.Value != 22 {
		t.Errorf("right right is invalid: got %v, wanted %v", n.left.right.Value, 22)
	}
}

func TestNode_Rotate(t *testing.T) {
	n := newNode(10)
	n.left = newNode(5)
	n.right = newNode(100)
	n.left.left = newNode(2)
	n.left.right = newNode(7)
	n.left.updateHeight()
	n.updateHeight()
	rotated := n.rotate(cmpInt)
	if n != rotated {
		t.Errorf("invalid rotation: got %s, wanted %s", rotated, n)
	}
	n.left.left.right = newNode(1)
	n.left.left.updateHeight()
	n.left.updateHeight()
	n.updateHeight()
	rotated = n.rotate(cmpInt)
	if n == rotated {
		t.Errorf("invalid rotation: did anything")
	}
}
