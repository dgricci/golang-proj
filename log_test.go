package proj

import (
    "testing"
    "bytes"
    "strings"
    "fmt"
)

// Tests :


func TestLogOnError (t *testing.T) {
    w := Log().Writer()
    f := Log().Flags()
    p := Log().Prefix()
    var buf bytes.Buffer
    Log().SetOutput(&buf)
    Log().SetFlags(0)
    Log().SetPrefix("ERR")
    LogOnError(fmt.Errorf("%s", "Fake"))
    if strings.TrimSpace(buf.String()) != Log().Prefix()+"Fake" {
        t.Errorf("Expected '"+Log().Prefix()+"Fake' error, got '%q'", strings.TrimSpace(buf.String()))
    }
    LogOnCError(nil)
    if strings.TrimSpace(buf.String()) != Log().Prefix()+"Fake" {
        t.Errorf("Expected '"+Log().Prefix()+"Fake' error, got '%q'", strings.TrimSpace(buf.String()))
    }
    // back to defaults :
    Log().SetOutput(w)
    Log().SetFlags(f)
    Log().SetPrefix(p)
}

// TestLogging checks the logging levels.
func TestLogging ( t *testing.T) {
    lvl := LogLevel(ctx)
    if lvl != None {
        t.Errorf("Expected default log level to be '%d', but got '%d'", None, lvl)
    }
    SetLogLevel(ctx,Error)
    lvl = LogLevel(ctx)
    if lvl != Error {
        t.Errorf("Expected default log level to be '%d', but got '%d'", Error, lvl)
    }
    SetLogLevel(ctx,Debug)
    lvl = LogLevel(ctx)
    if lvl != Debug {
        t.Errorf("Expected default log level to be '%d', but got '%d'", Debug, lvl)
    }
    SetLogLevel(ctx,Trace)
    lvl = LogLevel(ctx)
    if lvl != Trace {
        t.Errorf("Expected default log level to be '%d', but got '%d'", Trace, lvl)
    }
    SetLogLevel(ctx,None)
}

