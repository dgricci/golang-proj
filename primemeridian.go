package proj

/*
#cgo CFLAGS: -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include <stdlib.h>
#include <string.h>
#include "proj.h"

char *listcat2 ( PROJ_STRING_LIST sl ) {
    size_t l = 0;
    char *result = NULL;
    PROJ_STRING_LIST iterator = NULL;
    if (sl == NULL) return NULL ;
    for (iterator = sl; *iterator; iterator++) {
        l += strlen(*iterator);
    }
    result = (char *)malloc(l+1);
    if (result == NULL) return NULL;
    result[0] = '\0';
    for (iterator = sl; *iterator; iterator++) {
        result = strcat(result, *iterator);
    }
    return result;
}
 */
import "C"

import (
    "unsafe"
    "strings"
    "fmt"
)

// PrimeMeridian contains an internal object that holds everything related to a
// given prime meridian.
type PrimeMeridian struct {
    pj *C.PJ
}

// NewPrimeMeridian creates a prime meridian from a WKT string or URI.
//
func NewPrimeMeridian (ctx *Context, def string ) ( pm *PrimeMeridian, e error ) {
    var pj *C.PJ
    cdef := C.CString(def)
    defer C.free(unsafe.Pointer(cdef))
    switch dialect := C.proj_context_guess_wkt_dialect((*ctx).pj, cdef) ; GuessedWKTDialect(dialect) {
    case GuessedWKTUnknown  : // URI
        ac := strings.Split(def,":")
        if len(ac) != 2 {
            e = fmt.Errorf("%v does not yield a Prime Meridian", def)
            return
        }
        cauth := C.CString(ac[0])
        defer C.free(unsafe.Pointer(cauth))
        cname := C.CString(ac[1])
        defer C.free(unsafe.Pointer(cname))
        pj = C.proj_create_from_database((*ctx).pj, cauth, cname, C.PJ_CATEGORY_PRIME_MERIDIAN, 0, nil)
        if pj == (*C.PJ)(nil) {
            e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
            return
        }
    default                 : // WKT flavor wkt2_grammar.y : prime_meridian is missing from input !
        var ce C.PROJ_STRING_LIST
        pj = C.proj_create_from_wkt((*ctx).pj, cdef, nil, nil, &ce)
        if ce != (C.PROJ_STRING_LIST)(nil) {
            cm := C.listcat2(ce)
            defer C.free(unsafe.Pointer(cm))
            defer C.proj_string_list_destroy(ce)
            e = fmt.Errorf(C.GoString(cm))
            return
        }
    }
    pm = &PrimeMeridian{pj:pj}
    return
}

// DestroyPrimeMeridian deallocate the internal prime meridian object.
//
func (pm *PrimeMeridian) DestroyPrimeMeridian () {
    (*pm).pj = C.proj_destroy((*pm).pj)
    (*pm).pj = nil
}

// Handle returns the PROJ internal object to be passed to the PROJ library
//
func (pm *PrimeMeridian) Handle () (interface{}) {
    return (*pm).pj
}

// HandleIsNil returns true when the PROJ internal object is NULL.
//
func (pm *PrimeMeridian) HandleIsNil () bool {
    return (*pm).pj == (*C.PJ)(nil)
}

// Longitude returns the longitude of the prime meridian.
//
func (pm *PrimeMeridian) Longitude ( ctx *Context ) ( longitude float64, e error ) {
    _ = C.proj_errno_reset((*pm).pj)
    var cl C.double
    // proj_prime_meridian_get_parameters fails if pm is not a prime meridian ...
    _ = C.proj_prime_meridian_get_parameters((*ctx).pj, (*pm).pj, &cl, nil, nil)
    longitude = float64(cl)
    return
}

// ToRad returns the longitude of the prime meridian, in its native unit, 
// the conversion factor of the prime meridian longitude unit to radian and
// the unit name of the given prime meridian.
//
func (pm *PrimeMeridian) ToRad ( ctx *Context ) ( toRad float64, e error ) {
    _ = C.proj_errno_reset((*pm).pj)
    var cr C.double
    // proj_prime_meridian_get_parameters fails if pm is not a prime meridian ...
    _ = C.proj_prime_meridian_get_parameters((*ctx).pj, (*pm).pj, nil, &cr, nil)
    toRad = float64(cr)
    return
}

// Unit returns the longitude native unit of the given prime meridian.
//
func (pm *PrimeMeridian) Unit ( ctx *Context ) ( u string, e error ) {
    _ = C.proj_errno_reset((*pm).pj)
    var cu *C.char
    // proj_prime_meridian_get_parameters fails if pm is not a prime meridian ...
    _ = C.proj_prime_meridian_get_parameters((*ctx).pj, (*pm).pj, nil, nil, &cu)
    u = C.GoString(cu)
    return
}

// Parameters returns the longitude of the prime meridian, in its native unit, 
// the conversion factor of the prime meridian longitude unit to radian and
// the unit name of the given prime meridian.
//
func (pm *PrimeMeridian) Parameters ( ctx *Context ) ( longitude float64, toRad float64, u string, e error ) {
    _ = C.proj_errno_reset((*pm).pj)
    var cl, cr C.double
    var cu *C.char
    // proj_prime_meridian_get_parameters fails if pm is not a prime meridian ...
    _ = C.proj_prime_meridian_get_parameters((*ctx).pj, (*pm).pj, &cl, &cr, &cu)
    longitude = float64(cl)
    toRad = float64(cr)
    u = C.GoString(cu)
    return
}

// Info returns information about a specific ellipsoid object.
//
func (pm *PrimeMeridian) Info ( ) ( *ISOInfo ) {
    return &ISOInfo{pj:C.proj_pj_info((*pm).pj)}
}

