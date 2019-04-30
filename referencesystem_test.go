package proj

import (
    "testing"
    "reflect"
)

// Tests :

func TestUnknownCrs ( t *testing.T ) {
    _, e := NewReferenceSystem(ctx)
    if e == nil {
        t.Errorf("Unexpected creation of CRS without definition !")
    }
    _, e = NewReferenceSystem(ctx, "")
    if e == nil {
        t.Errorf("Unexpected creation of CRS with an empty definition !")
    }
    _, e = NewReferenceSystem(ctx, "Invalid")   // Beware "Unknown" works ...
    if e == nil {
        t.Errorf("Unexpected creation of CRS with 'Invalid' as definition !")
    }
}

// TestCrs checks some proj-strings
func TestCrs ( t *testing.T ) {
    s4326 := "EPSG:4326"
    p, e := NewReferenceSystem(ctx, s4326)
    if e != nil {
        t.Error(e)
    }
    p.DestroyReferenceSystem()
    if reflect.ValueOf(p.Handle()).Elem() != reflect.Zero(reflect.TypeOf(p.Handle())).Elem() {
        t.Errorf("Failed to deallocate the newly created CRS '%s'", s4326)
    }
    s := "urn:ogc:def:crs:EPSG::4326"
    p, e = NewReferenceSystem(ctx, s)
    if e != nil {
        t.Error(e)
    }
    if Type(p) != Geographic2DCRS {
        t.Errorf("Expected Geographic2DCRS")
    }
    p.DestroyReferenceSystem()
    if !p.HandleIsNil() {
        t.Errorf("Failed to deallocate the newly created CRS '%s'", s)
    }
}

// TestCrsMultiProjString checks creating with multi proj-strings
func TestCrsMultiProjString ( t *testing.T ) {
    p, e := NewReferenceSystem(ctx, "proj=utm", "zone=32", "ellps=GRS80")
    if e == nil {
        t.Error("Expected CRS creation to fail !")
    }
    p, e = NewReferenceSystem(ctx, "proj=utm", "zone=32", "ellps=GRS80", "type=crs")
    if e != nil {
        t.Error(e)
    }
    p.DestroyReferenceSystem()
}

// TestEPSG2154 checks with two different contexts
func TestEPSG2154 ( t *testing.T ) {
    s := "EPSG:2154"
    p, e := NewReferenceSystem(ctx, s)
    if e != nil {
        t.Error(e)
    }
    p.DestroyReferenceSystem()
    c := NewContext()
    p, e = NewReferenceSystem(c, s)
    if e != nil {
        t.Error(e)
    }
    p.DestroyReferenceSystem()
}

