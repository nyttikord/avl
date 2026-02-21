package avl

import (
	"cmp"
	"strings"
	"testing"
)

func TestKeyAVL(t *testing.T) {
	a := NewKey[int, string](cmp.Compare)
	a.Insert(1, "hello").
		Insert(2, "world").
		Insert(3, "!")
	got := a.Get(1)
	if got == nil {
		t.Errorf("invalid value: got nil")
		t.Logf("avl: %s", a)
	} else if *got != "hello" {
		t.Errorf("invalid value: got %s, wanted %q", *got, "hello")
		t.Logf("avl: %s", a)
	}
	st := a.Sort()
	if strings.Join(st, " ") != "hello world !" {
		t.Errorf("invalid value: got %s, wanted %q", strings.Join(st, " "), "hello world !")
		t.Logf("avl: %s", a)
	}
	a.Delete(2)
	st = a.Sort()
	if strings.Join(st, " ") != "hello !" {
		t.Errorf("invalid value: got %s, wanted %q", strings.Join(st, " "), "hello !")
		t.Logf("avl: %s", a)
	}
	if !a.Has(1) {
		t.Errorf("does not have 1")
		t.Logf("avl: %s", a)
	}
}
