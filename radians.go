package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"

import (
    "unsafe"
)

// Conversion between radians, degrees and grades
//
const (
	DegToRad = 0.017453292519943295769236907684886127134428718885417    // Pi/180
	RadToDeg = 57.295779513082320876798154814105170332405472466564      // 180/Pi

	GradToRad = 0.015707963267948966192313216916397514420985846996876   // Pi/200
	RadToGrad = 63.661977236758134307553505349005744813783858296183     // 200/Pi
)

// Conversion between radians and sexagecimal degrees

// DMSToRad converts a sexagecimal degrees string into radians.
//
func DMSToRad ( dms string ) (r float64) {
    cdms := C.CString(dms)
    defer C.free(unsafe.Pointer(cdms))
    return (float64)(C.wrapper_proj_dmstor(cdms))
}

// RadToDMS converts radians into a sexagecimal degrees string.
// `posNE` is usually 'N' (north) or 'E' (east) for positive values and
// `negSW` is usually 'S' (south) or 'W' (west) for negative values.
//
func RadToDMS ( r float64, posNE byte, negSW byte ) (dms string) {
    buf := string(make([]byte,128))
    cdms := C.CString(buf)
    defer C.free(unsafe.Pointer(cdms))
    cdms = C.proj_rtodms(cdms,C.double(r),C.int(posNE),C.int(negSW))
    return C.GoString(cdms)
}

