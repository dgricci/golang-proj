package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"

// Factors contains a opaque object describing various cartographic
// properties.
//
type Factors struct {
    pj C.PJ_FACTORS
}

// MeridionalScale returns the meridional scale calculated by ReferenceSystem.Factors.
//
func (f *Factors) MeridionalScale ( ) float64 {
    return float64((*f).pj.meridional_scale)
}

// ParallelScale returns the parallel scale calculated by ReferenceSystem.Factors.
//
func (f *Factors) ParallelScale ( ) float64 {
    return float64((*f).pj.parallel_scale)
}

// ArealScale returns the areal scale factor calculated by ReferenceSystem.Factors.
//
func (f *Factors) ArealScale ( ) float64 {
    return float64((*f).pj.areal_scale)
}

// AngularDistortion returns the angular distortion calculated by ReferenceSystem.Factors.
// Unit is radians.
//
func (f *Factors) AngularDistortion ( ) float64 {
    return float64((*f).pj.angular_distortion)
}

// MeridianParallelAngle returns the meridian/parallel angle, θ′, calculated by ReferenceSystem.Factors.
// Unit is radians.
//
func (f *Factors) MeridianParallelAngle ( ) float64 {
    return float64((*f).pj.meridian_parallel_angle)
}

// MeridianConvergence returns the meridian convergence, sometimes also
// described as grid declination, calculated by ReferenceSystem.Factors.
// Unit is radians.
//
func (f *Factors) MeridianConvergence ( ) float64 {
    return float64((*f).pj.meridian_convergence)
}

// MaximumScaleFactor returns the maximum scale factor calculated by ReferenceSystem.Factors.
//
func (f *Factors) MaximumScaleFactor ( ) float64 {
    return float64((*f).pj.tissot_semimajor)
}

// MinimumScaleFactor returns the minimum scale factor calculated by ReferenceSystem.Factors.
//
func (f *Factors) MinimumScaleFactor ( ) float64 {
    return float64((*f).pj.tissot_semiminor)
}

// PartialDerivativeXλ returns the partial derivative ∂x/∂λ calculated by ReferenceSystem.Factors.
//
func (f *Factors) PartialDerivativeXλ ( ) float64 {
    return float64((*f).pj.dx_dlam)
}

// PartialDerivativeYλ returns the partial derivative ∂y/∂λ calculated by ReferenceSystem.Factors.
//
func (f *Factors) PartialDerivativeYλ ( ) float64 {
    return float64((*f).pj.dy_dlam)
}

// PartialDerivativeXφ returns the partial derivative ∂x/∂φ calculated by ReferenceSystem.Factors.
//
func (f *Factors) PartialDerivativeXφ ( ) float64 {
    return float64((*f).pj.dx_dphi)
}

// PartialDerivativeYφ returns the partial derivative ∂y/∂φ calculated by ReferenceSystem.Factors.
//
func (f *Factors) PartialDerivativeYφ ( ) float64 {
    return float64((*f).pj.dy_dphi)
}

