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
        t.Errorf("Unexpected creation of '%s' Ellipsoid", s)
    }
    s = "PSG:7019"
    _, e = NewEllipsoid(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation of '%s' Ellipsoid", s)
    }
    s = "urn:ogc:def:ellipsoid::EPSG:code"
    _, e = NewEllipsoid(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation of '%s' Ellipsoid", s)
    }
    s = "EPSG:7019"
    ell, e := NewEllipsoid(ctx, s)
    if e != nil {
        t.Error(e)
    }
    if ell.TypeOf() != EllipsoidType {
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
    s := `SPHEROID["WGS 84",6378137,298.257223563,"unused"]`
    ell, e := NewEllipsoid(c, s)
    if e == nil {// fails because 6.0.0 does not allow PM creation via WKT ...
        t.Errorf("Unexpected creation of '%s' Ellipsoid", s)
    }
    s = `CONVERSION["PROJ-based coordinate operation",METHOD["PROJ-based operation method: +proj=lcc +lat_1=49 +lat_2=44 +lat_0=46.5 +lon_0=3 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m"]]`
    ell, e = NewEllipsoid(c, s)
    if e == nil {// fails because 6.0.0 does not allow PM creation via WKT ...
        t.Errorf("Unexpected creation of '%s' Ellipsoid", s)
    }
    s = `ELLIPSOID["GRS 1980",6378137,298.257222101,LENGTHUNIT["metre",1],ID["EPSG",7019]]`
    ell, e = NewEllipsoid(c, s)
    if e != nil {
        t.Error(e)
    }
    if ell.TypeOf() != EllipsoidType {
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
