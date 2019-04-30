package proj

import (
    "testing"
    "reflect"
)

// Tests :

// TestM checks creating a meter !
func TestM ( t *testing.T ) {
    u, e := GetUnitByID("m")
    if e != nil {
        t.Error(e)
    }
    if reflect.ValueOf(u.Handle()).Elem() == reflect.Zero(reflect.TypeOf(u.Handle())).Elem() {
        t.Errorf("Failed to get the 'm' unit")
    }
    if u.HandleIsNil() {
        t.Errorf("Failed to get the 'm' unit (again)")
    }
    if u.ID() != "m" {
        t.Errorf("Expected 'm', but got '%s'", u.ID())
    }
}

// TestUnknownUnit checks failure when asking for wrong ID
func TestUnknownUnit ( t *testing.T ) {
    u, e := GetUnitByID("UnknownUnit")
    if e == nil {
        t.Errorf("Unexpected '%s' unit !", u.ID())
    }
}

