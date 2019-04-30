package proj

import (
    "testing"
)

// Tests :

const (
    F = "Expecting %s equals %.1f, but got %.1f"
)

// Test8bytesTofloat64 checks reading float64 from empty buffer
//
func Test8bytesTofloat64 ( t *testing.T ) {
    bits := []byte{}
    if convert8bytesTofloat64(bits) != 0.0 {
        t.Errorf("Reading empty []bytes should return 0.0")
    }
}

func tλ ( t *testing.T, c *Coordinate, V float64 ) {
    if c.λ() != V {
        t.Errorf(F, "λ", V, c.λ())
    }
}

func tφ ( t *testing.T, c *Coordinate, V float64 ) {
    if c.φ() != V {
        t.Errorf(F, "φ", V, c.φ())
    }
}

func tλφ ( t *testing.T, c *Coordinate, V []float64 ) {
    tλ(t,c,V[0])
    tφ(t,c,V[1])
}

func tE ( t *testing.T, c *Coordinate, V float64 ) {
    if c.E() != V {
        t.Errorf(F, "E", V, c.E())
    }
}

func tN ( t *testing.T, c *Coordinate, V float64 ) {
    if c.N() != V {
        t.Errorf(F, "N", V, c.N())
    }
}

func tEN ( t *testing.T, c *Coordinate, V []float64 ) {
    tE(t,c,V[0])
    tN(t,c,V[1])
}

func tX ( t *testing.T, c *Coordinate, V float64 ) {
    if c.X() != V {
        t.Errorf(F, "X", V, c.X())
    }
}

func tY ( t *testing.T, c *Coordinate, V float64 ) {
    if c.Y() != V {
        t.Errorf(F, "Y", V, c.Y())
    }
}

func tXY ( t *testing.T, c *Coordinate, V []float64 ) {
    tX(t,c,V[0])
    tY(t,c,V[1])
}

func tΩ ( t *testing.T, c *Coordinate, V float64 ) {
    if c.Ω() != V {
        t.Errorf(F, "Ω", V, c.Ω())
    }
}

func tΩφ ( t *testing.T, c *Coordinate, V []float64 ) {
    tΩ(t,c,V[0])
    tφ(t,c,V[1])
}

// TestCoord2D checks 2D coordinates
func TestCoord2D ( t *testing.T ) {
    V := []float64{1.0, 2.0}
    c := NewCoordinate(V[0],V[1])
    tλφ(t,c,V)
    tEN(t,c,V)
    tXY(t,c,V)
    tΩφ(t,c,V)
}

func th ( t *testing.T, c *Coordinate, V float64 ) {
    if c.h() != V {
        t.Errorf(F, "h", V, c.h())
    }
}

func tλφh ( t *testing.T, c *Coordinate, V []float64 ) {
    tλφ(t,c,V)
    th(t,c,V[2])
}

func tENh ( t *testing.T, c *Coordinate, V []float64 ) {
    tEN(t,c,V)
    th(t,c,V[2])
}

func tZ ( t *testing.T, c *Coordinate, V float64 ) {
    if c.Z() != V {
        t.Errorf(F, "Z", V, c.Z())
    }
}

func tXYZ ( t *testing.T, c *Coordinate, V []float64 ) {
    tXY(t,c,V)
    tZ(t,c,V[2])
}

func tκ ( t *testing.T, c *Coordinate, V float64 ) {
    if c.κ() != V {
        t.Errorf(F, "κ", V, c.κ())
    }
}

func tΩφκ ( t *testing.T, c *Coordinate, V []float64 ) {
    tΩφ(t,c,V)
    tκ(t,c,V[2])
}

// TestCoord3D checks 3D coordinates
func TestCoord3D ( t *testing.T ) {
    V := []float64{1.1, 2.1, 3.1}
    c := NewCoordinate(V[0],V[1],V[2])
    tλφh(t,c,V)
    tENh(t,c,V)
    tXYZ(t,c,V)
    tΩφκ(t,c,V)
}

func tt ( t *testing.T, c *Coordinate, V float64 ) {
    if c.t() != V {
        t.Errorf(F, "t", V, c.t())
    }
}

func tXYZt ( t *testing.T, c *Coordinate, V []float64 ) {
    tXYZ(t,c,V)
    tt(t,c,V[3])
}

func tλφht ( t *testing.T, c *Coordinate, V []float64 ) {
    tλφh(t,c,V)
    tt(t,c,V[3])
}

// TestCoord4D checks 4D coordinates
func TestCoord4D ( t *testing.T ) {
    V := []float64{1.2, 2.2, 3.2, 4.2}
    c := NewCoordinate(V[0],V[1],V[2],V[3])
    tXYZt(t,c,V)
    tλφht(t,c,V)
}

// TestComponents checks 2D, 3D and 4D coordinates
func TestComponents ( tst *testing.T ) {
    V := []float64{1.3, 2.3, 3.3, 4.3}
    c := NewCoordinate(V[0],V[1],V[2],V[3])
    var X, Y, Z, t float64
    X, Y = c.Components2D()
    if X != V[0] && Y != V[1] {
        tst.Errorf("Expecting (X, Y) equals (%.1f,%.1f), but got (%.1f,%.1f)", V[0], V[1], X, Y)
    }
    X, Y, Z = c.Components3D()
    if X != V[0] && Y != V[1] && Z != V[2] {
        tst.Errorf("Expecting (X, Y, Z) equals (%.1f,%.1f,%.1f), but got (%.1f,%.1f,%.1f)", V[0], V[1], V[2], X, Y, Z)
    }
    X, Y, Z, t = c.Components4D()
    if X != V[0] && Y != V[1] && Z != V[2] && t != V[3] {
        tst.Errorf("Expecting (X, Y, Z, t) equals (%.1f,%.1f,%.1f,%.1f), but got (%.1f,%.1f,%.1f,%.1f)", V[0], V[1], V[2], V[3], X, Y, Z, t)
    }
}

// TestLocatable on Coordinate
func TestLocatable ( tst *testing.T ) {
    c0 := NewCoordinate(1.0, 2.0, 3.0, 4.0)
    cc := c0.Location()
    if cc != c0 {
        tst.Errorf("Expecting c == %p, but got %p", c0, cc)
    }
    c0.SetLocation(cc)
    x, y, z, t := c0.Components4D()
    if x != 1.0 || y != 2.0 || z != 3.0 || t != 4.0 {
        tst.Errorf("Expecting 1.0, 2.0, 3.0, 4.0, but got %.1f, %.1f, %.1f, %.1f", x, y, z, t)
    }
    cc = NewCoordinate(4.0, 3.0, 2.0, 1.0)
    if cc == c0 {
        tst.Errorf("Expecting x == %p, but got %p", cc, c0)
    }
    c0.SetLocation(cc)
    x, y, z, t = c0.Components4D()
    if x != 4.0 || y != 3.0 || z != 2.0 || t != 1.0 {
        tst.Errorf("Expecting 4.0, 3.0, 2.0, 1.0, but got %.1f, %.1f, %.1f, %.1f", x, y, z, t)
    }
}

