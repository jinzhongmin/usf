package usf

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"unsafe"
)

// 得到指针存储的值
func p2v(p unsafe.Pointer) uint64 {
	return *(*uint64)((unsafe.Pointer)(&p))
}

// 将一个值放入指针
func v2p(v uint64, p *unsafe.Pointer) {
	*(*uint64)((unsafe.Pointer)(p)) = v
}

func anyPtr(i interface{}) (len uint64, ptr unsafe.Pointer) {
	lp := (*(*[2]unsafe.Pointer)(unsafe.Pointer(&i)))
	return *(*uint64)(lp[0]), lp[1]
}

type slice struct {
	addr unsafe.Pointer
	len  uint64
	cap  uint64
}

func Sizeof(i interface{}) uint64 {
	l, _ := anyPtr(i)
	return l
}
func Malloc(size uint64, typeSize uint64) unsafe.Pointer {
	return C.malloc(C.ulonglong(size * typeSize))
}
func MallocOf(size uint64, zeroVal interface{}) unsafe.Pointer {
	return C.malloc(C.ulonglong(size * Sizeof(zeroVal)))
}
func Free(p interface{}) {
	_, a := anyPtr(p)
	C.free(a)
}
func Memset(p unsafe.Pointer, bytev int32, bytec uint64) {
	v := (C.int)(int32(bytev))
	c := (C.ulonglong)(bytec)
	C.memset(p, v, c)
}
func Slice(p unsafe.Pointer, n uint64) unsafe.Pointer {
	s := &slice{addr: p, len: n, cap: n}
	return unsafe.Pointer(s)
}

func Push(dst unsafe.Pointer, src unsafe.Pointer) {
	(*(*[1]uint64)(dst))[0] = p2v(src)
}
func PushAt(list unsafe.Pointer, idx uint64, ptr unsafe.Pointer) {
	p := unsafe.Add(list, idx*8)
	(*(*[1]uint64)(p))[0] = p2v(ptr)
}
func Pop(ptr unsafe.Pointer) (inPtr unsafe.Pointer) {
	return (*(*[1]unsafe.Pointer)(ptr))[0]
}
func PopAt(list unsafe.Pointer, idx uint64) (ptrAddr unsafe.Pointer) {
	p := unsafe.Add(list, idx*8)
	return (*(*[1]unsafe.Pointer)(p))[0]
}

func Ptrs(clist unsafe.Pointer, num uint64) []unsafe.Pointer {
	return *(*[]unsafe.Pointer)(Slice(clist, num))
}
func Clist(src []unsafe.Pointer) unsafe.Pointer {
	if src == nil {
		panic("src could not be nil")
	}
	dst := Malloc(uint64(len(src)), 8)
	p := *(*[]unsafe.Pointer)(Slice(dst, uint64(len(src))))
	copy(p, src)
	return dst
}
