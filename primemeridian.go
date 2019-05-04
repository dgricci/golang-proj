package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "wrapper.h"
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
        switch len(ac) {
        case 7 : // urn:ogc:def:meridian::EPSG:code
            pj = C.proj_create((*ctx).pj, cdef)
        case 2 :
            cauth := C.CString(ac[0])
            defer C.free(unsafe.Pointer(cauth))
            cname := C.CString(ac[1])
            defer C.free(unsafe.Pointer(cname))
            pj = C.proj_create_from_database((*ctx).pj, cauth, cname, C.PJ_CATEGORY_PRIME_MERIDIAN, 0, nil)
        default:
            e = fmt.Errorf("%v does not yield a Prime Meridian", def)
            return
        }
    default                 : // WKT flavor wkt2_grammar.y : prime_meridian is missing from input !
        var ce C.PROJ_STRING_LIST
        pj = C.proj_create_from_wkt((*ctx).pj, cdef, nil, nil, &ce)
        if pj == (*C.PJ)(nil) {
            if ce != (C.PROJ_STRING_LIST)(nil) {
                cm := C.listcat(ce)
                defer C.free(unsafe.Pointer(cm))
                defer C.proj_string_list_destroy(ce)
                e = fmt.Errorf(C.GoString(cm))
                //return
            }
            // not needed :
            //e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
            return
        }
    }
    if pj == (*C.PJ)(nil) {
        e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
        return
    }
    if C.proj_get_type(pj) != C.PJ_TYPE_PRIME_MERIDIAN {
        C.proj_destroy(pj)
        pj = nil
        e = fmt.Errorf("%v does not yield a Prime Meridian", def)
        return
    }
    pm = &PrimeMeridian{pj:pj}
    return
}

// DestroyPrimeMeridian deallocate the internal prime meridian object.
//
func (pm *PrimeMeridian) DestroyPrimeMeridian () {
    if (*pm).pj != nil {
        C.proj_destroy((*pm).pj)
        (*pm).pj = nil
    }
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

// TypeOf returns the ISOType of a prime meridian (PrimeMeridianType).
// UnKnownType on error.
//
func (pm *PrimeMeridian) TypeOf ( ) ISOType {
    return hasType(pm)
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

// Info returns information about a specific prime meridien object.
//
func (pm *PrimeMeridian) Info ( ) ( *ISOInfo ) {
    return &ISOInfo{pj:C.proj_pj_info((*pm).pj)}
}

// String returns a string representation of the prime meridien.
//
func (pm *PrimeMeridian) String ( ) string {
    return toString(pm)
}

// ProjString returns a proj-string representation of the prime meridian.
// Empty string is returned on error (sounds to be the case : no conversion).
//
func (pm *PrimeMeridian) ProjString ( ctx *Context, styp StringType, opts ...string ) string {
    return toProj(ctx, pm, styp, nil)
}

// Wkt return returns a WKT representation of the prime meridian.
// Empty string is returned on error.
// `opts` can be hold the following strings :
//
//   "MULTILINE=YES" Defaults to YES, except for styp equals WKT1_ESRI
//
//   "INDENTATION_WIDTH=<number>" Defaults to 4 (when multiline output is on)
//
//   "OUTPUT_AXIS=AUTO/YES/NO" In AUTO mode, axis will be output for WKT2
//   variants, for WKT1_GDAL for ProjectedCRS with easting/northing ordering
//   (otherwise stripped), but not for WKT1_ESRI. Setting to YES will output
//   them unconditionally, and to NO will omit them unconditionally.
//
func (pm *PrimeMeridian) Wkt ( ctx *Context, styp WKTType, opts ...string ) string {
    return toWkt(ctx, pm, styp, opts)
}

