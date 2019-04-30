package proj

/*
#cgo CFLAGS: -I${SRCDIR}/../usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/../usr/local/lib -lproj
#include "proj.h"
 */
import "C"

// GuessedWKTDialect lists guessed WKT 
//
type GuessedWKTDialect C.PJ_GUESSED_WKT_DIALECT
const (
    // GuessedWKTv2r2018 see WKTv2r2018
    GuessedWKTv2r2018 GuessedWKTDialect = C.PJ_GUESSED_WKT2_2018
    // GuessedWKTv2r2015 see WKTv2t2015
    GuessedWKTv2r2015 GuessedWKTDialect = C.PJ_GUESSED_WKT2_2015
    // GuessedWKTv1 see https://proj4.org/development/reference/cpp/cpp_general.html#namespaceWKT1
    GuessedWKTv1      GuessedWKTDialect = C.PJ_GUESSED_WKT1_GDAL
    // GuessedWKTv1ESRI is an ESRI variant of WKT1
    GuessedWKTv1ESRI  GuessedWKTDialect = C.PJ_GUESSED_WKT1_ESRI
    // GuessedWKTUnknown is either Not WKT or unrecognized WKT
    GuessedWKTUnknown GuessedWKTDialect = C.PJ_GUESSED_NOT_WKT
)

// WKTType lists WKT versions
//
type WKTType C.PJ_WKT_TYPE
const (
    // WKTv2r2015 conforming to ISO 19162:2015(E) / OGC 12-063r5 with all possible nodes and new keyword names.
    WKTv2r2015 WKTType = C.PJ_WKT2_2015
    // WKTv2r2015Simplified same as WKTv2t2015 with the following exceptions:
    //
    // * UNIT keyword used
    //
    // * ID node only on top element
    //
    // * No ORDER element in AXIS element
    //
    // * PRIMEM node omitted if it is Greenwich
    //
    // * ELLIPSOID.UNIT node omitted if it is UnitOfMeasure::METRE
    //
    // * PARAMETER.UNIT / PRIMEM.UNIT omitted if same as AXIS
    //
    // * AXIS.UNIT omitted and replaced by a common GEODCRS.UNIT if they are all the same on all axis
    WKTv2r2015Simplified WKTType = C.PJ_WKT2_2015_SIMPLIFIED
    // WKTv2r2018 Full WKT2 string, conforming to ISO 19162:2018 / OGC 18-010,
    // with all possible nodes and new keyword names. Non-normative list of
    // differences:
    //
    // * uses GEOGCRS / BASEGEOGCRS keywords for GeographicCRS
    WKTv2r2018 WKTType = C.PJ_WKT2_2018
    // WKTv2r2018Simplified same as WKTv2r2018 with the simplification rules
    // of WKTv2t2015Simplified
    WKTv2r2018Simplified WKTType = C.PJ_WKT2_2018_SIMPLIFIED
    // WKTv1GDAL is WKT1 as traditionally output by GDAL, deriving from OGC
    // 01-009. A notable departure from WKT1_GDAL with respect to OGC 01-009
    // is that in WKT1_GDAL, the unit of the PRIMEM value is always degrees.
    WKTv1GDAL WKTType = C.PJ_WKT1_GDAL
    // WKTv1ESRI is WKT1 as traditionally output by ESRI software, deriving
    // from OGC 99-049.
    WKTv1ESRI WKTType = C.PJ_WKT1_ESRI
)

// Category lists objects categories
//
type Category C.PJ_CATEGORY
const (
    // EllipsoidKind for Ellipsoid objects
    EllipsoidKind Category = C.PJ_CATEGORY_ELLIPSOID
    // PrimeMeridianKind for PrimeMeridian objects
    PrimeMeridianKind = C.PJ_CATEGORY_PRIME_MERIDIAN
    // DatumKind for Datum objects
    DatumKind = C.PJ_CATEGORY_DATUM
    // CRSKind for ReferenceSystem objects
    CRSKind = C.PJ_CATEGORY_CRS
    // OperationKind for Operation objects
    OperationKind = C.PJ_CATEGORY_COORDINATE_OPERATION
)

// ISOType lists objects types
//
type ISOType C.PJ_TYPE
const (
    // TypeUnknown stands for unknown category
    TypeUnknown ISOType = C.PJ_TYPE_UNKNOWN
    // EllipsoidType is for Ellipsoid objects
    EllipsoidType ISOType = C.PJ_TYPE_ELLIPSOID
    // PrimeMeridianType is for PrimeMeridian objects
    PrimeMeridianType ISOType = C.PJ_TYPE_PRIME_MERIDIAN
    // GeodeticReferenceFrame is for ? objects
    GeodeticReferenceFrame ISOType = C.PJ_TYPE_GEODETIC_REFERENCE_FRAME
    // DynamicGeodeticReferenceFrame is for ? objects
    DynamicGeodeticReferenceFrame ISOType = C.PJ_TYPE_DYNAMIC_GEODETIC_REFERENCE_FRAME
    // VerticalReferenceFrame is for ? objects
    VerticalReferenceFrame ISOType = C.PJ_TYPE_VERTICAL_REFERENCE_FRAME
    // DynamicVerticalReferenceFrame is for ? objects
    DynamicVerticalReferenceFrame ISOType = C.PJ_TYPE_DYNAMIC_VERTICAL_REFERENCE_FRAME
    // DatumEnsemble is for ? objects
    DatumEnsemble ISOType = C.PJ_TYPE_DATUM_ENSEMBLE
    // CRS is an abstract type
    CRS ISOType = C.PJ_TYPE_CRS
    // GeodeticCRS is for ReferenceSystem objects
    GeodeticCRS ISOType = C.PJ_TYPE_GEODETIC_CRS
    // GeocentricCRS is for ReferenceSystem objects
    GeocentricCRS ISOType = C.PJ_TYPE_GEOCENTRIC_CRS
    // GeographicCRS covers both Geographic2DCRS and Geographic3DCRS
    GeographicCRS ISOType = C.PJ_TYPE_GEOGRAPHIC_CRS
    // Geographic2DCRS is for ReferenceSystem objects
    Geographic2DCRS ISOType = C.PJ_TYPE_GEOGRAPHIC_2D_CRS
    // Geographic3DCRS is for ReferenceSystem objects
    Geographic3DCRS ISOType = C.PJ_TYPE_GEOGRAPHIC_3D_CRS
    // VerticalCRS is for ReferenceSystem objects
    VerticalCRS ISOType = C.PJ_TYPE_VERTICAL_CRS
    // ProjectedCRS is for ReferenceSystem objects
    ProjectedCRS ISOType = C.PJ_TYPE_PROJECTED_CRS
    // CompoundCRS is for ReferenceSystem objects
    CompoundCRS ISOType = C.PJ_TYPE_COMPOUND_CRS
    // TemporalCRS is for ReferenceSystem objects
    TemporalCRS ISOType = C.PJ_TYPE_TEMPORAL_CRS
    // EngineeringCRS is for ReferenceSystem objects
    EngineeringCRS ISOType = C.PJ_TYPE_ENGINEERING_CRS
    // BoundCRS is for ReferenceSystem objects
    BoundCRS ISOType = C.PJ_TYPE_BOUND_CRS
    // OtherCRS is for ReferenceSystem objects
    OtherCRS ISOType = C.PJ_TYPE_OTHER_CRS
    // Conversion is for Operation objects
    Conversion ISOType = C.PJ_TYPE_CONVERSION
    // Transformation is for Operation objects
    Transformation ISOType = C.PJ_TYPE_TRANSFORMATION
    // ConcatenatedOperation is for Operation objects
    ConcatenatedOperation ISOType = C.PJ_TYPE_CONCATENATED_OPERATION
    // OtherCoordinateOperation is for Operation objects
    OtherCoordinateOperation ISOType = C.PJ_TYPE_OTHER_COORDINATE_OPERATION
)

// ComparisonCriterion expresses comparison cases between ISO19111
// objects
//
type ComparisonCriterion C.PJ_COMPARISON_CRITERION
const (
    // Strict stands for all properties are identical
    Strict ComparisonCriterion = C.PJ_COMP_STRICT
    // Equivalent means that objects are equivalent for the purpose
    // of coordinate operations. They can differ by the name of their objects,
    // identifiers, other metadata. Parameters may be expressed in different
    // units, provided that the value is (with some tolerance) the same once
    // expressed in a common unit.
    Equivalent ComparisonCriterion = C.PJ_COMP_EQUIVALENT
    // EquivalentExceptAxisOrder is same as Equivalent relaxed
    // with an exception that the axis order of the base CRS of a
    // DerivedCRS/ProjectedCRS or the axis order of a GeographicCRS is
    // ignored. Only to be used with DerivedCRS/ProjectedCRS/GeographicCRS.
    EquivalentExceptAxisOrder ComparisonCriterion = C.PJ_COMP_EQUIVALENT_EXCEPT_AXIS_ORDER_GEOGCRS
)

// CRSExtentUse speficies how source and target CRS extent should be used to
// restrict candidate operations (only taken into account if no explicit area
// of interest is specified).
type CRSExtentUse C.PROJ_CRS_EXTENT_USE
const (
    // NoExtent ignores CRS extent
    NoExtent CRSExtentUse = C.PJ_CRS_EXTENT_NONE
    // BothExtent tests coordinate operation extent against both CRS extent
    BothExtent CRSExtentUse = C.PJ_CRS_EXTENT_BOTH
    // IntersectionExtent tests coordinate operation extent against both CRS extent
    IntersectionExtent CRSExtentUse = C.PJ_CRS_EXTENT_INTERSECTION
    // SmallestExtent tests coordinate operation against the smallest of both CRS extent
    SmallestExtent CRSExtentUse = C.PJ_CRS_EXTENT_SMALLEST
)

// GridAvailabilityUse describe how grid availability is used.
//
type GridAvailabilityUse C.PROJ_GRID_AVAILABILITY_USE
const (
    // SortGrids is only used for sorting results. Operations where some grids are missing will be sorted last.
    SortGrids GridAvailabilityUse = C.PROJ_GRID_AVAILABILITY_USED_FOR_SORTING
    // DiscardMissingGrid completely discards an operation if a required grid is missing.
    DiscardMissingGrid GridAvailabilityUse = C.PROJ_GRID_AVAILABILITY_DISCARD_OPERATION_IF_MISSING_GRID
    // IgnoreGrids ignores grid availability at all. Results will be presented as if all grids were available.
    IgnoreGrids GridAvailabilityUse = C.PROJ_GRID_AVAILABILITY_IGNORED
)

// StringType indicates the PROJ string version
//
type StringType C.PJ_PROJ_STRING_TYPE
const (
    // Version5 PROJ v5 (or later versions) string.
    Version5 StringType = C.PJ_PROJ_5
    // Version4 PROJ v4 string as output by GDAL exportToProj4()
    Version4 StringType = C.PJ_PROJ_4
)

// SpatialCriterion to restrict candidate operations
//
type SpatialCriterion C.PROJ_SPATIAL_CRITERION
const (
    // StrictContainment the area of validity of transforms should strictly contain the area of interest.
    StrictContainment SpatialCriterion = C.PROJ_SPATIAL_CRITERION_STRICT_CONTAINMENT
    // PartialIntersection the area of validity of transforms should at least intersect the area of interest.
    PartialIntersection SpatialCriterion = C.PROJ_SPATIAL_CRITERION_PARTIAL_INTERSECTION
)

// IntermediateCRSUse describes if and how intermediate CRS should be used
//
type IntermediateCRSUse C.PROJ_INTERMEDIATE_CRS_USE
const (
    // AlwaysUse searches for intermediate CRS.
    AlwaysUse IntermediateCRSUse = C.PROJ_INTERMEDIATE_CRS_USE_ALWAYS
    // WhenNoDirectTransformation only attempt looking for intermediate CRS if there is no direct transformation available.
    WhenNoDirectTransformation IntermediateCRSUse = C.PROJ_INTERMEDIATE_CRS_USE_IF_NO_DIRECT_TRANSFORMATION
    // NeverUse does not use intermediate CRS.
    NeverUse IntermediateCRSUse = C.PROJ_INTERMEDIATE_CRS_USE_NEVER
)

// CoordinateSystemType describes type of coordinate system.
//
type CoordinateSystemType C.PJ_COORDINATE_SYSTEM_TYPE
const (
    // UnknownCS for unknown coordinate system type
    UnknownCS CoordinateSystemType = C.PJ_CS_TYPE_UNKNOWN
    // CartesianCS for set of coordinates measured in the same unit
    // (usually meter).
    CartesianCS CoordinateSystemType = C.PJ_CS_TYPE_CARTESIAN
    // EllisoidalCS for set of coordinates surface that meet at right angles
    // on an ellisoid
    EllisoidalCS CoordinateSystemType = C.PJ_CS_TYPE_ELLIPSOIDAL
    // VerticalCS for a position along a vertical direction above or below a given vertical datum.
    VerticalCS CoordinateSystemType = C.PJ_CS_TYPE_VERTICAL
    // SphericalCS for set of coordinates surface that meet at right angles
    // on a sphere
    SphericalCS CoordinateSystemType = C.PJ_CS_TYPE_SPHERICAL
    // OrdinalCS for ...
    OrdinalCS CoordinateSystemType = C.PJ_CS_TYPE_ORDINAL
    // ParametricCS for set of functions
    ParametricCS CoordinateSystemType = C.PJ_CS_TYPE_PARAMETRIC
    // DateTimeTemporalCS for ...
    DateTimeTemporalCS CoordinateSystemType = C.PJ_CS_TYPE_DATETIMETEMPORAL
    // TemporalCountCS for ...
    TemporalCountCS CoordinateSystemType = C.PJ_CS_TYPE_TEMPORALCOUNT
    // TemporalMeasureCs for ...
    TemporalMeasureCs CoordinateSystemType = C.PJ_CS_TYPE_TEMPORALMEASURE
)

// Type returns the type of a `*ReferenceSystem`, `*Operation`, `Ellipsoid`,
// `PrimeMeridian` object.
//
func Type ( iso interface{} ) ISOType {
    switch iso.(type) {
    case *ReferenceSystem   :
        return ISOType(C.proj_get_type( ((iso.(*ReferenceSystem)).Handle()).(*C.PJ) ))
    case *Operation         :
        return ISOType(C.proj_get_type( ((iso.(*Operation)).Handle()).(*C.PJ) ))
    case *Ellipsoid         :
        return ISOType(C.proj_get_type( ((iso.(*Ellipsoid)).Handle()).(*C.PJ) ))
    case *PrimeMeridian     :
        return ISOType(C.proj_get_type( ((iso.(*PrimeMeridian)).Handle()).(*C.PJ) ))
    default                 :
        return TypeUnknown
    }
}
