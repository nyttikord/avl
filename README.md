# AVL

Library implementing AVL trees in Go.
                                                                                                              
## Usage

Install the library with
```bash
go get -u github.com/nyttikord/avl@latest
```

You can create use simple AVL tree storing int with:
```go
// create a new AVL storing int using the standard cmp.Compare.
tree := avl.NewSimple[int]()

// insert new nodes
tree.Insert(1, 2, 3, 8, 7, 5)

// delete nodes
tree.Delete(5, 2)

// get a node
var got *int
got = tree.Get(func (v int) { return 1 - v }) // returns the node containing 1
if got == nil {
    println("not found!")
} else {
    println("got", *got)
}

// get the maximum
got = tree.Max()

// get the minimum
got = tree.Min()

// get number of nodes
n := tree.Size()

// get sorted nodes
sorted := tree.Sort()

// clone the tree
cloned := tree.Clone()
```

### Complex types 

If you want to store more complex types, you can:
```go
// `comp` is a function comparing two values.
// It must returns 0 if a == b, > 0 if a > b, and < 0 if a < b.
// The order must be linear (or total, or simple) on the set used.
//
// comp = func(a, b int) { return a - b } is a good function for ordering the set of integer.
tree := avl.New(comp)

// `find` is a function returning 0 for the requested value.
// It must returns < 0 if the value checked if before the requested value and > 0 if it is after.
// It is useful when you are sorting a set with a key.
//
// find = func(v int) { return v - 5 } is a good function for getting the 5.
got := avl.Get(find)
```

### Immutable data

By default, the tree returned does not contain immutable data.
It has side effects if it contains pointers (like an array).
To avoid side effects, you can create an AVL storing immutable data.
```go
// `clone` is a function cloning the given value.
// It must perform a deep copy to completely avoid side effects.
//
// clone = func(v []int) []int {
//   dst := make([]int, len(v))
//   copy(dst, v)
//   return dst
// } is a good function for cloning []int.
tree := avl.NewImmutable(comp, clone)
```

If the type stored in the AVL implements `avl.Clonable[T]`, `avl.New` automatically creates an immutable data using the
method defined by this implementation.
```go
package main

type intArr []int

// Clone implements avl.Clonable[intArr] for intArr to creates immutable data.
func (i intArr) Clone() intArr {
    dst := make(intArr, len(v))
    copy(dst, v)
    return dst
}

func main() {
    // automatically creates an AVL storing immutable data, because intArr implements avl.Clonable[intArr]
    tree := avl.New(func(a, b intArr) { return len(a) - len(b) }) // this is a bad comp function, because {0} == {1}!
}
```

If you want to store a type implementing `avl.Clonable[T]` in a mutable AVL, use `avl.NewMutable`.
```go
// http.Header already implements avl.Clonable[http.Header]
var header http.Header
// creates an AVL storing mutable data, despite its implementation of avl.Clonable[http.Header]
tree := avl.NewMutable(header)
```

### Helping functions

You can use helping functions to avoid writing common comparison function.
```go
// create a new AVL using the standard cmp.Compare function.
tree := avl.NewSimple[int]()

// Create a new AVL storing strings using the standard strings.Compare function.
// This function is faster than avl.NewOrdered (because strings.Compare is faster than cmp.Compare).
tree = avl.NewString()
```

`avl.NewSimpleImmutable` and `avl.NewSimpleMutable` are available.

These functions return a `*SimpleAVL[T]` that has the method `Has(T) bool` checking if the value is present in the AVL.
```go
tree := avl.NewSimple[int]()
tree.Insert(1, 2, 3)
if !tree.Has(1) {
    panic("tree does not have 1")
}
```

### Key-value storage

You can use AVL trees as maps easily with the type `KeyAVL`.
```go
// Create a new key-value AVL.
// The compare function only targets keys.
tree := avl.NewKey[int, string](cmp.Compare)
// you can also write
tree = avl.NewKeySimple[int, string]()

// insert "hello" with the key 1
tree.Insert(1, "hello").
    Insert(2, "world").
    Insert(3, "world")

// print "hello world !"
st := tree.Sort()
println(strings.Join(st, " "))

// remove the value associated with the key 2
tree.Delete(2)

// print "hello !"
st = tree.Sort()
println(strings.Join(st, " "))

// check if a value is associated with key 2
if !tree.Has(2) {
    println("no world :(")
    tree.Insert(2, "new world")
}

// get the value associated with the key 2
got := tree.Get(2)
if got == nil {
    println("value not found")
} else {
    println(*got)
}
```

Functions `NewKeyImmutable`, `NewKeyMutable`, `NewKeySimpleImmutable` and `NewKeySimpleMutable` are available too.
