package proj

import (
    "fmt"
)

func ExampleInfo () {

    fmt.Printf("%s\n", Release())
    fmt.Printf("%s\n", VersionNumber())
    fmt.Printf("%d\n", Major())
    fmt.Printf("%d\n", Minor())
    fmt.Printf("%d\n", Patch())

    // Output:
    // Rel. 6.2.0, September 1st, 2019
    // 6.2.0
    // 6
    // 2
    // 0
}

func ExampleISOInfo_referencesystem () {
    c := NewContext()
    defer c.DestroyContext()
    p, _ := NewReferenceSystem(c, "EPSG:4326")
    defer p.DestroyReferenceSystem()
    pi := p.Info()
    fmt.Printf("id :%s\n", pi.ID())
    fmt.Printf("dsc:%s\n", pi.Description())
    fmt.Printf("def:%s\n", pi.Definition())
    fmt.Printf("inv:%t\n", pi.HasInverse())
    fmt.Printf("acc:%e\n", pi.Accuracy())

    u, _ := NewReferenceSystem(c, "+proj=utm +zone=32 +ellps=GRS80 +type=crs")
    defer u.DestroyReferenceSystem()
    ui := u.Info()
    fmt.Printf("id :%s\n", ui.ID())
    fmt.Printf("dsc:%s\n", ui.Description())
    fmt.Printf("def:%s\n", ui.Definition())
    fmt.Printf("inv:%t\n", ui.HasInverse())
    fmt.Printf("acc:%e\n", ui.Accuracy())

    // Output:
    // id :
    // dsc:WGS 84
    // def:
    // inv:false
    // acc:-1.000000e+00
    // id :
    // dsc:unknown
    // def:
    // inv:false
    // acc:-1.000000e+00

}

func ExampleISOInfo_operation () {
    c := NewContext()
    defer c.DestroyContext()
    o, _ := NewOperation(c, &Area{}, "EPSG:4326", "EPSG:32631")
    defer o.DestroyOperation()
    oi := o.Info()
    fmt.Printf("id :%s\n", oi.ID())
    fmt.Printf("dsc:%s\n", oi.Description())
    fmt.Printf("def:%s\n", oi.Definition())
    fmt.Printf("inv:%t\n", oi.HasInverse())
    fmt.Printf("acc:%e\n", oi.Accuracy())

    // Output:
    // id :pipeline
    // dsc:UTM zone 31N
    // def:proj=pipeline step proj=axisswap order=2,1 step proj=unitconvert xy_in=deg xy_out=rad step proj=utm zone=31 ellps=WGS84
    // inv:true
    // acc:0.000000e+00

}

func ExampleGridInfo () {
    ng := NewGridInfo("null")
    fmt.Printf("name:%s\n", ng.GridName())
    fmt.Printf("fmt :%s\n", ng.Format())
    fmt.Printf("%+7.2f %+7.2f\n", RadToDeg*ng.LowerLeft().λ(), RadToDeg*ng.LowerLeft().φ())
    fmt.Printf("%+7.2f %+7.2f\n", RadToDeg*ng.UpperRight().λ(), RadToDeg*ng.UpperRight().φ())
    fmt.Printf("%d x %d\n", ng.LongitudinalLen(), ng.LatitudinalLen())
    fmt.Printf("%+7.2f %+7.2f\n", RadToDeg*ng.LongitudinalCellSize(), RadToDeg*ng.LatitudinalCellSize())

    // Output:
    // name:null
    // fmt :ctable2
    // -180.00  -90.00
    // +360.00 +180.00
    // 3 x 3
    // +180.00  +90.00

}

