package proj

import (
    "fmt"
)

func ExamplePrimeMeridian () {
    s := "EPSG:8901"
    c := NewContext()
    defer c.DestroyContext()
    greenwhich, e := NewPrimeMeridian(c, s)
    if e != nil {
        fmt.Println("Ooops")
        return
    }
    greenwhichI := greenwhich.Info()
    fmt.Printf("id :%s\n", greenwhichI.ID())
    fmt.Printf("dsc:%s\n", greenwhichI.Description())
    fmt.Printf("def:%s\n", greenwhichI.Definition())
    l, f, u, e := greenwhich.Parameters(c)
    if e != nil {
        fmt.Println("Ooops (Parameters)")
    }
    fmt.Printf("l :%10.2f\n", l)
    fmt.Printf("f :%10.2f\n", f)
    fmt.Printf("u :%s\n", u)
    // Output:
    // id :
    // dsc:Greenwich
    // def:
    // l :      0.00
    // f :      0.02
    // u :degree
}

