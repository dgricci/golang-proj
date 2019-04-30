package proj

import (
    "testing"
    "reflect"
    "strings"
    "math"
)

// Tests :

// TestOperation checks creating a pipeline between two crs
func TestOperation ( t *testing.T ) {
    s25832 := "EPSG:25832"
    s25833 := "EPSG:25833"
    o, e := NewOperation(ctx, nil, s25832)
    if e == nil {
        t.Errorf("Expected NewOperation to failed (ReferenceSystem)")
    }
    o, e = NewOperation(ctx, &Area{}, s25832)
    if e == nil {
        t.Errorf("Expected NewOperation to failed (missing target CRS)")
    }
    o, e = NewOperation(ctx, &Area{}, s25832, s25833)
    if e != nil {
        t.Error(e)
    }
    o.DestroyOperation()
    if reflect.ValueOf(o.Handle()).Elem() != reflect.Zero(reflect.TypeOf(o.Handle())).Elem() {
        t.Errorf("Failed to deallocate the newly created transformation from '%s' to '%s'", s25832, s25833)
    }
    s := "urn:ogc:def:coordinateOperation:EPSG::1671" // RGF93 to WGS 84
    o, e = NewOperation(ctx, nil, s)
    if e != nil {
        t.Error(e)
    }
    if Type(o) != Transformation {
        t.Errorf("Expected Transformation ...")
    }
    o.DestroyOperation()
    if !o.HandleIsNil() {
        t.Errorf("Failed to deallocate the newly created transformation '%s'", s)
    }
    s = `CONVERSION["PROJ-based coordinate operation",METHOD["PROJ-based operation method: +proj=lcc +lat_1=49 +lat_2=44 +lat_0=46.5 +lon_0=3 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m"]]`
    o, e = NewOperation(ctx, nil, s)
    if e != nil {
        t.Error(e)
    }
    o.DestroyOperation()
    sv := []string{"proj=utm", "zone=32", "ellps=GRS80"}
    o, e = NewOperation(ctx, nil, strings.Join(sv," +"))
    if e != nil {
        t.Error(e)
    }
    o.DestroyOperation()
    o, e = NewOperation(ctx, nil, sv...)
    if e != nil {
        t.Error(e)
    }
    o.DestroyOperation()
}

// TestOperation_2 checks creating a pipeline between two crs and transform
// coordinate
func TestOperation_2 ( t *testing.T ) {
    s4326 := "EPSG:4326"
    s32631 := "EPSG:32631"
    o, e := NewOperation(ctx, &Area{}, s4326, s32631)
    if e != nil {
        t.Error(e)
    }
    defer o.DestroyOperation()
    a := NewCoordinate(0.0,3.0)         // latitude, longitude in degrees as expected by EPSG:4326
    i, _ := o.Transform(Forward, a)     // from EPSG:4326 to EPSG:32631
    b := i.(*Coordinate)
    if math.Abs(b.X() - 500000.0) > 1e-9 {
        t.Errorf("Expected X near 500000.0, but got %.1f", b.X())
    }
    if math.Abs(b.Y() - 0.0) > 1e-9 {
        t.Errorf("Expected Y near 0.0, but got %.1f", b.Y())
    }
    if _, e = o.Transform(Forward, b) ; e == nil {
        t.Errorf("Expected Transform to fail (wrong direction)")
    }
}

func TestOperation_2again (t *testing.T ) {
    s4326 := "EPSG:4326"
    s32631 := "EPSG:32631"
    crsS, _ := NewReferenceSystem(ctx, s4326)
    defer crsS.DestroyReferenceSystem()
    crsT, _ := NewReferenceSystem(ctx, s32631)
    defer crsT.DestroyReferenceSystem()
    ope, e := crsS.NewOperation(ctx, crsT)
    if e != nil {
        t.Error(e)
    }
    defer ope.DestroyOperation()
    a := NewCoordinate(0.0,3.0)         // latitude, longitude in degrees as expected by EPSG:4326
    i, _ := ope.Transform(Forward, a)     // from EPSG:4326 to EPSG:32631
    b := i.(*Coordinate)
    if math.Abs(b.X() - 500000.0) > 1e-9 {
        t.Errorf("Expected X near 500000.0, but got %.1f", b.X())
    }
    if math.Abs(b.Y() - 0.0) > 1e-9 {
        t.Errorf("Expected Y near 0.0, but got %.1f", b.Y())
    }
    if _, e = ope.Transform(Forward, b) ; e == nil {
        t.Errorf("Expected Transform to fail (wrong direction)")
    }
}

func TestOperation_3 (t *testing.T ) {
    sREUN47GAUSSL := "IGNF:REUN47GAUSSL"
    sRGAF09UTM20 := "IGNF:RGAF09UTM20"
    crsS, _ := NewReferenceSystem(ctx, sREUN47GAUSSL)
    defer crsS.DestroyReferenceSystem()
    crsT, _ := NewReferenceSystem(ctx, sRGAF09UTM20)
    defer crsT.DestroyReferenceSystem()
    filter := DefaultFilter
    filter.Authority = "IGNF"
    filter.Accuracy = 1.0
    filter.BBox = []float64{-21.42,-20.76,55.17,55.92}
    filter.XUse = IntersectionExtent
    filter.SCriterion = PartialIntersection
    filter.GUse = IgnoreGrids
    filter.AltGrid = false
    filter.PivotUse = WhenNoDirectTransformation
    filter.Pivots = make(map[string][]string)
    filter.Pivots["IGNF"] = []string{"LAMB93", "LAMBE"}
    ope, e := crsS.NewOperation(ctx, crsT, filter)
    if e == nil {
        opeI := ope.Info()
        t.Errorf("Unexpected IGNF operation between '%s' and '%s' : %s", sREUN47GAUSSL, sRGAF09UTM20, opeI.Definition())
    }
    defer ope.DestroyOperation()
}

func TestOperation_conversion ( t *testing.T ) {
    o, e := NewOperation(ctx, nil, "urn:ogc:def:coordinateOperation:EPSG::15593") // Geographic3D to Geographic2D
    if e != nil {
        t.Error(e)
    }
    if Type(o) != Conversion {
        t.Errorf("Expected Conversion ...")
    }
    o.DestroyOperation()
}

func TestOperation_concatenatedoperation ( t *testing.T ) {
    o, e := NewOperation(ctx, &Area{}, "EPSG:3758", "EPSG:2157") // Web-Mercator to Lambert-93
    if e != nil {
        t.Error(e)
    }
    if Type(o) != ConcatenatedOperation {
        t.Errorf("Expected ConcatenatedOperation ...")
    }
    o.DestroyOperation()
}

func TestOperation_othercoordinateoperation ( t *testing.T ) {
    t.Skip("Does not work ???")
    // NGF-IGN 1969 vers EVRF2000 (UELN-95/98)(EUROPEAN VERTICAL REFERENCE FRAME)
    o, e := NewOperation(ctx, nil, "IGNF:TSG1250")
    if e != nil {
        t.Error(e)
    }
    if Type(o) != OtherCoordinateOperation {
        t.Errorf("Expected OtherCoordinateOperation ...")
    }
    o.DestroyOperation()
}

