package goparsetime

/*
#include <stdlib.h>
#include <time.h>
#include "parsetime.h"
*/
import "C"
import (
	"errors"
	"time"
	"unsafe"
)

func Parsetime(t string) (time.Time, error) {
	var (
		argc C.int
		argv *C.char
		res  C.time_t
	)

	argc = 1
	argv = C.CString(t)
	defer func() {
		C.free(unsafe.Pointer(argv))
	}()

	if ret := C.parsetime(argc, &argv, &res); ret == 0 {
		return time.Unix(int64(res), 0), nil
	} else {
		err := errors.New(C.GoString(C.ErrorMessages[ret]))
		return time.Time{}, err
	}
}
