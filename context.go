package proj

/*
#cgo CFLAGS: -I. -I${SRCDIR}/usr/local/include
#cgo LDFLAGS: -L${SRCDIR}/usr/local/lib -lproj
#include "wrapper.h"
 */
import "C"
import "unsafe"

// Context handles an internal threads context of the PROJ library
//
type Context struct {
    pj *C.PJ_CONTEXT
}

var (
    authorities map[string]bool
)

// NewContext creates a new threading-context into the PROJ library.
//
func NewContext () (*Context) {
    return &Context{pj:C.proj_context_create()}
}

// DestroyContext deallocates the internal threading-context into the PROJ library.
//
func (ctx *Context) DestroyContext () {
    if (*ctx).pj != nil {
        C.proj_context_destroy((*ctx).pj)
        (*ctx).pj = nil
    }
}

// Handle returns the PROJ internal object to be passed to the PROJ library
//
func (ctx *Context) Handle () (interface{}) {
    return (*ctx).pj
}

// HandleIsNil returns true when the PROJ internal object is NULL.
//
func (ctx *Context) HandleIsNil () bool {
    return (*ctx).pj == (*C.PJ_CONTEXT)(nil)
}

// DatabasePath returns the path to the database, empty string if none.
//
func (ctx *Context) DatabasePath () string {
    p := C.proj_context_get_database_path((*ctx).pj)
    if p == nil { return "" }
    return C.GoString(p)
}

// SetDatabasePath assigns the path to the 'proj.db' file.
//
func (ctx *Context) SetDatabasePath ( p string ) {
    dbp := C.CString(p)
    defer C.free(unsafe.Pointer(dbp))
    _ = C.proj_context_set_database_path((*ctx).pj,dbp,nil,nil)
}

// IsAnAuthority checks whether the proposed name is an authority or not
//
func (ctx *Context) IsAnAuthority ( name string ) bool {
    return authorities[name]
}

// LogLevel returns the current log level of PROJ.
//
func (ctx *Context) LogLevel ( ) LoggingLevel {
    return (LoggingLevel)(C.proj_log_level( (*ctx).pj, C.PJ_LOG_TELL) )
}

// SetLogLevel assigns the log level of PROJ.
//
func (ctx *Context) SetLogLevel ( lvl LoggingLevel ) {
    _ = C.proj_log_level( (*ctx).pj, (C.PJ_LOG_LEVEL)(lvl) )
}

// init package initialisation
//
func init () {
    if auths := C.proj_get_authorities_from_database(nil) ; auths != nil {
        authorities = make(map[string]bool)
        for i := 0 ; i >= 0 ; i++ {
            cauth := C.getAuthorityFromPROJ(auths, C.int(i))
            if cauth == nil {
                break
            }
            authorities[C.GoString(cauth)] = true
        }
        C.proj_string_list_destroy(auths)
    }
}

