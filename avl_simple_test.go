package avl

import "testing"

func TestSimpleAVL_Has(t *testing.T) {
	a := NewSimple[int]()
	for i := range 10 {
		a.Insert(i)
	}
	for i := range 10 {
		if !a.Has(i) {
			t.Errorf("avl does not have %d", i)
			t.Logf("avl: %s", a)
		}
	}
}
