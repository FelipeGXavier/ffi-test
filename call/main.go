package main

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>

typedef void (*AddNumberFunction)(int, int);

int add_number(void* f, int x, int y) {
    ((AddNumberFunction) f)(x, y);
}

*/
import "C"
import "fmt"

func ffi() {
	handle := C.dlopen(C.CString("../lib/libhello.so"), C.RTLD_LAZY)
	add_number_ptr := C.dlsym(handle, C.CString("add_number"))
	result := C.add_number(add_number_ptr, C.int(99), C.int(10))
	fmt.Println(handle, result)
}

func main() {
	ffi()
}
