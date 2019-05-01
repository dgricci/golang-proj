// Package proj implements a wrapper to the PROJ https://proj4.org library.
//
// (c) 2019 Didier Richard. All rights reserved.
// Use of this source code is governed by a Apache-style
// license that can be found in the LICENSE.md file.
package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"

import "unsafe"

var (
    // global information (private)
    projInfo *Info
)

// handle exposes methods to return and tests the private PROJ pointer.
// It implies type conversion to the right type.
//
// Internal use only.
//
type handle interface {
    Handle() interface{}    // return handle to the underlaying PROJ pointer
    HandleIsNil() bool      // return true if the underlying PROJ pointer is nil
}

// pj exposes handle methods as well as Info() to get metadata on the
// underlaying PROJ pointer.
//
// Internal use only.
//
type pj interface {
    handle
    Info() *ISOInfo         // return information about a specific object
}

// toString returns a string representation of the struct implementing a pj
// interface.
//
func toString ( o pj ) string {
    return o.Info().Description()
}

// toProj returns a proj-string representation of the struct
// implementing a pj interface.
// Empty string is returned on error.
//
func toProj ( ctx *Context, o pj, styp StringType ) string {
    cs := C.proj_as_proj_string((*ctx).pj, o.Handle().(*C.PJ), C.PJ_PROJ_STRING_TYPE(styp), nil)
    return C.GoString(cs)
}

// toWkt returns a WKT representation of the struct implementing a pj
// interface.
// Empty string is returned on error.
// Operation struct can only be exported to WKT2:2018 (WKTv2r2018 or
// WKTv2r2018Simplified for `styp`).
// `opts` can hold the following strings :
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
func toWkt ( ctx *Context, o pj, styp WKTType, opts []string ) string {
    var copts **C.char
    l := len(opts)
    if l > 0 {
        copts = C.makeStringArray(C.size_t(l+1))
        for i, opt := range opts {
            copt := C.CString(opt)
            C.setStringArrayItem(copts, C.size_t(i), copt)
        }
        C.setStringArrayItem(copts, C.size_t(l), nil)
    }
    cs := C.proj_as_wkt((*ctx).pj, o.Handle().(*C.PJ), C.PJ_WKT_TYPE(styp), copts)
    if l > 0 {
        for i := 0 ; i < l ; i++ {
            C.free(unsafe.Pointer(C.getStringArrayItem(copts, C.size_t(i))))
        }
        C.destroyStringArray(&copts)
    }
    return C.GoString(cs)
}

// init package initialisation
//
func init () {
    projInfo = &Info{pj:C.proj_info()}
}

