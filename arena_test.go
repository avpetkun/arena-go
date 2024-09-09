// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arena_test

import (
	"testing"

	"github.com/avpetkun/arena-go"
)

type T1 struct {
	n int
}
type T2 [1 << 20]byte // 1MiB

func TestSmoke(t *testing.T) {
	a := arena.NewArena()
	defer a.Free()

	tt := arena.New[T1](a)
	tt.n = 1

	ts := arena.MakeSlice[T1](a, 99, 100)
	if len(ts) != 99 {
		t.Errorf("Slice() len = %d, want 99", len(ts))
	}
	if cap(ts) != 100 {
		t.Errorf("Slice() cap = %d, want 100", cap(ts))
	}
	ts[1].n = 42
}

func TestSmokeLarge(t *testing.T) {
	a := arena.NewArena()

	defer a.Free()
	for i := 0; i < 10*64; i++ {
		_ = arena.New[T2](a)
	}
}

// goos: darwin
// goarch: arm64
// pkg: github.com/avpetkun/arena-go
// cpu: Apple M1 Pro
// BenchmarkArena/new			11620 ns/op     16000 B/op		1000 allocs/op
// BenchmarkArena/arena-8        7250 ns/op		16084 B/op         3 allocs/op

func BenchmarkArena(b *testing.B) {
	type Object struct {
		A int
		B int
	}

	b.Run("new", func(b *testing.B) {
		refs := make([]*Object, 1000)
		b.ResetTimer()
		for range b.N {
			for i := range 1000 {
				refs[i] = new(Object)
			}
		}
	})
	b.Run("arena", func(b *testing.B) {
		refs := make([]*Object, 1000)
		newObject := arena.NewFactory[Object]()
		b.ResetTimer()
		for range b.N {
			a := arena.NewArena()
			for i := range 1000 {
				refs[i] = newObject(a)
			}
			a.Free()
		}
	})
}
