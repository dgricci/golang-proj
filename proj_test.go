package proj

import (
    "testing"
    "os"
)

// Tests :
var (
    ctx *Context
)

func setup () {
    ctx = NewContext()
}

func teardown () {
    ctx.DestroyContext()
}

func TestMain ( m *testing.M ) {
    // call flag.Parse() here if TestMain uses flags
    setup()

    res := m.Run()

    teardown()

    os.Exit(res)
}

