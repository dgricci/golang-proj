package proj

import (
    "testing"
    "reflect"
)

// Tests :

// TestPrimeMeridian checks creating the Greenwich's prime meridian !
func TestPrimeMeridian ( t *testing.T ) {
    s := "Greenwich"
    _, e := NewPrimeMeridian(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation od '%s' PrimeMeridian", s)
    }
    s = "PSG:8901"
    _, e = NewPrimeMeridian(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation of '%s' PrimeMeridian", s)
    }
    s = "urn:ogc:def:ellipsoid::EPSG:code"
    _, e = NewPrimeMeridian(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation of '%s' PrimeMeridian", s)
    }
    s = "EPSG:8901"
    pm, e := NewPrimeMeridian(ctx, s)
    if e != nil {
        t.Error(e)
    }
    if pm.TypeOf() != PrimeMeridianType {
        t.Errorf("Expected PrimeMeridianType")
    }
    _, e = pm.Longitude(ctx)
    if e != nil {
        t.Error(e)
    }
    _, e = pm.ToRad(ctx)
    if e != nil {
        t.Error(e)
    }
    _, e = pm.Unit(ctx)
    if e != nil {
        t.Error(e)
    }
    pm.DestroyPrimeMeridian()
    if reflect.ValueOf(pm.Handle()).Elem() != reflect.Zero(reflect.TypeOf(pm.Handle())).Elem() {
        t.Errorf("Failed to deallocate the newly created PrimeMeridian '%s'", s)
    }
    if !pm.HandleIsNil() {
        t.Errorf("Failed to deallocate the newly created PrimeMeridian '%s'", s)
    }
}

// TestPrimeMeridianWKT checks failure when asking for wrong ID
func TestPrimeMeridianWKT ( t *testing.T ) {
    c := NewContext()
    defer c.DestroyContext()
    s := `SPHEROID["WGS 84",6378137,298.257223563,"unused"]`
    pm, e := NewPrimeMeridian(c, s)
    if e == nil {
        t.Errorf("Unexpected creation of '%s' PrimeMeridian", s)
    }
    s = `CONVERSION["PROJ-based coordinate operation",METHOD["PROJ-based operation method: +proj=lcc +lat_1=49 +lat_2=44 +lat_0=46.5 +lon_0=3 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m"]]`
    pm, e = NewPrimeMeridian(c, s)
    if e == nil {
        t.Errorf("Unexpected creation of '%s' PrimeMeridian", s)
    }
    t.Skip("WKT creation not (yet) supported for PrimeMeridian")
    s = `PRIMEM["Greenwich",0,ANGLEUNIT["degree",0.0174532925199433],ID["EPSG",8901]]`
    pm, e = NewPrimeMeridian(c, s)
    if e != nil {
        t.Error(e)
    }
    if pm.TypeOf() != PrimeMeridianType {
        t.Errorf("Expected PrimeMeridianType")
    }
    _, e = pm.Longitude(c)
    if e != nil {
        t.Error(e)
    }
    _, e = pm.ToRad(c)
    if e != nil {
        t.Error(e)
    }
    _, e = pm.Unit(c)
    if e != nil {
        t.Error(e)
    }
    pm.DestroyPrimeMeridian()
    if reflect.ValueOf(pm.Handle()).Elem() != reflect.Zero(reflect.TypeOf(pm.Handle())).Elem() {
        t.Errorf("Failed to deallocate the newly created PrimeMeridian '%s'", s)
    }
    if !pm.HandleIsNil() {
        t.Errorf("Failed to deallocate the newly created PrimeMeridian '%s'", s)
    }
}

