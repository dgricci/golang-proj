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
//   ope, e = NewOperation(ctx, nil, "EPSG:9616")
//
//   ope, e = NewOperation(ctx, nil, "+proj=utm +zone=32 +ellps=GRS80")
//
//   ope, e = NewOperation(ctx, nil, "urn:ogc:def:coordinateOperation:EPSG::1671")
//
//   ope, e = NewOperation(ctx, nil, "WKT string")
//
// with the exception of :
//
//   ope, e = NewOperation(ctx, nil, "proj=utm", "zone=32", "ellps=GRS80")
//
// otherwise the two first elements are considered to create a transformation object
// that is a pipeline between two known coordinate reference system
// definitions.
//
//   ope, e := NewOperation(ctx, bbox, "EPSG:25832", "EPSG:25833")
//
func NewOperation ( ctx *Context, bbox *Area, def ...string ) (op *Operation, e error) {
    var pj *C.PJ
    l := len(def)
    if l == 0 {
        e = fmt.Errorf(C.GoString(C.proj_errno_string(-1)))
        return
    }
    dialect := GuessedWKTUnknown
    cdef := C.CString(def[0])
    defer C.free(unsafe.Pointer(cdef))
    if l == 1 {
        dialect = GuessedWKTDialect(C.proj_context_guess_wkt_dialect((*ctx).pj, cdef))
    }
    switch dialect {
    case GuessedWKTUnknown  :
        switch {
        case l==1 :
            ac := strings.Split(def[0],":")
            if len(ac) == 2 {// <AUTH>:<CODE>
                cauth := C.CString(ac[0])
                defer C.free(unsafe.Pointer(cauth))
                cname := C.CString(ac[1])
                defer C.free(unsafe.Pointer(cname))
                pj = C.proj_create_from_database((*ctx).pj, cauth, cname, C.PJ_CATEGORY_COORDINATE_OPERATION, 0, nil)
            } else {
                d := C.CString(def[0])
                defer C.free(unsafe.Pointer(d))
                pj = C.proj_create((*ctx).pj, d)
            }
        case l==2 && bbox != nil :// src and tgt CRSs
            // proj_create_crs_to_crs() is a high level function over
            // proj_create_operations() : it can then returns several
            // operations (Cf. projinfo -s  -o PROJ -s IGNF:NTFLAMB2E.NGF84 -t IGNF:ETRS89LCC.EVRF2000
            src, se := NewReferenceSystem(ctx, def[0])
            if se != nil { e = se ; return }
            defer src.DestroyReferenceSystem()
            tgt, te := NewReferenceSystem(ctx, def[1])
            if te != nil { e = te ; return }
            defer tgt.DestroyReferenceSystem()
            opeFactory := C.proj_create_operation_factory_context((*ctx).pj, nil)
            if opeFactory == (*C.PJ_OPERATION_FACTORY_CONTEXT)(nil) {
                e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
                return
            }
            defer C.proj_operation_factory_context_destroy(opeFactory)
            candidateOps := C.proj_create_operations((*ctx).pj, (*src).pj, (*tgt).pj, opeFactory)
            if candidateOps == (*C.PJ_OBJ_LIST)(nil) {
                e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
                return
            }
            defer C.proj_list_destroy(candidateOps)
            if C.proj_list_get_count(candidateOps) == 0 {
                e = fmt.Errorf("No operation found between '%s' and '%s'", def[0], def[1])
                return
            }
            pj = C.proj_list_get((*ctx).pj, candidateOps, C.int(0))
        default :
            defs := C.makeStringArray(C.size_t(l))
            for i, partdef := range def {
                partd := C.CString(partdef)
                C.setStringArrayItem(defs, C.size_t(i), partd)
            }
            pj = C.proj_create_argv((*ctx).pj, C.int(l), defs)
            for i := 0 ; i < l ; i++ {
                C.free(unsafe.Pointer(C.getStringArrayItem(defs,C.size_t(i))))
            }
            C.destroyStringArray(&defs)
        }
    default:// WKT
        var ce C.PROJ_STRING_LIST
        pj = C.proj_create_from_wkt((*ctx).pj, cdef, nil, nil, &ce)
        if ce != (C.PROJ_STRING_LIST)(nil) {// FIXME : PROJ 6.1.0 should return an error with proj_context_errno
            cm := C.listcat(ce)
            defer C.free(unsafe.Pointer(cm))
            defer C.proj_string_list_destroy(ce)
            e = fmt.Errorf(C.GoString(cm))
            return
        }
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
        e = fmt.Errorf("%v does not yield an Operation (%v)", def, Type(op))
        op.DestroyOperation()
        op = nil
    }
    return
}

// DestroyOperation deallocates the internal Operation object.
//
func (op *Operation) DestroyOperation () {
    if (*op).pj != nil {
        C.proj_destroy((*op).pj)
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

