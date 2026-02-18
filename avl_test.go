package avl

import (
	"testing"
)

var cmpInt = func(a, b int) int { return a - b }

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

func TestAVL_Insert(t *testing.T) {
	a := NewInt()
	a.Insert(10)
	if a.root.heigth != 1 {
		t.Errorf("invalid heigth: got %d, wanted %d", a.root.heigth, 1)
		t.Logf("avl: %s", a)
	}
	a.Insert(5, 100)
	if a.root.heigth != 2 {
		t.Errorf("invalid heigth: got %d, wanted %d", a.root.heigth, 2)
		t.Logf("avl: %s", a)
	}
	a.Insert(1)
	if a.root.heigth != 3 {
		t.Errorf("invalid heigth: got %d, wanted %d", a.root.heigth, 3)
		t.Logf("avl: %s", a)
	}
	a = NewInt()
	for i := range 10 {
		a.Insert(i)
	}
	a = NewInt()
	for i := range 10 {
		a.Insert(9 - i)
	}
}

func TestAVL_Get(t *testing.T) {
	a := NewInt()
	for i := range 10 {
		a.Insert(i)
	}
	for i := range 10 {
		g := a.Get(func(v int) int { return v - i })
		if g == nil {
			t.Errorf("get not found, wanted %d", i)
		} else if *g != i {
			t.Errorf("invalid get: got: %d, wanted %d", *g, i)
		}
	}

}

func TestAVL_Delete(t *testing.T) {
	a := NewInt()
	for i := range 10 {
		a.Insert(i)
	}
	for i := range 10 {
		a.Delete(i)
		g := a.Get(func(v int) int { return v - i })
		if g != nil {
			t.Errorf("Get found %d, wanted not", *g)
			t.Logf("avl: %s", a)
		}
	}
}

func TestAVL_Min(t *testing.T) {
	a := NewInt()
	for i := range 10 {
		a.Insert(i)
	}
	println(a.String())
	for i := range 10 {
		got := a.Min()
		if got == nil {
			t.Errorf("invalid min: got nil, wanted %v", i)
			t.Logf("avl: %s", a)
		} else if i != *got {
			t.Errorf("invalid min: got %v, wanted %v", *got, i)
			t.Logf("avl: %s", a)
		}
		a.Delete(i)
	}
}

func TestAVL_Max(t *testing.T) {
	a := NewInt()
	for i := range 10 {
		a.Insert(i)
	}
	for i := range 10 {
		i = 9 - i
		got := a.Max()
		if got == nil {
			t.Errorf("invalid max: got nil, wanted %v", i)
			t.Logf("avl: %s", a)
		} else if i != *got {
			t.Errorf("invalid min: got %v, wanted %v", *got, i)
			t.Logf("avl: %s", a)
		}
		a.Delete(i)
	}
}
