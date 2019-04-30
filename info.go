package proj

/*
#cgo CFLAGS: -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include <stdlib.h>
#include "proj.h"
 */
import "C"
import "unsafe"

// Info holds global information about PROJ library
//
type Info struct {
    pj C.PJ_INFO
}

// ISOInfo holds information about a `ReferenceSystem`, an `Operation`, an
// `Ellipsoid`, a `PrimeMeridian`, a `CoordinateSystem`, a `Datum`.
//
type ISOInfo struct {
    pj C.PJ_PROJ_INFO
}

// GridInfo holds information about a specific grid in the search path of PROJ.
//
type GridInfo struct {
    pj C.PJ_GRID_INFO
}

// InitInfo holds information about a specific init file in the search path of PROJ.
//
type InitInfo struct {
    pj C.PJ_INIT_INFO
}

// Release returns PROJ library release information such as 'Rel. 6.0.0, March 1st, 2019'.
//
func Release () string {
    return C.GoString((*projInfo).pj.release)
}

// VersionNumber returns text representation of the full version number, e.g. '6.0.0'.
//
func VersionNumber () string {
    return C.GoString((*projInfo).pj.version)
}

// Major returns the major version number, e.g. '6'.
//
func Major () uint {
    return uint((*projInfo).pj.major)
}

// Minor returns the minor version number, e.g. '0'.
//
func Minor () uint {
    return uint((*projInfo).pj.minor)
}

// Patch returns the patch level of release, e.g. '0'.
//
func Patch () uint {
    return uint((*projInfo).pj.patch)
}

// ID returns the short id of the `ReferenceSystem` or
// `Operation` type is based on, that is, what comes afther the +proj= in a
// proj-string, e.g. 'merc'.
//
func (i *ISOInfo) ID () string {
    return C.GoString((*i).pj.id)
}

// Description describes the operation the `ReferenceSystem` or `Operation`
// type is based on, e.g. 'Mercator Cyl, Sph&Ell lat_ts='.
//
func (i *ISOInfo) Description () string {
    return C.GoString((*i).pj.description)
}

// Definition returns the proj-string that was used to create the
// `ReferenceSystem` or `Operation` type
// with, e.g. '+proj=merc +lat_0=24 +lon_0=53 +ellps=WGS84'.
//
func (i *ISOInfo) Definition () string {
    return C.GoString((*i).pj.definition)
}

// HasInverse returns true if an inverse mapping of the defined operation
// exists, otherwise false.
//
func (i *ISOInfo) HasInverse () bool {
    if (*i).pj.has_inverse == C.int(1) {
        return true
    }
    return false
}

// Accuracy returns the expected accuracy of the transformation. -1 if
// unknown.
//
func (i *ISOInfo) Accuracy () float64 {
    return float64((*i).pj.accuracy)
}

// NewGridInfo creates information about a specific grid.
//
func NewGridInfo ( gridName string ) (*GridInfo) {
    gn := C.CString(gridName)
    defer C.free(unsafe.Pointer(gn))
    return &GridInfo{pj:C.proj_grid_info(gn)}
}

// GridName returns the name of grid, e.g. 'BETA2007.gsb'.
//
func (i *GridInfo) GridName () string {
    return C.GoString(&((*i).pj.gridname)[0])
}

// GridPath returns the full path of grid file, e.g.
// 'C:\OSGeo4W64\share\proj\BETA2007.gsb'.
//
func (i *GridInfo) GridPath () string {
    return C.GoString(&((*i).pj.filename)[0])
}

// Format returns the file format of grid file, e.g. 'ntv2'.
//
func (i *GridInfo) Format () string {
    return C.GoString(&((*i).pj.format)[0])
}

// LowerLeft returns the geodetic coordinate (λ,φ) of lower left corner of grid.
// Longitude and Latitude are expressed in radians.
//
func (i *GridInfo) LowerLeft () (*Coordinate) {
    return NewCoordinate(float64((*i).pj.lowerleft.lam), float64((*i).pj.lowerleft.phi))
}

// UpperRight returns the geodetic coordinate (λ,φ) of upper right corner of grid.
// Longitude and Latitude are expressed in radians.
//
func (i *GridInfo) UpperRight () (*Coordinate) {
    return NewCoordinate(float64((*i).pj.upperright.lam), float64((*i).pj.upperright.phi))
}

// LongitudinalLen returns the number of grid cells in the longitudinal direction.
//
func (i *GridInfo) LongitudinalLen () uint {
    return uint((*i).pj.n_lon)
}

// LatitudinalLen returns the number of grid cells in the latitudianl direction.
//
func (i *GridInfo) LatitudinalLen () uint {
    return uint((*i).pj.n_lat)
}

// LongitudinalCellSize returns the cell size in the longitudinal direction, in radians.
//
func (i *GridInfo) LongitudinalCellSize () float64 {
    return float64((*i).pj.cs_lon)
}

// LatitudinalCellSize returns the cell size in the latitudinal direction, in radians.
//
func (i *GridInfo) LatitudinalCellSize () float64 {
    return float64((*i).pj.cs_lat)
}

// NewInitInfo creates information about a specific init file.
//
func NewInitInfo ( initName string ) (*InitInfo) {
    in := C.CString(initName)
    defer C.free(unsafe.Pointer(in))
    return &InitInfo{pj:C.proj_init_info(in)}
}

// InitName returns the name of init file, e.g. 'epsg'.
//
func (i *InitInfo) InitName () string {
    return C.GoString(&((*i).pj.name)[0])
}

// InitPath returns the full path of init file, e.g.
// 'C:\OSGeo4W64\share\proj\epsg'.
//
func (i *InitInfo) InitPath () string {
    return C.GoString(&((*i).pj.filename)[0])
}

// VersionNumber returns the version number of init-file, e.g. '9.0.0'.
//
func (i *InitInfo) VersionNumber () string {
    return C.GoString(&((*i).pj.version)[0])
}

// Authority returns the originating entity of the init file, e.g. 'EPSG'.
//
func (i *InitInfo) Authority () string {
    return C.GoString(&((*i).pj.origin)[0])
}

// LastUpdateDate returns the date of last update of the init-file.
//
func (i *InitInfo) LastUpdateDate () string {
    return C.GoString(&((*i).pj.lastupdate)[0])
}

