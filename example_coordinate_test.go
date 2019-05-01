package proj

import (
    "fmt"
)

func ExampleCoordinate () {
    c := NewCoordinate(2.3488000*DegToRad,48.8534100*DegToRad)
    fmt.Printf("Paris (λ, φ) in degrees : (%14.7f, %14.7f)\n", RadToDeg*c.λ(), RadToDeg*c.φ())
    fmt.Printf("Paris (λ, φ) in radians : (%14.10f, %14.10f)\n", c.λ(), c.φ())
    fmt.Printf("Paris (λ, φ) in grades  : (%14.7f, %14.7f)\n", RadToGrad*c.λ(), RadToGrad*c.φ())
    fmt.Printf("Paris (λ, φ) in dms     : (%14s, %14s)\n", RadToDMS(c.λ(), 'W', 'E'), RadToDMS(c.φ(), 'N', 'S'))

    // Output:
    // Paris (λ, φ) in degrees : (     2.3488000,     48.8534100)
    // Paris (λ, φ) in radians : (  0.0409942935,   0.8526528553)
    // Paris (λ, φ) in grades  : (     2.6097778,     54.2815667)
    // Paris (λ, φ) in dms     : (  2d20'55.68"W, 48d51'12.276"N)

}

