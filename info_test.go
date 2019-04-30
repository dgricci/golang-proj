package proj

import (
    "testing"
)

// Tests :

// TestInfo gets information from underlaying PROJ library.
// Change this test when linking with a new PROJ library.
func TestInfo ( t *testing.T ) {
    r := Release()
    if r == "" {
        t.Errorf("Expected release set")
    }
    if r != "Rel. 6.0.0, March 1st, 2019" {
        t.Errorf("Expected release to equal 'Rel. 6.0.0, March 1st, 2019'")
    }
}

// TestGridInfo checks access to grid information
func TestGridInfo ( t *testing.T ) {
    ng := NewGridInfo("null")
    if ng.GridPath() == "" {
        t.Errorf("Expected path for 'null' grid")
    }
    if ng.GridName() != "null" {
        t.Errorf("Expected 'null' grid name to be 'null'")
    }
    g := NewGridInfo("nonexistinggrid")
    if g.GridPath() != "" {
        t.Errorf("Unexpected path for 'nonexistinggrid' grid")
    }
}

// TestInitInfo checks acess to init file information
func TestInitInfo ( t *testing.T ) {
    i := NewInitInfo("unknowninit")
    if i.InitPath() != "" {
        t.Errorf("Unexpected path for 'unknowninit' init file")
    }
    ei := NewInitInfo("world")
    if ei.InitPath() == "" {
        t.Errorf("Expected path for 'world' init file")
    }
    if ei.InitName() != "world" {
        t.Errorf("Expected 'world' init file name to be 'world'")
    }
    if ei.VersionNumber() == "" {
        t.Errorf("Expected version number for 'world' init file")
    }
    // Need to allow for "Unknown" until all commonly distributed EPSG-files comes with a metadata section
    if !(ei.Authority() == "EPSG" || ei.Authority() == "Unknown") {
        t.Errorf("Expected authory for 'world' init file")
    }
    if ei.LastUpdateDate() == "" {
        t.Errorf("Expected last update date for 'world' init file")
    }
}

