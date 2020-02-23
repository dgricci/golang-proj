#!/bin/bash
PROJ_PATH="${GOPATH}/src/osgeo.org/proj"
PROJ_LIB_PATH="${PROJ_PATH}/usr/local/lib"
LD_LIBRARY_PATH="${PROJ_LIB_PATH}"
PROJ_LIB="${PROJ_PATH}/usr/local/share/proj"

PROJ_CMD="${PROJ_PATH}/usr/local/bin/$(basename $0)"

eval ${PROJ_CMD} $*
exit $?

