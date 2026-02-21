package avl

import (
	"testing"
)

var cmpInt = func(a, b int) int { return a - b }

func TestAVL_Insert(t *testing.T) {
	a := NewOrdered[int]()
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
	a = NewOrdered[int]()
	for i := range 10 {
		a.Insert(i)
	}
	a = NewOrdered[int]()
	for i := range 10 {
		a.Insert(9 - i)
	}
}

func TestAVL_Get(t *testing.T) {
	a := NewOrdered[int]()
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
	a := NewOrdered[int]()
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
	a := NewOrdered[int]()
	for i := range 10 {
		a.Insert(i)
	}
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
	a := NewOrdered[int]()
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

func TestAVL_Sort(t *testing.T) {
	a := NewOrdered[int]()
	for i := range 10 {
		a.Insert(i)
	}
	st := a.Sort()
	for i := range 10 {
		if st[i] != i {
			t.Errorf("invalid value: got %d, wanted %d", st[i], i)
		}
	}
}

func TestAVL_Clone(t *testing.T) {
	a := NewOrdered[int]()
	for i := range 10 {
		a.Insert(i)
	}
	cl := a.Clone()
	a.Insert(11)
	got := cl.Get(func(v int) int { return v - 11 })
	if got != nil {
		t.Errorf("invalid value: got %d, wanted nothing", *got)
		t.Logf("avl: %s", a)
		t.Logf("clone: %s", cl)
	}
}

type cloneStruct []int

func (c cloneStruct) Clone() cloneStruct {
	dst := make([]int, len(c))
	copy(dst, c)
	return dst
}

func TestClonable_AutoClone(t *testing.T) {
	a := New(func(a, b cloneStruct) int { return len(a) - len(b) })
	cl := cloneStruct{1, 2, 3}
	a.Insert(cl)
	got := a.Get(func(v cloneStruct) int { return 3 - len(v) })
	if got == nil {
		t.Errorf("invalid value: got nil")
		t.Logf("avl: %s", a)
		return
	}
	for i := range *got {
		(*got)[i] = (i + 1) << 3
	}
	ng := a.Get(func(v cloneStruct) int { return 3 - len(v) })
	for i, r := range *ng {
		if i+1 != r {
			t.Errorf("invalid value: got %d, wanted %d", r, i+1)
			t.Logf("avl: %s", a)
		}
	}
}
