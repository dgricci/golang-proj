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
  // LoggerFlags for default logging
  LoggerFlags = log.LstdFlags
)

// LogOnCError wraps message from PROJ to this logger
//
//export LogOnCError
func LogOnCError ( msg *C.char ) {
    if msg != nil {
        LogOnError(fmt.Errorf(C.GoString(msg)))
        return
    }
}

// LogOnError is a helper for logging error
//
func LogOnError ( err error ) {
    if err != nil {
        qlog.Printf("%v", err.(error))
        return
    }
}

// Log returns the underlaying logger
//
func Log () *log.Logger {
    return qlog
}

// SetLog overrides the C logging function with `logFuncToGo`
//
func SetLog ( ctx *Context ) {
    C.proj_log_func((*ctx).pj, nil, C.PJ_LOG_FUNCTION(C.logFuncToGo));
}

// LogLevel returns the current log level of PROJ.
//
func LogLevel ( ctx *Context ) LoggingLevel {
    return (LoggingLevel)(C.proj_log_level( (*ctx).pj, C.PJ_LOG_TELL) )
}

// SetLogLevel assigns the log level of PROJ.
//
func SetLogLevel ( ctx *Context, lvl LoggingLevel ) {
    _ = C.proj_log_level( (*ctx).pj, (C.PJ_LOG_LEVEL)(lvl) )
}

// init package initialisation : logger writes to os.Stderr using LoggerPrefix
// and LoggerFlags as default values.
//
func init () {
    qlog= log.New(os.Stderr, LoggerPrefix, LoggerFlags)
}

