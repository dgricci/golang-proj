package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"
import "fmt"

// Unit contains data needed to define distance units
//
type Unit struct {
    pj *C.PJ_UNITS
}

var (
    units map[string]*Unit
)

// GetUnitByID returns an ellipsoid from its identifier
//
func GetUnitByID ( id string ) (u *Unit, e error) {
    if u = units[id] ; u == nil {
        e = fmt.Errorf("No Unit with that identifier '%s'", id)
    }
    return
}

// ID returns the keyword name of the unit.
//
func (u *Unit) ID () string {
    return C.GoString((*u).pj.id)
}

// ToMeterString returns a text representation of the factor that converts a given
// unit to meters.
//
func (u *Unit) ToMeterString () string {
    return C.GoString((*u).pj.to_meter)
}

// ToMeter returns the conversion factor that converts the unit to meters.
//
func (u *Unit) ToMeter () float64 {
    return float64((*u).pj.factor)
}

// Name returns the name of the unit.
//
func (u *Unit) Name () string {
    return C.GoString((*u).pj.name)
}

// Handle returns the PROJ internal object to be passed to the PROJ library
//
func (u *Unit) Handle () (interface{}) {
    return (*u).pj
}

// HandleIsNil returns true when the PROJ internal object is NULL.
//
func (u *Unit) HandleIsNil () bool {
    return (*u).pj == (*C.PJ_UNITS)(nil)
}

// init package initialisation
//
func init () {
    lus := int(C.nbUnitsFromPROJ())
    units = make(map[string]*Unit)
    for i := 0 ; i < lus ; i++ {
        u := &Unit{pj:nil}
        (*u).pj = C.getUnitFromPROJ(C.int(i))
        units[u.ID()] = u
    }
}

