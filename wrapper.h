#ifndef PROJ_WRAPPER_H
#define PROJ_WRAPPER_H

#include <stdlib.h>
#include <stddef.h>  /* size_t */
#include <string.h>

#include "proj.h"

#ifdef __cplusplus
extern "C" {
#endif

char PROJ_DLL *getAuthorityFromPROJ ( PROJ_STRING_LIST l, int i );
char PROJ_DLL *listcat ( PROJ_STRING_LIST sl );
char PROJ_DLL **makeStringArray ( size_t l );
void PROJ_DLL setStringArrayItem ( const char **t, size_t i, const char *v);
const char PROJ_DLL *getStringArrayItem ( const char **t, size_t i);
void PROJ_DLL destroyStringArray ( char ***t );
double PROJ_DLL wrapper_proj_dmstor ( const char *dms );
int PROJ_DLL nbUnitsFromPROJ ( void );
PJ_UNITS PROJ_DLL *getUnitFromPROJ ( int i );

#ifdef __cplusplus
}
#endif

#endif /* PROJ_WRAPPER_H */
