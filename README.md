# gonote

Library and CLI tool for parsing Go project documentation notes

## Overview

`gonote` provides the ability to recursively parse a Go project space for godoc notes. These notes [are described in the godoc library documentation](https://golang.org/pkg/go/doc/#Note).

Godoc notes are structured comments found in Go source code. These notes take the form of a marker, an UID, and a body:

`MARKER(uid): note body`

The marker labels the category of comment. For example, you may have a comment related to describing a known bug, in which case the marker would be `BUG`. You could also have a marker for labeling a TODO item, where the marker would be `TODO`. In both cases, the UID could be a username responsible for following up on that issue. The note body contains any related details.

### Examples

Here are a few examples of godoc notes from the Go standard library:

- [`TODO(mundaym): merge with case 24.`](https://github.com/golang/go/blob/9fa988547a778540eebfe0358536b7433efe6748/src/cmd/internal/obj/s390x/asmz.go#L3083)
- [`TODO(rogpeppe): re-enable this test in https://go-review.googlesource.com/#/c/5910/`](https://github.com/golang/go/blob/38c561cb2caa8019e44059e2be71c909ceef30a6/src/encoding/xml/marshal_test.go#L1768)
- [`TODO(rsc): Decide which tests are enabled by default.`](https://github.com/golang/go/blob/834d2244a0150d8ae29b587ed2193e81e552d601/src/cmd/go/internal/test/test.go#L505)
- [`BUG(brainman): This package is not implemented on Windows. As the syslog package is frozen, Windows users are encouraged to use a package outside of the standard library. For background, see https://golang.org/issue/1108.`](https://github.com/golang/go/blob/f9027d61ab48154e4cb29c50e356a3f462840e01/src/log/syslog/doc.go#L19)
- [`BUG(akumar): This package is not implemented on Plan 9.`](https://github.com/golang/go/blob/f9027d61ab48154e4cb29c50e356a3f462840e01/src/log/syslog/doc.go#L24)
- [`BUG(rsc): FieldByName and related functions consider struct field names to be equal if the names are equal, even if they are unexported names originating in different packages.`](https://github.com/golang/go/blob/997d7a1893ae15df1438c46487dd69903f16c57f/src/reflect/type.go#L211)

## Installation

`gonote` requires Go v1.10+ and a valid Go workspace.

To install:

`go get -u github.com/pokstad/gonote/gonote`

## Usage

```
  OVERVIEW:

    gonote scans Go source code files for notations

  USAGE:

    gonote [OPTIONS] TARGET

  TARGET:

    the Go package path you wish to scan

  OPTIONS:

    -output [json|text|textmate]
```