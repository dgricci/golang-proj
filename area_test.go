package proj

import (
    "testing"
    "reflect"
)

// Tests :

// TestWorldArea checks creating a world area !
func TestWorldArea ( t *testing.T ) {
    w := NewArea(-180,-90,180,90)
    w.DestroyArea()
    if reflect.ValueOf(w.Handle()).Elem() != reflect.Zero(reflect.TypeOf(w.Handle())).Elem() {
        t.Errorf("Failed to deallocate the newly created world area")
    }
    if !w.HandleIsNil() {
        t.Errorf("Failed to deallocate the newly created world area (again)")
    }
}

