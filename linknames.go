package arena

import "unsafe"

//go:linkname arena_newArena arena.runtime_arena_newArena
func arena_newArena() unsafe.Pointer

//go:linkname arena_arena_New arena.runtime_arena_arena_New
func arena_arena_New(arena unsafe.Pointer, typ any) any

//go:linkname arena_arena_Slice arena.runtime_arena_arena_Slice
func arena_arena_Slice(arena unsafe.Pointer, slice any, cap int)

//go:linkname arena_arena_Free arena.runtime_arena_arena_Free
func arena_arena_Free(arena unsafe.Pointer)

//go:linkname arena_heapify arena.runtime_arena_heapify
func arena_heapify(s any) any
