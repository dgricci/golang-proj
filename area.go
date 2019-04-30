package proj

/*
#cgo CFLAGS: -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "proj.h"
 */
import "C"

// Area contains a opaque object describing an area in which a transformation is performed.
//
type Area struct {
    pj *C.PJ_AREA
}

/*
var (
    World := NewArea(-180, -90, 180, 90)
)
 */

// NewArea creates an area of use.
// `lonmin` West longitude, in degrees. In [-180,180] range.
// `latmin` South latitude, in degrees. In [-90,90] range.
// `lonmax` East longitude, in degrees. In [-180,180] range.
// `latmax` North latitude, in degrees. In [-90,90] range.
// In the case of an area of use crossing the antimeridian (longitude +/- 180
// degrees), `lonmin` will be greater than `lonmax`.
//
func NewArea ( lonmin float64, latmin float64, lonmax float64, latmax float64) (*Area) {
    a := &Area{pj:C.proj_area_create()}
    C.proj_area_set_bbox(a.pj, C.double(lonmin), C.double(latmin), C.double(lonmax), C.double(latmax))
    return a
}

// DestroyArea deallocates the internal PROJ area pointer
//
func (a *Area) DestroyArea () {
    if a != nil {
        C.proj_area_destroy((*a).pj)
        (*a).pj = nil
    }
}

// Handle returns the PROJ internal object to be passed to the PROJ library
//
func (a *Area) Handle () (interface{}) {
    return (*a).pj
}

// HandleIsNil returns true when the PROJ internal object is NULL.
//
func (a *Area) HandleIsNil () bool {
    return (*a).pj == (*C.PJ_AREA)(nil)
}

/*
// SetBBox sets the bounding box of the area of use.
// In the case of an area of use crossing the antimeridian (longitude +/- 180
// degrees), `lonmin` will be greater than `lonmax`.
//
func (a *Area) SetBBox ( lonmin float64, latmin float64, lonmax float64, latmax float64 ) {
    C.proj_area_set_bbox(a.pj, C.double(lonmin), C.double(latmin), C.double(lonmax), C.double(latmax))
}
 */

/* TODO how to destroy opaque object in the end ?
// init package initialisation
//
func init () {
    runtime.SetFinalizer(World, func (a *Area) {
        a.DestroyArea()
    }
}
 */
