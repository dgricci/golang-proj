package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"

import (
    "unsafe"
    "fmt"
)

// ReferenceSystem contains an internal object that holds everything related to a given
// reference system and derivatives.
//
type ReferenceSystem struct {
    pj *C.PJ
}

// Direction applies transformation to observation - in forward or inverse direction
//
type Direction C.PJ_DIRECTION
const (
    // Forward is forward transformation
    Forward    Direction =  C.PJ_FWD
    // Identity does nothing
    Identity   Direction =  C.PJ_IDENT
    // Inverse is inverse transformation
    Inverse    Direction = C.PJ_INV
)

// NewReferenceSystem creates a reference system object from a proj-string, a WKT string,
// or object code.
//
//   crs := NewReferenceSystem(ctx, "+proj=utm +zone=32 +ellps=GRS80 +type=crs")
//
//   crs := NewReferenceSystem(ctx, "EPSG:4326")
//
//   crs := NewReferenceSystem(ctx, "urn:ogc:def:crs:EPSG::4326")
//
// in the defining proj-string is an entry in `def` :
//
//   crs := NewReferenceSystem(ctx, "proj=utm", "zone=32", "ellps=GRS80", "type=crs")
//
func NewReferenceSystem ( ctx *Context, def ...string ) (crs *ReferenceSystem, e error) {
    var pj *C.PJ
    l := len(def)
    if l == 0 {
        e = fmt.Errorf(C.GoString(C.proj_errno_string(-1)))
        return
    }
    if l == 1 {
        d := C.CString(def[0])
        pj = C.proj_create((*ctx).pj, d)
        C.free(unsafe.Pointer(d))
    } else {
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
    if pj == (*C.PJ)(nil) {
        e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
        return
    }
    if C.proj_is_crs(pj) == C.int(0) {
        C.proj_destroy(pj)
        e = fmt.Errorf("%v does not yield a CRS", def)
        return
    }
    crs = &ReferenceSystem{pj:pj}
    return
}

// DestroyReferenceSystem deallocates the internal ReferenceSystem object.
//
func (crs *ReferenceSystem) DestroyReferenceSystem () {
    (*crs).pj = C.proj_destroy((*crs).pj)
    (*crs).pj = nil
}

// Handle returns the PROJ internal object to be passed to the PROJ library
// Cannot be tested against nil as it returns a pointer to a type, so use :
//   if p.HandleIsNil() { ... }
//
func (crs *ReferenceSystem) Handle () (interface{}) {
    return (*crs).pj
}

// HandleIsNil returns true when the PROJ internal object is NULL.
//
func (crs *ReferenceSystem) HandleIsNil () bool {
    return (*crs).pj == (*C.PJ)(nil)
}

// OperationFilter allows setting different parameter when searching
// operations in `NewOperation`
//
type OperationFilter struct {
    // Authority can be "", "any" or any non-empty string different of "any" :
    //   ""    : coordinate operations from any authority will be searched, with the
    //           restrictions set in the authority_to_authority_preference database table.
    //   "any" : coordinate operations from any authority will be searched.
    //   "..." : non-empty string different of "any", then coordinate operations
    //           will be searched only in that authority namespace.
    Authority   string
    // in meter, 0 to disable this constraint
    Accuracy    float64
    // desired area of interest for the resulting coordinate transformations.
    // the slice holds the west longitude (in degrees), the south latitude (in
    // degrees), the east longitude (in degrees) and the north latitude (in
    // degrees).
    BBox        []float64
    // how source and target CRS extent should be used when considering if a
    // transformation can be used (only takes effect if no area of interest is
    // explicitly defined). Default is `SmallestExtent`
    XUse        CRSExtentUse
    // spatial criterion to use when comparing the area of validity of
    // coordinate operations with the area of interest / area of validity of
    // source and target CRS. Default is `StrictContainment`
    SCriterion  SpatialCriterion
    // how grid availability is used. Default is `SortGrids`
    GUse        GridAvailabilityUse
    //  whether PROJ alternative grid names should be substituted to the
    //  official authority names. Default is `true`
    AltGrid     bool
    // whether an intermediate pivot CRS can be used for researching
    // coordinate operations between a source and target CRS. Default is
    // `AlwaysUse`
    PivotUse    IntermediateCRSUse
    // restrict the potential pivot CRSs that can be used when trying to build
    // a coordinate operation between two CRS that have no direct operation.
    // Default is no restriction
    Pivots      map[string][]string
}

var (
    // DefaultFilter for `NewOperation`, it does no filter at all, but can use
    // as base to build a filter.
    DefaultFilter = OperationFilter{
        Authority : "",
        Accuracy  : 0.0,
        BBox      : nil,
        XUse      : SmallestExtent,
        SCriterion: StrictContainment,
        GUse      : SortGrids,
        AltGrid   : true,
        PivotUse  : AlwaysUse,
        Pivots    : nil,
    }
)

// NewOperation creates a transformation from the reference system to the
// given reference system. An area may be added to the creation to restrict
// the bounding box of the transformation.
//
func (crs *ReferenceSystem) NewOperation ( ctx *Context, targetCrs *ReferenceSystem, filter ...OperationFilter ) ( op *Operation, e error) {
    _ = C.proj_errno_reset((*crs).pj)
    var opFilter OperationFilter
    if len(filter) == 0 {
        opFilter = DefaultFilter
    } else {
        opFilter = filter[0]
    }
    cauth := C.CString(opFilter.Authority)
    defer C.free(unsafe.Pointer(cauth))
    opeFactory := C.proj_create_operation_factory_context((*ctx).pj, cauth)
    if opeFactory == (*C.PJ_OPERATION_FACTORY_CONTEXT)(nil) {// no more memory ??
        e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
        return
    }
    defer C.proj_operation_factory_context_destroy(opeFactory)
    if opFilter.Accuracy != 0.0 {
        C.proj_operation_factory_context_set_desired_accuracy(
            (*ctx).pj, opeFactory, C.double(opFilter.Accuracy),
        )
    }
    if opFilter.BBox != nil {
        C.proj_operation_factory_context_set_area_of_interest(
            (*ctx).pj, opeFactory, C.double(opFilter.BBox[0]), C.double(opFilter.BBox[1]), C.double(opFilter.BBox[2]), C.double(opFilter.BBox[3]),
        )
    }
    if opFilter.XUse != SmallestExtent {
        C.proj_operation_factory_context_set_crs_extent_use((*ctx).pj, opeFactory, C.PROJ_CRS_EXTENT_USE(opFilter.XUse))
    }
    if opFilter.SCriterion != StrictContainment {
        C.proj_operation_factory_context_set_spatial_criterion((*ctx).pj, opeFactory, C.PROJ_SPATIAL_CRITERION(opFilter.SCriterion))
    }
    if opFilter.GUse != SortGrids {
        C.proj_operation_factory_context_set_grid_availability_use((*ctx).pj, opeFactory, C.PROJ_GRID_AVAILABILITY_USE(opFilter.GUse))
    }
    if !opFilter.AltGrid {
        C.proj_operation_factory_context_set_use_proj_alternative_grid_names((*ctx).pj, opeFactory, C.int(0))
    }
    if opFilter.PivotUse != AlwaysUse {
        C.proj_operation_factory_context_set_allow_use_intermediate_crs((*ctx).pj, opeFactory, C.PROJ_INTERMEDIATE_CRS_USE(opFilter.PivotUse))
    }
    if opFilter.Pivots != nil {
        // PROJ expects an array of strings NULL terminated, with the format { “auth_name1”, “code1”, “auth_name2”, “code2”, … NULL }
        // count total number of elements
        n := 0
        for _, codes := range opFilter.Pivots {
            n += 2*len(codes)
        }
        n++ // last element is NULL
        cpivots := C.makeStringArray(C.size_t(n))
        n = 0
        for auth, codes := range opFilter.Pivots {
            for _, code := range codes {
                cauth := C.CString(auth)
                C.setStringArrayItem(cpivots, C.size_t(n), cauth)
                n++
                ccode := C.CString(code)
                C.setStringArrayItem(cpivots, C.size_t(n), ccode)
                n++
            }
        }
        C.setStringArrayItem(cpivots, C.size_t(n), nil)
        n++
        C.proj_operation_factory_context_set_allowed_intermediate_crs((*ctx).pj, opeFactory, cpivots)
        for i := 0 ; i < n ; i++ {
            C.free(unsafe.Pointer(C.getStringArrayItem(cpivots,C.size_t(i))))
        }
        C.destroyStringArray(&cpivots)
    }
    candidateCrs := C.proj_create_operations((*ctx).pj, (*crs).pj, (*targetCrs).pj, opeFactory)
    if candidateCrs == (*C.PJ_OBJ_LIST)(nil) {// one of the crs is not a CRS, no more memory
        e = fmt.Errorf(C.GoString(C.proj_errno_string(C.proj_context_errno((*ctx).pj))))
        return
    }
    defer C.proj_list_destroy(candidateCrs)
    if C.proj_list_get_count(candidateCrs) == 0 {
        crsS := C.GoString(C.proj_get_name((*crs).pj))
        crsT := C.GoString(C.proj_get_name((*targetCrs).pj))
        e = fmt.Errorf("No operation found between '%s' and '%s'", crsS, crsT)
        return
    }
    // return the first operation as the operations are sorted with the most
    // relevant ones first: by descending area (intersection of the
    // transformation area with the area of interest, or intersection of the
    // transformation with the area of use of the CRS), and by increasing
    // accuracy. Operations with unknown accuracy are sorted last, whatever
    // their area.
    // counting is done for 0 (not documented, but code says : result->objects[index]
    op = &Operation{pj:C.proj_list_get((*ctx).pj, candidateCrs, C.int(0))}
    return
}

// Info returns information about a specific reference system object.
//
func (crs *ReferenceSystem) Info ( ) ( *ISOInfo ) {
    return &ISOInfo{pj:C.proj_pj_info((*crs).pj)}
}

