package proj

import (
    "fmt"
)

func ExampleEllipsoid () {
    s := "EPSG:7019"
    c := NewContext()
    defer c.DestroyContext()
    grs80, e := NewEllipsoid(c, s)
    if e != nil {
        fmt.Println("Ooops")
        return
    }
    grs80I := grs80.Info()
    fmt.Printf("id :%s\n", grs80I.ID())
    fmt.Printf("dsc:%s (%s)\n", grs80I.Description(), grs80)
    fmt.Printf("def:%s\n", grs80I.Definition())
    a, b, isComputed, rf, e := grs80.Parameters(c)
    if e != nil {
        fmt.Println("Ooops (Parameters)")
    }
    fmt.Printf("a  :%7.2f\n", a)
    if (isComputed) {
        fmt.Printf("rf :%10.2f\n", rf)
    } else {
        fmt.Printf("b  :%10.2f\n", b)
    }
    fmt.Printf("proj-string : %s\n", grs80.ProjString(c, Version4))
    fmt.Printf("WKT : %s\n", grs80.Wkt(c, WKTv1GDAL, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))

    // Output:
    // id :
    // dsc:GRS 1980 (GRS 1980)
    // def:
    // a  :6378137.00
    // rf :    298.26
    // proj-string : +ellps=GRS80
    // WKT : SPHEROID["GRS 1980",6378137,298.257222101,AUTHORITY["EPSG","7019"]]


}

