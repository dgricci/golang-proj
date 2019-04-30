package proj

/*
#cgo CFLAGS: -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include <stdlib.h>
#include <string.h>
#include "proj.h"
 */
import "C"

import (
    "unsafe"
    "fmt"
)

// Operation contains an internal object that holds everything related to a given
// coordinate transformation.
//
type Operation struct {
    pj *C.PJ
}

var (
    operations map[string]*C.PJ_OPERATIONS
)

// NewOperation creates a reference system object from a proj-string, a WKT string,
// or object code.
//
// When `bbox` is not defined then only the first element is considered :
//
//   ope, e = NewOperation(ctx, nil, "+proj=utm +zone=32 +ellps=GRS80")
//   ope, e = NewOperation(ctx, nil, "urn:ogc:def:coordinateOperation:EPSG::1671")
//
// otherwise the two first elements are considered to create a transformation object
// that is a pipeline between two known coordinate reference system
// definitions.
//
//   ope, e := NewOperation(ctx, bbox, "EPSG:25832", "EPSG:25833")
//
func NewOperation ( ctx *Context, bbox *Area, def ...string ) (op *Operation, e error) {
    var pj *C.PJ
    dA := C.CString(def[0])
    defer C.free(unsafe.Pointer(dA))
    if bbox != nil {
        if len(def) >= 2 {
            dB := C.CString(def[1])
            defer C.free(unsafe.Pointer(dB))
            pj = C.proj_create_crs_to_crs((*ctx).pj, dA, dB, (*bbox).pj)
        }
    } else {
        pj = C.proj_create((*ctx).pj, dA)
    }
    if pj == (*C.PJ)(nil) {
        e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
        return
    }
    op = &Operation{pj:pj}
    switch Type(op) {
    case Conversion,
         Transformation,
         ConcatenatedOperation,
         OtherCoordinateOperation :
        return
    default :
        op.DestroyOperation()
        e = fmt.Errorf("%v does not yield an Operation", def)
    }
    return
}

// DestroyOperation deallocates the internal Operation object.
//
func (op *Operation) DestroyOperation () {
    if op != nil {
        (*op).pj = C.proj_destroy((*op).pj)
        (*op).pj = nil
    }
}

// Handle returns the PROJ internal object to be passed to the PROJ library.
// Cannot be tested against nil as it returns a pointer to a type, so use :
//   if p.HandleIsNil() { ... }
//
func (op *Operation) Handle () (interface{}) {
    return (*op).pj
}

// HandleIsNil returns true when the PROJ internal object is NULL.
//
func (op *Operation) HandleIsNil () bool {
    return (*op).pj == (*C.PJ)(nil)
}

func (op *Operation) fwdinv ( d Direction, aC *Coordinate ) ( aR *Coordinate, e error ) {
    var cpj, cc C.PJ_COORD
    _ = C.proj_errno_reset((*op).pj)
    // make a copy not to change coord in case of error :
    _ = C.memcpy(unsafe.Pointer(&cc), unsafe.Pointer(&((*aC).pj)), C.sizeof_PJ_COORD)
    cpj = C.proj_trans((*op).pj, C.PJ_DIRECTION(d), cc)
    if En := C.proj_errno((*op).pj) ; En != C.int(0) {
        e = fmt.Errorf(C.GoString(C.proj_errno_string(En)))
    } else {
        // everything's ok, copy back :
        _ = C.memcpy(unsafe.Pointer(&((*aC).pj)), unsafe.Pointer(&cpj), C.sizeof_PJ_COORD)
        aR = aC
    }
    return
}

// Transform applies the transformation of coordinates to object
// implementing `Locatable` either from or to the CRS.
// Returns the object with transformed coordinates or nil on error.
//
func (op *Operation) Transform ( d Direction, c Locatable ) ( r Locatable, e error ) {
    xyzt := c.Location()
    xyzt, e = op.fwdinv(d, xyzt)
    if e != nil { return }
    c.SetLocation(xyzt)
    r = c
    return
}

// Factors creates various cartographic properties, such as scale factors,
// angular distortion and meridian convergence.
// Depending on the underlying projection values will be calculated either
// numerically (default) or analytically.
//
// The function also calculates the partial derivatives of the given
// geographical coordinate.
//
func (op *Operation) Factors ( c *Coordinate) ( f *Factors, e error ) {
    _ = C.proj_errno_reset((*op).pj)
    var pjf C.PJ_FACTORS
    pjf = C.proj_factors((*op).pj, (*c).pj)
    if En := C.proj_errno((*op).pj) ; En != C.int(0) {
        e = fmt.Errorf(C.GoString(C.proj_errno_string(En)))
    } else {
        f = &Factors{pj:pjf}
    }
    return
}

// Info returns information about a specific reference system object.
//
func (op *Operation) Info ( ) ( *ISOInfo ) {
    return &ISOInfo{pj:C.proj_pj_info((*op).pj)}
}

