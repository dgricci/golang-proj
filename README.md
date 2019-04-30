% GOPROJ : A go wrapper for PROJ the coordinate transformation software
% Didier Richard
% 2019/03/10

---

Revision:
    * 2019/03/10 : first release

---

# Rational #

To provide a wrapper in Golang for [PROJ](https://proj4.org/index.html) and
learn `cgo` mechanism :smiley:

## Compiling PROJ 6.0.0 ##

In the `$GOPATH/src/osgeo.org/proj/usr/local/src/proj-x.y.z` directory :

```bash
export PROJ_LIB=$GOPATH/src/osgeo.org/proj/usr/local/share/proj
./configure --prefix=$GOPATH/src/osgeo.org/proj/usr/local --disable-static
make -j$(nproc)
make install
make check
make clean
unset PROJ_LIB
```

## Building the project ##

* To generate the Gopkg.lock file (first time) :

```bash
make dep-install
```

* To test the wrapper :

```bash
make test
```

* To install the wrapper :

```
make install
```

## Miscellaneous

If the PROJ library and the PROJ files are not installed in usual places
(like `/usr/lib` or `/usr/lib/x86_64-linux-gnu/` for the former and
`/usr/share.proj` for the latter), don't forget upon wrapper's installation to
have the `LD_LIBRARY_PATH` set to point at the PROJ library and `PROJ_LIB` to
the directory where `proj.db` is.

