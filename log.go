package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/../usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"

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

