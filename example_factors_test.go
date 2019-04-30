package proj

import (
    "fmt"
)

func ExampleFactors () {
    c := NewContext()
    defer c.DestroyContext()
    op, e := NewOperation(c, nil, epsg2154PROJString)
    defer op.DestroyOperation()
    if e != nil {
        fmt.Println(e)
        return
    }
    // a coordinate representing Paris: 2d20'55.68"E 48d51'12.276"N
    // note: PROJ works in radians, hence the DegToRad conversion
    a := NewCoordinate(2.3488000*DegToRad, 48.8534100*DegToRad)
    fmt.Printf("Paris (geographical coordinate) : %s %s\n", RadToDMS(a.λ(),'E','W'), RadToDMS(a.φ(),'N','S'))

    f, e := op.Factors(a)
    if e != nil {
        fmt.Println(e)
        return
    }

    fmt.Printf("MeridionalScale       : %7.2f\n", f.MeridionalScale())
    fmt.Printf("ParallelScale         : %7.2f\n", f.ParallelScale())
    fmt.Printf("ArealScale            : %7.2f\n", f.ArealScale())
    fmt.Printf("AngularDistortion     : %7.2f\n", f.AngularDistortion()*RadToDeg)
    fmt.Printf("MeridianParallelAngle : %7.2f\n", f.MeridianParallelAngle()*RadToDeg)
    fmt.Printf("MeridianConvergence   : %7.2f\n", f.MeridianConvergence()*RadToDeg)
    fmt.Printf("MaximumScaleFactor    : %7.2f\n", f.MaximumScaleFactor())
    fmt.Printf("MinimumScaleFactor    : %7.2f\n", f.MinimumScaleFactor())
    fmt.Printf("PartialDerivativeXλ   : %7.2f\n", f.PartialDerivativeXλ())
    fmt.Printf("PartialDerivativeYλ   : %7.2f\n", f.PartialDerivativeYλ())
    fmt.Printf("PartialDerivativeXφ   : %7.2f\n", f.PartialDerivativeXφ())
    fmt.Printf("PartialDerivativeYφ   : %7.2f\n", f.PartialDerivativeYφ())

    // Output:
    // Paris (geographical coordinate) : 2d20'55.68"E 48d51'12.276"N
    // MeridionalScale       :    1.00
    // ParallelScale         :    1.00
    // ArealScale            :    1.00
    // AngularDistortion     :    0.00
    // MeridianParallelAngle :   90.00
    // MeridianConvergence   :   -0.47
    // MaximumScaleFactor    :    1.00
    // MinimumScaleFactor    :    1.00
    // PartialDerivativeXλ   :    0.66
    // PartialDerivativeYλ   :   -0.01
    // PartialDerivativeXφ   :    0.01
    // PartialDerivativeYφ   :    1.00
}

