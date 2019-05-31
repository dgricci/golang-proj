package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/../usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"

import (
    "os"
    "log"
    "fmt"
)

// LoggingLevel lists logging levels in PROJ. Used to set the logging level in PROJ.
//
type LoggingLevel C.PJ_LOG_LEVEL
const (
    // None don’t log anything.
    None LoggingLevel = C.PJ_LOG_NONE
    // Error log only errors.
    Error LoggingLevel = C.PJ_LOG_ERROR
    // Debug log errors and additional debug information.
    Debug LoggingLevel = C.PJ_LOG_DEBUG
    // Trace highest logging level. Log everything including very detailed debug information.
    Trace  LoggingLevel = C.PJ_LOG_TRACE
)

var (
  qlog *log.Logger
  // LoggerPrefix for logged messages
  LoggerPrefix = "[proj]: "
)

// LogOnCError wraps message from PROJ to this logger
//
//export LogOnCError
func LogOnCError ( err *C.char ) {
    LogOnError(fmt.Errorf(C.GoString(err)))
}

// LogOnError function helper for logging error
//
func LogOnError ( err error ) {
    if err.(error) != nil {
        qlog.Printf("%v", err.(error))
    }
}

// Log function returns the underlaying logger
//
func Log () *log.Logger {
    return qlog
}

// SetLog function overrides the C logging function with `logFuncToGo`
//
func SetLog ( ctx *Context ) {
    C.proj_log_func((*ctx).pj, nil, C.PJ_LOG_FUNCTION(C.logFuncToGo));
}

// init package initialisation
//
func init () {
    qlog= log.New(os.Stderr, LoggerPrefix, log.LstdFlags)
}

