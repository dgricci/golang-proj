package proj

import (
    "testing"
    "math"
)

// Tests :

// TestFactors tests NewFactors and computed values
func TestFactors ( t *testing.T) {
    c := NewContext()
    defer c.DestroyContext()
    merc, e := NewOperation(c, nil, "+proj=merc +ellps=WGS84")
    defer merc.DestroyOperation()
    if merc == nil {
        t.Error(e)
        return
    }
    a := NewCoordinate(DegToRad*12, DegToRad*55,0.0,0.0)
    f, e := merc.Factors(a)
    if e != nil {
        t.Error(e)
        return
    }
    // check a few key characteristics of the Mercator projection
    if math.Abs(f.AngularDistortion() - 0.0) > 1e-7 {
        t.Errorf("Angular distortion should be 0.0")
    }
    if math.Abs(f.MeridianParallelAngle() - math.Pi/2.0) > 1e-7 {
        t.Errorf("Meridian/parallel angle should be 90 deg")
    }
    // meridian convergence should be 0
    if f.MeridianConvergence() != 0.0 {
        t.Errorf("Meridian convergence should be 0")
    }

    i, _ := merc.Transform(Forward, a)
    a = i.(*Coordinate)
    f, e = merc.Factors(a)
    if e == nil {
        t.Error("Unexpected Factors success on non geographical coordinates")
    }
}

