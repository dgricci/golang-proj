package proj

import (
    "fmt"
)

func ExampleUnit () {
    c := NewContext()
    defer c.DestroyContext()
    // Get km unit from PROJ
    km, e := GetUnitByID("km")
    if e != nil {
        fmt.Println("Ooops")
        return
    }
    fmt.Printf("id=%s\n", km.ID())
    fmt.Printf("tom=%s\n", km.ToMeterString())
    fmt.Printf("scl=%20.15f\n", km.ToMeter())
    fmt.Printf("name=%s\n", km.Name())

    // Output:
    // id=km
    // tom=1000
    // scl=1000.000000000000000
    // name=Kilometer
}

