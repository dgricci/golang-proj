package proj

import (
    "fmt"
)

func ExampleReferenceSystem () {
    c := NewContext()
    defer c.DestroyContext()

    crs, _ := NewReferenceSystem(c, utm32PROJString + " +type=crs")
    if crs == nil {
        fmt.Println("Oops (utm)")
        return
    }
    crsI := crs.Info()
    fmt.Printf("id :%s\n", crsI.ID())
    fmt.Printf("dsc:%s\n", crsI.Description())
    fmt.Printf("def:%s\n", crsI.Definition())
    fmt.Printf("inv:%t\n", crsI.HasInverse())
    fmt.Printf("acc:%e\n", crsI.Accuracy())
    crs.DestroyReferenceSystem()

    crs, _ = NewReferenceSystem(c, "EPSG:4326")
    if crs == nil {
        fmt.Println("Oops (wgs84g)")
    }
    crsI = crs.Info()
    fmt.Printf("id :%s\n", crsI.ID())
    fmt.Printf("dsc:%s\n", crsI.Description())
    fmt.Printf("def:%s\n", crsI.Definition())
    fmt.Printf("inv:%t\n", crsI.HasInverse())
    fmt.Printf("acc:%e\n", crsI.Accuracy())
    crs.DestroyReferenceSystem()

    crs, _ = NewReferenceSystem(c, epsg2154PROJString + " +type=crs")
    crsI = crs.Info()
    fmt.Printf("id :%s\n", crsI.ID())
    fmt.Printf("dsc:%s\n", crsI.Description())
    fmt.Printf("def:%s\n", crsI.Definition())
    fmt.Printf("inv:%t\n", crsI.HasInverse())
    fmt.Printf("acc:%e\n", crsI.Accuracy())
    crs.DestroyReferenceSystem()

    // Output:
    // id :
    // dsc:unknown
    // def:
    // inv:false
    // acc:-1.000000e+00
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

func ExampleReferenceSystem_epsg () {
    s2154 := "EPSG:2154"
    fmt.Printf("DatabasePath: %s\n", ctx.DatabasePath())
    p, _ := NewReferenceSystem(ctx, s2154)
    pi := p.Info()
    fmt.Printf("id :%s\n", pi.ID())
    fmt.Printf("dsc:%s\n", pi.Description())
    fmt.Printf("def:%s\n", pi.Definition())
    fmt.Printf("inv:%t\n", pi.HasInverse())
    fmt.Printf("acc:%e\n", pi.Accuracy())
    p.DestroyReferenceSystem()
    c := NewContext()
    fmt.Printf("DatabasePath: %s\n", c.DatabasePath())
    p, _ = NewReferenceSystem(c, s2154)
    pi = p.Info()
    fmt.Printf("id :%s\n", pi.ID())
    fmt.Printf("dsc:%s\n", pi.Description())
    fmt.Printf("def:%s\n", pi.Definition())
    fmt.Printf("inv:%t\n", pi.HasInverse())
    fmt.Printf("acc:%e\n", pi.Accuracy())
    p.DestroyReferenceSystem()
    c.DestroyContext()

    // Output:
    // DatabasePath: ./usr/local/share/proj/proj.db
    // id :
    // dsc:RGF93 / Lambert-93
    // def:
    // inv:false
    // acc:-1.000000e+00
    // DatabasePath: ./usr/local/share/proj/proj.db
    // id :
    // dsc:RGF93 / Lambert-93
    // def:
    // inv:false
    // acc:-1.000000e+00
}

