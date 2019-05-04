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
    s := ""
    _, e = NewReferenceSystem(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation of CRS with an empty definition !")
    }
    s = "Invalid"  // Beware "Unknown" works ...
    _, e = NewReferenceSystem(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation of CRS with '%s' as definition !", s)
    }
    s = "PSG:4326"
    _, e = NewReferenceSystem(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation of CRS with '%s' as definition !", s)
    }
    s = "urn:ogc:def:crs::EPSG:code"
    _, e = NewReferenceSystem(ctx, s)
    if e == nil {
        t.Errorf("Unexpected creation of CRS with '%s' as definition !", s)
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
    if p.TypeOf() != Geographic2DCRS {
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

// TestCrsWkt checks WKT strings
func TestCrsWkt ( t *testing.T ) {
    c := NewContext()
    s := `SPHEROID["WGS 84",6378137,298.257223563,"unused"]`
    crs, e := NewReferenceSystem(c, s)
    if e == nil {
        t.Errorf("Unexpected creation of '%s' ReferenceSystem", s)
    }
    s = `CONVERSION["PROJ-based coordinate operation",METHOD["PROJ-based operation method: +proj=lcc +lat_1=49 +lat_2=44 +lat_0=46.5 +lon_0=3 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m"]]`
    crs, e = NewReferenceSystem(c, s)
    if e == nil {
        t.Errorf("Unexpected creation of '%s' ReferenceSystem", s)
    }
    s = `PROJCS["RGF93 / Lambert-93",GEOGCS["RGF93",DATUM["Reseau_Geodesique_Francais_1993",SPHEROID["GRS 1980",6378137,298.257222101,AUTHORITY["EPSG","7019"]],AUTHORITY["EPSG","6171"]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9122"]],AUTHORITY["EPSG","4171"]],PROJECTION["Lambert_Conformal_Conic_2SP"],PARAMETER["latitude_of_origin",46.5],PARAMETER["central_meridian",3],PARAMETER["standard_parallel_1",49],PARAMETER["standard_parallel_2",44],PARAMETER["false_easting",700000],PARAMETER["false_northing",6600000],UNIT["metre",1,AUTHORITY["EPSG","9001"]],AXIS["Easting",EAST],AXIS["Northing",NORTH],AUTHORITY["EPSG","2154"]]`
    crs, e = NewReferenceSystem(c, s)
    if e != nil {
        t.Error(e)
    }
    if crs.TypeOf() != ProjectedCRS {
        t.Errorf("Expected ProjectedCRS")
    }
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

