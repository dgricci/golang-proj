package proj

import (
    "testing"
    "math"
)

// Tests :

// TestRadianConversions checks radians to DMS both ways.
func TestRadianConversions ( t *testing.T) {
    buf := RadToDMS(math.Pi,'N','S')
    if "180dN" != buf {
        t.Errorf("Expected 180dN")
    }
    if DMSToRad(buf) != math.Pi {
        t.Errorf("Expected Pi")
    }
    buf = RadToDMS(-2, 'N', 'S')
    if `114d35'29.612"S` != buf {
        t.Errorf(`Expected 114d35'29.612"S`)
    }
    if math.Abs(DMSToRad(buf) + 2.0) > 1e-7 {
        t.Errorf("Expected -2")
    }
}

