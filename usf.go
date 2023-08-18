package usf

/*
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
*/
import "C"
import (
	"unsafe"
)

type (
	_any struct {
		size *uint64
		addr unsafe.Pointer
	}
	_slice struct {
		addr unsafe.Pointer
		len  uint64
		cap  uint64
	}
	_ptr struct{ p unsafe.Pointer }
)

func AddrOf(i interface{}) unsafe.Pointer {
	return (*_any)(unsafe.Pointer(&i)).addr
}
func SizeOf(i interface{}) uint64 {
	return *(*_any)(unsafe.Pointer(&i)).size
}
func Malloc(size uint64) unsafe.Pointer {
	return C.malloc(C.uint64_t(size))
}
func MallocN(n uint64, typeSize uint64) unsafe.Pointer {
	return C.malloc(C.uint64_t(n * typeSize))
}
func MallocOf(size uint64, zeroVal interface{}) unsafe.Pointer {
	return C.malloc(C.uint64_t(size * SizeOf(zeroVal)))
}
func Free(p unsafe.Pointer) {
	if p == nil {
		return
	}
	C.free(p)
}
func Memset(p unsafe.Pointer, bytev int32, bytec uint64) {
	C.memset(p, C.int(bytev), C.uint64_t(bytec))
}
func Memcpy(dest, src unsafe.Pointer, n uint64) { C.memcpy(dest, src, C.uint64_t(n)) }
func Slice(p unsafe.Pointer, n uint64) unsafe.Pointer {
	return unsafe.Pointer(&_slice{addr: p, len: n, cap: n})
}

func Push(dst unsafe.Pointer, src unsafe.Pointer)   { (*_ptr)(dst).p = src }
func Pop(ptr unsafe.Pointer) (inPtr unsafe.Pointer) { return (*_ptr)(ptr).p }
func PushAt(list unsafe.Pointer, idx uint64, ptr unsafe.Pointer) {
	(*_ptr)(unsafe.Add(list, idx*8)).p = ptr
}
func PopAt(list unsafe.Pointer, idx uint64) (ptrAddr unsafe.Pointer) {
	return (*_ptr)(unsafe.Add(list, idx*8)).p
}

// func Ptrs(clist unsafe.Pointer, num uint64) []unsafe.Pointer {
// 	return *(*[]unsafe.Pointer)(Slice(clist, num))
// }
// func Clist(src []unsafe.Pointer) unsafe.Pointer {
// 	if src == nil {
// 		return nil
// 	}
// 	dst := Malloc(uint64(len(src)), 8)
// 	p := *(*[]unsafe.Pointer)(Slice(dst, uint64(len(src))))
// 	copy(p, src)
// 	return dst
// }
