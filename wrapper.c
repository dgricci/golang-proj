/**
 * PROJ (http://proj4.org) wrapper in Golang
 * Release under the MIT License (MIT)
 * 
 * Copyright © `2019` `Didier RICHARD`
 *
 * Permission is hereby granted, free of charge, to any person
 * obtaining a copy of this software and associated documentation
 * files (the “Software”), to deal in the Software without
 * restriction, including without limitation the rights to use,
 * copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following
 * conditions:
 * 
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 * 
 * THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
 * OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
 * HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 **/ 

#include "wrapper.h"
#include  "_cgo_export.h"


char *getAuthorityFromPROJ ( PROJ_STRING_LIST l, int i ) {
    return l[i];
}

char *listcat ( PROJ_STRING_LIST sl ) {
    size_t l = 0;
    char *result = NULL;
    PROJ_STRING_LIST iterator = NULL;
    if (sl == NULL) return NULL ;
    for (iterator = sl; *iterator; iterator++) {
        l += strlen(*iterator);
    }
    result = (char *)malloc(l+1);
    if (result == NULL) return NULL;
    result[0] = '\0';
    for (iterator = sl; *iterator; iterator++) {
        result = strcat(result, *iterator);
    }
    return result;
}   

double wrapper_proj_dmstor ( const char *dms ) {
    char *rs = NULL;
    return proj_dmstor(dms,&rs);
}

char **makeStringArray ( size_t l ) {
    return (char **)calloc(l,sizeof(char*));
}   
void setStringArrayItem ( const char **t, size_t i, const char *v) {
    t[i] = v;
}   
const char *getStringArrayItem ( const char **t, size_t i) {
    return t[i];
}   
void destroyStringArray ( char ***t ) {
    free(*t);
    *t = NULL;
}   

int nbUnitsFromPROJ ( ) {
    int n = 0 ;
    PJ_UNITS *us;
    for (us = (PJ_UNITS *)proj_list_units(); us->id; us++) { n++; }
    return n;
}       
        
PJ_UNITS *getUnitFromPROJ ( int i ) {
    PJ_UNITS *us;
    us = (PJ_UNITS *)proj_list_units(); 
    return us+i;
}   

void logFuncToGo ( void *udata, int llvl, const char *emsg ) {
    switch (llvl) {
    case PJ_LOG_ERROR :
        LogOnCError((char *)emsg);
    default           :
    case PJ_LOG_DEBUG :
    case PJ_LOG_TRACE :
    case PJ_LOG_NONE  :
        return ;
    }
    return;
}

