package proj

import (
    "fmt"
)

var (
    utm32PROJString    = "+proj=utm +zone=32 +ellps=GRS80"
    epsg2154PROJString = "+proj=lcc +lat_1=49 +lat_2=44 +lat_0=46.5 +lon_0=3 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m"
)

func ExampleOperation () {
    c := NewContext()
    defer c.DestroyContext()

    utm, _ := NewOperation(c, nil, utm32PROJString)
    if utm == nil {
        fmt.Println("Oops (utm)")
        return
    }
    defer utm.DestroyOperation()
    // a coordinate representing Copenhagen: 55d N, 12d E
    // note: PROJ works in radians, hence the DegToRad conversion
    a := NewCoordinate(DegToRad*12, DegToRad*55)
    fmt.Printf("Copenhagen (geographical coordinate) : %.0f %.0f\n", a.λ()*RadToDeg, a.φ()*RadToDeg)

    i, _ := utm.Transform(Forward, a)
    b := i.(*Coordinate)
    fmt.Printf("easting: %.2f, northing: %.2f\n", b.E(), b.N())

    i, _ = utm.Transform(Inverse, b)
    b = i.(*Coordinate)
    fmt.Printf("longitude: %s, latitude: %s\n", RadToDMS(b.λ(),'E','W'), RadToDMS(b.φ(),'N','S'))

    lcc, _ := NewOperation(c, nil, epsg2154PROJString)
    defer lcc.DestroyOperation()
    if lcc == nil {
        fmt.Println("Oops (lcc)")
        return
    }
    // a coordinate representing Paris: 2d20'55.68"E 48d51'12.276"N
    u := NewCoordinate(2.3488000*DegToRad, 48.8534100*DegToRad)
    fmt.Printf("Paris (geographical coordinate) : %s %s\n", RadToDMS(u.λ(),'E','W'), RadToDMS(u.φ(),'N','S'))

    i, _ = lcc.Transform(Forward, u)
    v := i.(*Coordinate)
    fmt.Printf("easting: %.2f, northing: %.2f\n", v.E(), v.N())

    i, _ = lcc.Transform(Inverse, v)
    v = i.(*Coordinate)
    fmt.Printf("longitude: %s, latitude: %s\n", RadToDMS(v.λ(),'E','W'), RadToDMS(v.φ(),'N','S'))

    // Output:
    // Copenhagen (geographical coordinate) : 12 55
    // easting: 691875.63, northing: 6098907.83
    // longitude: 12dE, latitude: 55dN
    // Paris (geographical coordinate) : 2d20'55.68"E 48d51'12.276"N
    // easting: 652216.64, northing: 6861682.61
    // longitude: 2d20'55.68"E, latitude: 48d51'12.276"N
}

// XYQS implements Locatable interface
type XYQS struct {
    x, y float64
    qualityLevel int
    surveyorName []byte
}
func (s *XYQS) Location () *Coordinate {
    return NewCoordinate(s.x, s.y, 23.45, 0.0)
}
func (s *XYQS) SetLocation ( c *Coordinate ) {
    x, y, _, _ := c.Components4D()
    s.x = x/1000.0
    s.y = y/1000.0
}

func ExampleOperation_locatable () {
    c := NewContext()
    defer c.DestroyContext()
    lcc, _ := NewOperation(c, nil, epsg2154PROJString)
    defer lcc.DestroyOperation()
    lccI := lcc.Info()
    fmt.Printf("id :%s\n", lccI.ID())
    fmt.Printf("dsc:%s (%s)\n", lccI.Description(), lcc)
    fmt.Printf("def:%s\n", lccI.Definition())
    fmt.Printf("inv:%t\n", lccI.HasInverse())
    fmt.Printf("acc:%e\n", lccI.Accuracy())
    fmt.Printf("proj-string : %s\n", lcc.ProjString(c, Version4))
    fmt.Printf("WKT : %s\n", lcc.Wkt(c, WKTv2r2018, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))

    // switch to radians ...
    survey := XYQS{x:2.3488000*DegToRad, y:48.8534100*DegToRad, qualityLevel:1, surveyorName:[]byte("me")}
    i, _ := lcc.Transform(Forward, &survey)
    survey = *(i.(*XYQS))
    fmt.Printf("easting: %.5f, northing: %.5f\n", survey.x, survey.y)

    // Output:
    // id :lcc
    // dsc:PROJ-based coordinate operation (PROJ-based coordinate operation)
    // def:proj=lcc lat_1=49 lat_2=44 lat_0=46.5 lon_0=3 x_0=700000 y_0=6600000 ellps=GRS80 towgs84=0,0,0,0,0,0,0 units=m
    // inv:true
    // acc:-1.000000e+00
    // proj-string : +proj=lcc +lat_1=49 +lat_2=44 +lat_0=46.5 +lon_0=3 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m
    // WKT : CONVERSION["PROJ-based coordinate operation",METHOD["PROJ-based operation method: +proj=lcc +lat_1=49 +lat_2=44 +lat_0=46.5 +lon_0=3 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m"]]
    // easting: 652.21664, northing: 6861.68261

}

func ExampleOperation_locatablearray () {
    s4326 := "EPSG:4326"
    s2154 := "EPSG:2154"
    c := NewContext()
    defer c.DestroyContext()
    wgs84g, _ := NewReferenceSystem(c, s4326)
    wgs84gI := wgs84g.Info()
    fmt.Printf("id :%s\n", wgs84gI.ID())
    fmt.Printf("dsc:%s (%s)\n", wgs84gI.Description(), wgs84g)
    fmt.Printf("def:%s\n", wgs84gI.Definition())
    fmt.Printf("inv:%t\n", wgs84gI.HasInverse())
    fmt.Printf("acc:%e\n", wgs84gI.Accuracy())
    fmt.Printf("proj-string : %s\n", wgs84g.ProjString(c, Version5))
    fmt.Printf("WKT : %s\n", wgs84g.Wkt(c, WKTv2r2018, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))
    wgs84g.DestroyReferenceSystem()
    epsg2154, _ := NewReferenceSystem(c, s2154)
    epsg2154I := epsg2154.Info()
    fmt.Printf("id :%s\n", epsg2154I.ID())
    fmt.Printf("dsc:%s (%s)\n", epsg2154I.Description(), epsg2154)
    fmt.Printf("def:%s\n", epsg2154I.Definition())
    fmt.Printf("inv:%t\n", epsg2154I.HasInverse())
    fmt.Printf("acc:%e\n", epsg2154I.Accuracy())
    fmt.Printf("proj-string : %s\n", epsg2154.ProjString(c, Version5))
    fmt.Printf("WKT : %s\n", epsg2154.Wkt(c, WKTv2r2018, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))
    epsg2154.DestroyReferenceSystem()
    wgs84gToEpsg2154, _ := NewOperation(c, &Area{}, s4326, s2154)
    defer wgs84gToEpsg2154.DestroyOperation()
    wgs84gToEpsg2154I := wgs84gToEpsg2154.Info()
    fmt.Printf("id :%s\n", wgs84gToEpsg2154I.ID())
    fmt.Printf("dsc:%s (%s)\n", wgs84gToEpsg2154I.Description(), wgs84gToEpsg2154)
    fmt.Printf("def:%s\n", wgs84gToEpsg2154I.Definition())
    fmt.Printf("inv:%t\n", wgs84gToEpsg2154I.HasInverse())
    fmt.Printf("acc:%e\n", wgs84gToEpsg2154I.Accuracy())
    fmt.Printf("proj-string : %s\n", wgs84gToEpsg2154.ProjString(c, Version5))
    fmt.Printf("WKT : %s\n", wgs84gToEpsg2154.Wkt(c, WKTv2r2018, "MULTILINE=YES", "INDENTATION_WIDTH=2", "OUTPUT_AXIS=AUTO"))

    var surveys [365]XYQS
    for r := 0 ; r < 365 ; r++ {
        // beware of EPSG:4326 axis order :)
        // the transformation reverse axis order AND switch to radians ...
        surveys[r] = XYQS{y:2.3488000, x:48.8534100, qualityLevel:1, surveyorName:[]byte("me")}
    }
    for r, survey := range surveys {
        i, _ := wgs84gToEpsg2154.Transform(Forward, &survey)
        surveys[r] = *(i.(*XYQS))
    }
    fmt.Printf("easting: %.5f, northing: %.5f\n", surveys[  0].x, surveys[  0].y)
    fmt.Printf("easting: %.5f, northing: %.5f\n", surveys[364].x, surveys[364].y)

    // Output:
    // id :
    // dsc:WGS 84 (WGS 84)
    // def:
    // inv:false
    // acc:-1.000000e+00
    // proj-string : +proj=longlat +datum=WGS84 +no_defs +type=crs
    // WKT : GEOGCRS["WGS 84",DATUM["World Geodetic System 1984",ELLIPSOID["WGS 84",6378137,298.257223563,LENGTHUNIT["metre",1]]],PRIMEM["Greenwich",0,ANGLEUNIT["degree",0.0174532925199433]],CS[ellipsoidal,2],AXIS["geodetic latitude (Lat)",north,ORDER[1],ANGLEUNIT["degree",0.0174532925199433]],AXIS["geodetic longitude (Lon)",east,ORDER[2],ANGLEUNIT["degree",0.0174532925199433]],USAGE[SCOPE["unknown"],AREA["World"],BBOX[-90,-180,90,180]],ID["EPSG",4326]]
    // id :
    // dsc:RGF93 / Lambert-93 (RGF93 / Lambert-93)
    // def:
    // inv:false
    // acc:-1.000000e+00
    // proj-string : +proj=lcc +lat_0=46.5 +lon_0=3 +lat_1=49 +lat_2=44 +x_0=700000 +y_0=6600000 +ellps=GRS80 +units=m +no_defs +type=crs
    // WKT : PROJCRS["RGF93 / Lambert-93",BASEGEOGCRS["RGF93",DATUM["Reseau Geodesique Francais 1993",ELLIPSOID["GRS 1980",6378137,298.257222101,LENGTHUNIT["metre",1]]],PRIMEM["Greenwich",0,ANGLEUNIT["degree",0.0174532925199433]]],CONVERSION["Lambert-93",METHOD["Lambert Conic Conformal (2SP)",ID["EPSG",9802]],PARAMETER["Latitude of false origin",46.5,ANGLEUNIT["degree",0.0174532925199433],ID["EPSG",8821]],PARAMETER["Longitude of false origin",3,ANGLEUNIT["degree",0.0174532925199433],ID["EPSG",8822]],PARAMETER["Latitude of 1st standard parallel",49,ANGLEUNIT["degree",0.0174532925199433],ID["EPSG",8823]],PARAMETER["Latitude of 2nd standard parallel",44,ANGLEUNIT["degree",0.0174532925199433],ID["EPSG",8824]],PARAMETER["Easting at false origin",700000,LENGTHUNIT["metre",1],ID["EPSG",8826]],PARAMETER["Northing at false origin",6600000,LENGTHUNIT["metre",1],ID["EPSG",8827]]],CS[Cartesian,2],AXIS["easting (X)",east,ORDER[1],LENGTHUNIT["metre",1]],AXIS["northing (Y)",north,ORDER[2],LENGTHUNIT["metre",1]],USAGE[SCOPE["unknown"],AREA["France"],BBOX[41.15,-9.86,51.56,10.38]],ID["EPSG",2154]]
    // id :pipeline
    // dsc:Inverse of RGF93 to WGS 84 (1) + Lambert-93 (Inverse of RGF93 to WGS 84 (1) + Lambert-93)
    // def:proj=pipeline step proj=axisswap order=2,1 step proj=unitconvert xy_in=deg xy_out=rad step proj=lcc lat_0=46.5 lon_0=3 lat_1=49 lat_2=44 x_0=700000 y_0=6600000 ellps=GRS80
    // inv:true
    // acc:1.000000e+00
    // proj-string : +proj=pipeline +step +proj=axisswap +order=2,1 +step +proj=unitconvert +xy_in=deg +xy_out=rad +step +proj=lcc +lat_0=46.5 +lon_0=3 +lat_1=49 +lat_2=44 +x_0=700000 +y_0=6600000 +ellps=GRS80
    // WKT : CONCATENATEDOPERATION["Inverse of RGF93 to WGS 84 (1) + Lambert-93",
    //   SOURCECRS[
    //     GEOGCRS["WGS 84",
    //       DATUM["World Geodetic System 1984",
    //         ELLIPSOID["WGS 84",6378137,298.257223563,
    //           LENGTHUNIT["metre",1]]],
    //       PRIMEM["Greenwich",0,
    //         ANGLEUNIT["degree",0.0174532925199433]],
    //       CS[ellipsoidal,2],
    //         AXIS["geodetic latitude (Lat)",north,
    //           ORDER[1],
    //           ANGLEUNIT["degree",0.0174532925199433]],
    //         AXIS["geodetic longitude (Lon)",east,
    //           ORDER[2],
    //           ANGLEUNIT["degree",0.0174532925199433]],
    //       USAGE[
    //         SCOPE["unknown"],
    //         AREA["World"],
    //         BBOX[-90,-180,90,180]],
    //       ID["EPSG",4326]]],
    //   TARGETCRS[
    //     PROJCRS["RGF93 / Lambert-93",
    //       BASEGEOGCRS["RGF93",
    //         DATUM["Reseau Geodesique Francais 1993",
    //           ELLIPSOID["GRS 1980",6378137,298.257222101,
    //             LENGTHUNIT["metre",1]]],
    //         PRIMEM["Greenwich",0,
    //           ANGLEUNIT["degree",0.0174532925199433]]],
    //       CONVERSION["Lambert-93",
    //         METHOD["Lambert Conic Conformal (2SP)",
    //           ID["EPSG",9802]],
    //         PARAMETER["Latitude of false origin",46.5,
    //           ANGLEUNIT["degree",0.0174532925199433],
    //           ID["EPSG",8821]],
    //         PARAMETER["Longitude of false origin",3,
    //           ANGLEUNIT["degree",0.0174532925199433],
    //           ID["EPSG",8822]],
    //         PARAMETER["Latitude of 1st standard parallel",49,
    //           ANGLEUNIT["degree",0.0174532925199433],
    //           ID["EPSG",8823]],
    //         PARAMETER["Latitude of 2nd standard parallel",44,
    //           ANGLEUNIT["degree",0.0174532925199433],
    //           ID["EPSG",8824]],
    //         PARAMETER["Easting at false origin",700000,
    //           LENGTHUNIT["metre",1],
    //           ID["EPSG",8826]],
    //         PARAMETER["Northing at false origin",6600000,
    //           LENGTHUNIT["metre",1],
    //           ID["EPSG",8827]]],
    //       CS[Cartesian,2],
    //         AXIS["easting (X)",east,
    //           ORDER[1],
    //           LENGTHUNIT["metre",1]],
    //         AXIS["northing (Y)",north,
    //           ORDER[2],
    //           LENGTHUNIT["metre",1]],
    //       USAGE[
    //         SCOPE["unknown"],
    //         AREA["France"],
    //         BBOX[41.15,-9.86,51.56,10.38]],
    //       ID["EPSG",2154]]],
    //   STEP[
    //     COORDINATEOPERATION["Inverse of RGF93 to WGS 84 (1)",
    //       SOURCECRS[
    //         GEOGCRS["WGS 84",
    //           DATUM["World Geodetic System 1984",
    //             ELLIPSOID["WGS 84",6378137,298.257223563,
    //               LENGTHUNIT["metre",1]]],
    //           PRIMEM["Greenwich",0,
    //             ANGLEUNIT["degree",0.0174532925199433]],
    //           CS[ellipsoidal,2],
    //             AXIS["geodetic latitude (Lat)",north,
    //               ORDER[1],
    //               ANGLEUNIT["degree",0.0174532925199433]],
    //             AXIS["geodetic longitude (Lon)",east,
    //               ORDER[2],
    //               ANGLEUNIT["degree",0.0174532925199433]]]],
    //       TARGETCRS[
    //         GEOGCRS["RGF93",
    //           DATUM["Reseau Geodesique Francais 1993",
    //             ELLIPSOID["GRS 1980",6378137,298.257222101,
    //               LENGTHUNIT["metre",1]]],
    //           PRIMEM["Greenwich",0,
    //             ANGLEUNIT["degree",0.0174532925199433]],
    //           CS[ellipsoidal,2],
    //             AXIS["geodetic latitude (Lat)",north,
    //               ORDER[1],
    //               ANGLEUNIT["degree",0.0174532925199433]],
    //             AXIS["geodetic longitude (Lon)",east,
    //               ORDER[2],
    //               ANGLEUNIT["degree",0.0174532925199433]]]],
    //       METHOD["Geocentric translations (geog2D domain)",
    //         ID["EPSG",9603]],
    //       PARAMETER["X-axis translation",0,
    //         LENGTHUNIT["metre",1],
    //         ID["EPSG",8605]],
    //       PARAMETER["Y-axis translation",0,
    //         LENGTHUNIT["metre",1],
    //         ID["EPSG",8606]],
    //       PARAMETER["Z-axis translation",0,
    //         LENGTHUNIT["metre",1],
    //         ID["EPSG",8607]],
    //       OPERATIONACCURACY[1.0],
    //       USAGE[
    //         SCOPE["unknown"],
    //         AREA["France"],
    //         BBOX[41.15,-9.86,51.56,10.38]],
    //       ID["INVERSE(EPSG)",1671]]],
    //   STEP[
    //     CONVERSION["Lambert-93",
    //       METHOD["Lambert Conic Conformal (2SP)",
    //         ID["EPSG",9802]],
    //       PARAMETER["Latitude of false origin",46.5,
    //         ANGLEUNIT["degree",0.0174532925199433],
    //         ID["EPSG",8821]],
    //       PARAMETER["Longitude of false origin",3,
    //         ANGLEUNIT["degree",0.0174532925199433],
    //         ID["EPSG",8822]],
    //       PARAMETER["Latitude of 1st standard parallel",49,
    //         ANGLEUNIT["degree",0.0174532925199433],
    //         ID["EPSG",8823]],
    //       PARAMETER["Latitude of 2nd standard parallel",44,
    //         ANGLEUNIT["degree",0.0174532925199433],
    //         ID["EPSG",8824]],
    //       PARAMETER["Easting at false origin",700000,
    //         LENGTHUNIT["metre",1],
    //         ID["EPSG",8826]],
    //       PARAMETER["Northing at false origin",6600000,
    //         LENGTHUNIT["metre",1],
    //         ID["EPSG",8827]],
    //       ID["EPSG",18085]]],
    //   USAGE[
    //     SCOPE["unknown"],
    //     AREA["France"],
    //     BBOX[41.15,-9.86,51.56,10.38]]]
    // easting: 652.21664, northing: 6861.68261
    // easting: 652.21664, northing: 6861.68261

}

