package proj

import (
    "testing"
    "os"
    "reflect"
)

// Tests :

// TestContext creates and destroy threading-context.
func TestContext ( t *testing.T) {
    pj := os.Getenv("PROJ_LIB")
    os.Setenv("PROJ_LIB",os.TempDir()) // fake path to prevent retrieval of proj.db
    c := NewContext()
    if c.HandleIsNil() {
        t.Errorf("Failed to create a new threading-context")
    }
    c.SetDatabasePath(os.TempDir()) // this should failed ... but PROJ keeps the previous path which is wrong too !
    if c.DatabasePath() != "" {
        t.Errorf("Unexpected database path %s for threading-context", c.DatabasePath())
    }
    if pj != "" {
        os.Setenv("PROJ_LIB",pj) // back to the right path ... 
    }
    if c.DatabasePath() == "" {
        t.Errorf("Expected database path for threading-context")
    }
    c.DestroyContext()
    if reflect.ValueOf(c.Handle()).Elem() != reflect.Zero(reflect.TypeOf(c.Handle())).Elem() {
        t.Errorf("Failed to deallocate the newly created threading-context")
    }
}

// TestDefaultContext checks the default context.
func TestDefaultContext ( t *testing.T) {
    if ctx.HandleIsNil() {
        t.Errorf("Failed to create tests context safe thread")
    }
    if ctx.DatabasePath() == "" {
        t.Errorf("Expected database path for default context")
    }
    s := "EPSG"
    if !ctx.IsAnAuthority(s) {
        t.Errorf("Expected '%s' to be an authority", s)
    }
    s = "UnKnownAuthority"
    if ctx.IsAnAuthority(s) {
        t.Errorf("Unexpected '%s' to be an authority", s)
    }
}

