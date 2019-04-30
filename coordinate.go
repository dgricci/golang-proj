package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"

import (
    "bytes"
    "encoding/binary"
)

// Coordinate holds a data type for generic geodetic 3D data plus epoch information
// It is seen as an array of bytes as it is a union. It holds a vector of 4
// doubles in C.
//
type Coordinate struct {
    pj C.PJ_COORD
}

// NewCoordinate initializer
//
func NewCoordinate (coords ...float64) *Coordinate {
    var x, y, z, t C.double
    switch len(coords) {
    case 4 :
        t = (C.double)(coords[3])
        fallthrough
    case 3 :
        z = (C.double)(coords[2])
        fallthrough
    case 2 :
        y = (C.double)(coords[1])
        fallthrough
    case 1 :
        x = (C.double)(coords[0])
    }
    return &Coordinate{pj:C.proj_coord(x, y, z, t)}
}

// convert8bytesTofloat64 converts [8]byte from C.double to float64
//
func convert8bytesTofloat64 ( bits []byte ) (v float64) {
    buf := bytes.NewBuffer(bits)
    var cval C.double
    if err := binary.Read(buf, binary.LittleEndian, &cval); err == nil {
        v = float64(cval)
    } else {
        v = 0.0
    }
    return
}

// get1stComponentCoordinate returns the 1st coordinate's component
//
func (c *Coordinate) get1stComponentCoordinate () (float64) {
                                          // <- from 1st byte (0) up to 8th byte excluded
    return convert8bytesTofloat64(((*c).pj)[ 0: 8])
}

// get2ndComponentCoordinate returns the 2nd coordinate's component
//
func (c *Coordinate) get2ndComponentCoordinate () (float64) {
    return convert8bytesTofloat64(((*c).pj)[ 8:16])
}

// get3rdComponentCoordinate returns the 3rd coordinate's component
//
func (c *Coordinate) get3rdComponentCoordinate () (float64) {
    return convert8bytesTofloat64(((*c).pj)[16:24])
}

// get4thComponentCoordinate returns the 4th coordinate's component
//
func (c *Coordinate) get4thComponentCoordinate () (float64) {
    return convert8bytesTofloat64(((*c).pj)[24:32])
}

// λ returns the longitude in radians
//
// note : under vim, hit <Ctrl-k>l* (digraph)
//
func (c *Coordinate) λ () (float64) {
    return c.get1stComponentCoordinate()
}

// φ returns the latitude in radians or
// the second rotation angle phi
//
// note : under vim, hit <Ctrl-k>f* (digraph)
//
func (c *Coordinate) φ () (float64) {
    return c.get2ndComponentCoordinate()
}

// X returns the abscissa in meters
//
func (c *Coordinate) X () (float64) {
    return c.get1stComponentCoordinate()
}

// Y returns the ordinate in meters
//
func (c *Coordinate) Y () (float64) {
    return c.get2ndComponentCoordinate()
}

// E returns the easting in meters
//
func (c *Coordinate) E () (float64) {
    return c.get1stComponentCoordinate()
}

// N returns the northing in meters
//
func (c *Coordinate) N () (float64) {
    return c.get2ndComponentCoordinate()
}

// Z returns the altitude depending on the
// underlaying CRS in meters
//
func (c *Coordinate) Z () (float64) {
    return c.get3rdComponentCoordinate()
}

// h returns the height
//
func (c *Coordinate) h () (float64) {
    return c.get3rdComponentCoordinate()
}

// Ω returns the first rotation angle omega
//
// note : under vim, hit <Ctrl-k>W* (digraph)
//
func (c *Coordinate) Ω () (float64) {
    return c.get1stComponentCoordinate()
}

// κ returns the third rotation angle kappa
//
// note : under vim, hit <Ctrl-k>k* (digraph)
//
func (c *Coordinate) κ () (float64) {
    return c.get3rdComponentCoordinate()
}

// t returns the time value
//
func (c *Coordinate) t () (float64) {
    return c.get4thComponentCoordinate()
}

// Components4D returns the quadruplet representing this coordinate :
// could be (X, Y, Z, t), (λ, φ, h, t)
//
func (c *Coordinate) Components4D () (float64, float64, float64, float64) {
    return c.get1stComponentCoordinate(), c.get2ndComponentCoordinate(),
           c.get3rdComponentCoordinate(), c.get4thComponentCoordinate()
}

// Components3D returns the triplet representing this coordinate :
// could be (Ω, φ, κ), (E, N, h), (X, Y, Z), (λ, φ, h)
//
func (c *Coordinate) Components3D () (float64, float64, float64) {
    return c.get1stComponentCoordinate(), c.get2ndComponentCoordinate(), c.get3rdComponentCoordinate()
}

// Components2D returns the pair representing this coordinate :
// could be (E, N), (X, Y), (λ, φ)
//
func (c *Coordinate) Components2D () (float64, float64) {
    return c.get1stComponentCoordinate(), c.get2ndComponentCoordinate()
}

// Locatable allows working on coordinates for a type 
//
type Locatable interface {
    // Location returns coordinates from the interface
    Location   ()               (*Coordinate)
    // SetLocation assigns coordinates to the interface
    SetLocation( *Coordinate )
}

// Location returns coordinates. Here itself !
//
func (c *Coordinate) Location() (xyzt *Coordinate) {
    xyzt = c
    return
}

// SetLocation sets coordinates.
//
func (c *Coordinate) SetLocation ( xyzt *Coordinate ) {
    if c == xyzt { return }
    x, y, z, t := xyzt.Components4D()
    (*c).pj = C.proj_coord(C.double(x), C.double(y), C.double(z), C.double(t))
}

