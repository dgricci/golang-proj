package proj

import (
    "testing"
    "reflect"
)

// Tests :

func TestEllipsoid ( t *testing.T ) {
    s := "GRS80"
    _, e := NewEllipsoid(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation od '%s' Ellipsoid", s)
    }
    s = "EPSG:7019"
    ell, e := NewEllipsoid(ctx, s)
    if e != nil {
        t.Error(e)
    }
    if Type(ell) != EllipsoidType {
        t.Errorf("Expected EllipsoidType")
    }
    _, e = ell.SemiMajor(ctx)
    if e != nil {
        t.Error(e)
    }
    _, _, e = ell.SemiMinor(ctx)
    if e != nil {
        t.Error(e)
    }
    _, e = ell.InverseFlattening(ctx)
    if e != nil {
        t.Error(e)
    }
    ell.DestroyEllipsoid()
    if reflect.ValueOf(ell.Handle()).Elem() != reflect.Zero(reflect.TypeOf(ell.Handle())).Elem() {
        t.Errorf("Failed to deallocate the newly created Ellipsoid '%s'", s)
    }
    if !ell.HandleIsNil() {
        t.Errorf("Failed to deallocate the newly created Ellipsoid '%s'", s)
    }
}

func TestEllipsoidWKT ( t *testing.T ) {
    c := NewContext()
    s := `ELLIPSOID["GRS 1980",6378137,298.257222101,LENGTHUNIT["metre",1],ID["EPSG",7019]]`
    ell, e := NewEllipsoid(c, s)
    if e != nil {
        t.Error(e)
    }
    if Type(ell) != EllipsoidType {
        t.Errorf("Expected EllipsoidType")
    }
    _, e = ell.SemiMajor(c)
    if e != nil {
        t.Error(e)
    }
    _, _, e = ell.SemiMinor(c)
    if e != nil {
        t.Error(e)
    }
    _, e = ell.InverseFlattening(c)
    if e != nil {
        t.Error(e)
    }
    ell.DestroyEllipsoid()
    if reflect.ValueOf(ell.Handle()).Elem() != reflect.Zero(reflect.TypeOf(ell.Handle())).Elem() {
        t.Errorf("Failed to deallocate the newly created Ellipsoid '%s'", s)
    }
    if !ell.HandleIsNil() {
        t.Errorf("Failed to deallocate the newly created Ellipsoid '%s'", s)
    }
}
