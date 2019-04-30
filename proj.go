// Package proj implements a wrapper to the PROJ https://proj4.org library.
//
// (c) 2019 Didier Richard. All rights reserved.
// Use of this source code is governed by a Apache-style
// license that can be found in the LICENSE.md file.
package proj

/*
#cgo CFLAGS: -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "proj.h"
 */
import "C"

var (
    // global information (private)
    projInfo *Info
)

// proj exposes methods the return and tests the private PROJ pointer.
// It implies type conversion to the right type.
//
// Internal use only.
type proj interface {
    Handle() interface{}    // return handle to the underlying PROJ pointer
    HandleIsNil() bool      // return true if the underlying PROJ pointer is nil
}

// init package initialisation
//
func init () {
    projInfo = &Info{pj:C.proj_info()}
}

