package packaging_c_method_qsort

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_cmp(void* a, void* b);
*/
import "C"
import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var go_sort_cmpinfo struct {
	base    unsafe.Pointer
	eleSize int
	eleNum  int
	cmp     func(a, b int) bool
	sync.Mutex
}

func SortByqsort(slice interface{}, cmp func(a, b int) bool) {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic(fmt.Sprintf("qsort called with non-slice value of type %T", slice))
	}
	if value.Len() == 0 {
		return
	}
	go_sort_cmpinfo.Mutex.Lock()
	defer go_sort_cmpinfo.Mutex.Unlock()

	defer func() {
		go_sort_cmpinfo.eleNum = 0
		go_sort_cmpinfo.eleSize = 0
		go_sort_cmpinfo.base = nil
		go_sort_cmpinfo.cmp = nil
	}()

	go_sort_cmpinfo.eleSize = int(value.Type().Elem().Size())
	go_sort_cmpinfo.eleNum = value.Len()
	go_sort_cmpinfo.base = unsafe.Pointer(value.Index(0).Addr().Pointer())
	go_sort_cmpinfo.cmp = cmp

	C.qsort(go_sort_cmpinfo.base, C.size_t(go_sort_cmpinfo.eleNum), C.size_t(go_sort_cmpinfo.eleSize), C.qsort_cmp_func_t(C._cgo_qsort_cmp))
}

//export _cgo_qsort_cmp
func _cgo_qsort_cmp(a, b unsafe.Pointer) C.int {
	var (
		base    = uintptr(go_sort_cmpinfo.base)
		eleSize = uintptr(go_sort_cmpinfo.eleSize)
	)
	i := int((uintptr(a) - base) / eleSize)
	j := int((uintptr(b) - base) / eleSize)

	switch {
	case go_sort_cmpinfo.cmp(i, j):
		return -1
	case go_sort_cmpinfo.cmp(j, i):
		return 1
	default:
		return 0
	}

}
