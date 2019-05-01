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
    fmt.Printf("proj-string : %s\n", crs.ProjString(c, Version4))
    fmt.Printf("WKT : %s\n", crs.Wkt(c, WKTv1GDAL, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))
    crs.DestroyReferenceSystem()

    crs, _ = NewReferenceSystem(c, "EPSG:4326")
    if crs == nil {
        fmt.Println("Oops (wgs84g)")
        return
    }
    crsI = crs.Info()
    fmt.Printf("id :%s\n", crsI.ID())
    fmt.Printf("dsc:%s\n", crsI.Description())
    fmt.Printf("def:%s\n", crsI.Definition())
    fmt.Printf("inv:%t\n", crsI.HasInverse())
    fmt.Printf("acc:%e\n", crsI.Accuracy())
    fmt.Printf("proj-string : %s\n", crs.ProjString(c, Version4))
    fmt.Printf("WKT : %s\n", crs.Wkt(c, WKTv1GDAL, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))
    crs.DestroyReferenceSystem()

    crs, _ = NewReferenceSystem(c, epsg2154PROJString + " +type=crs")
    crsI = crs.Info()
    fmt.Printf("id :%s\n", crsI.ID())
    fmt.Printf("dsc:%s\n", crsI.Description())
    fmt.Printf("def:%s\n", crsI.Definition())
    fmt.Printf("inv:%t\n", crsI.HasInverse())
    fmt.Printf("acc:%e\n", crsI.Accuracy())
    fmt.Printf("proj-string : %s\n", crs.ProjString(c, Version4))
    fmt.Printf("WKT : %s\n", crs.Wkt(c, WKTv1GDAL, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))
    crs.DestroyReferenceSystem()

    // Output:
    // id :
    // dsc:unknown
    // def:
    // inv:false
    // acc:-1.000000e+00
    // proj-string : +proj=utm +zone=32 +ellps=GRS80 +units=m +no_defs +type=crs
    // WKT : PROJCS["unknown",GEOGCS["unknown",DATUM["Unknown_based_on_GRS80_ellipsoid",SPHEROID["GRS 1980",6378137,298.257222101,AUTHORITY["EPSG","7019"]]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9122"]]],PROJECTION["Transverse_Mercator"],PARAMETER["latitude_of_origin",0],PARAMETER["central_meridian",9],PARAMETER["scale_factor",0.9996],PARAMETER["false_easting",500000],PARAMETER["false_northing",0],UNIT["metre",1,AUTHORITY["EPSG","9001"]],AXIS["Easting",EAST],AXIS["Northing",NORTH]]
    // id :
    // dsc:WGS 84
    // def:
    // inv:false
    // acc:-1.000000e+00
    // proj-string : +proj=longlat +datum=WGS84 +no_defs +type=crs
    // WKT : GEOGCS["WGS 84",DATUM["WGS_1984",SPHEROID["WGS 84",6378137,298.257223563,AUTHORITY["EPSG","7030"]],AUTHORITY["EPSG","6326"]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9122"]],AUTHORITY["EPSG","4326"]]
    // id :
    // dsc:unknown
    // def:
    // inv:false
    // acc:-1.000000e+00
    // proj-string : +proj=lcc +lat_0=46.5 +lon_0=3 +lat_1=49 +lat_2=44 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m +no_defs +type=crs
    // WKT : PROJCS["unknown",GEOGCS["unknown",DATUM["Unknown_based_on_GRS80_ellipsoid",SPHEROID["GRS 1980",6378137,298.257222101,AUTHORITY["EPSG","7019"]],TOWGS84[0,0,0,0,0,0,0]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9122"]]],PROJECTION["Lambert_Conformal_Conic_2SP"],PARAMETER["latitude_of_origin",46.5],PARAMETER["central_meridian",3],PARAMETER["standard_parallel_1",49],PARAMETER["standard_parallel_2",44],PARAMETER["false_easting",700000],PARAMETER["false_northing",6600000],UNIT["metre",1,AUTHORITY["EPSG","9001"]],AXIS["Easting",EAST],AXIS["Northing",NORTH]]

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
    fmt.Printf("proj-string : %s\n", p.ProjString(ctx, Version4))
    fmt.Printf("WKT : %s\n", p.Wkt(ctx, WKTv1GDAL, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))
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
    fmt.Printf("proj-string : %s\n", p.ProjString(c, Version4))
    fmt.Printf("WKT : %s\n", p.Wkt(c, WKTv1GDAL, "MULTILINE=NO", "OUTPUT_AXIS=AUTO"))
    p.DestroyReferenceSystem()
    c.DestroyContext()

    // Output:
    // DatabasePath: ./usr/local/share/proj/proj.db
    // id :
    // dsc:RGF93 / Lambert-93
    // def:
    // inv:false
    // acc:-1.000000e+00
    // proj-string : +proj=lcc +lat_0=46.5 +lon_0=3 +lat_1=49 +lat_2=44 +x_0=700000 +y_0=6600000 +ellps=GRS80 +units=m +no_defs +type=crs
    // WKT : PROJCS["RGF93 / Lambert-93",GEOGCS["RGF93",DATUM["Reseau_Geodesique_Francais_1993",SPHEROID["GRS 1980",6378137,298.257222101,AUTHORITY["EPSG","7019"]],AUTHORITY["EPSG","6171"]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9122"]],AUTHORITY["EPSG","4171"]],PROJECTION["Lambert_Conformal_Conic_2SP"],PARAMETER["latitude_of_origin",46.5],PARAMETER["central_meridian",3],PARAMETER["standard_parallel_1",49],PARAMETER["standard_parallel_2",44],PARAMETER["false_easting",700000],PARAMETER["false_northing",6600000],UNIT["metre",1,AUTHORITY["EPSG","9001"]],AXIS["Easting",EAST],AXIS["Northing",NORTH],AUTHORITY["EPSG","2154"]]
    // DatabasePath: ./usr/local/share/proj/proj.db
    // id :
    // dsc:RGF93 / Lambert-93
    // def:
    // inv:false
    // acc:-1.000000e+00
    // proj-string : +proj=lcc +lat_0=46.5 +lon_0=3 +lat_1=49 +lat_2=44 +x_0=700000 +y_0=6600000 +ellps=GRS80 +units=m +no_defs +type=crs
    // WKT : PROJCS["RGF93 / Lambert-93",GEOGCS["RGF93",DATUM["Reseau_Geodesique_Francais_1993",SPHEROID["GRS 1980",6378137,298.257222101,AUTHORITY["EPSG","7019"]],AUTHORITY["EPSG","6171"]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9122"]],AUTHORITY["EPSG","4171"]],PROJECTION["Lambert_Conformal_Conic_2SP"],PARAMETER["latitude_of_origin",46.5],PARAMETER["central_meridian",3],PARAMETER["standard_parallel_1",49],PARAMETER["standard_parallel_2",44],PARAMETER["false_easting",700000],PARAMETER["false_northing",6600000],UNIT["metre",1,AUTHORITY["EPSG","9001"]],AXIS["Easting",EAST],AXIS["Northing",NORTH],AUTHORITY["EPSG","2154"]]

}

