# go-saber
A comprehensive, efficient, and reusable util function library of Go.

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.20-9cf)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zzopen/go-saber/blob/main/LICENSE)

<div STYLE="page-break-after: always;"></div>

## Feature

-   👏 Comprehensive, efficient and reusable.
-   💪 more go util functions, support string, datetime, net ...
-   💅 Only depends on two kinds of libraries: go standard library and golang.org/x.
-   🌍 Unit test for every exported function.

## Installation

### Note:
```go
go get github.com/zzopen/go-saber
```

## Usage

saber organizes the code into package structure, and you need to import the corresponding package name when use it. For example, if you use string-related functions,import the strutil package like below:

```go
import "github.com/zzopen/go-saber/timex"
```

## Example


```go
package main

import (
    "fmt"
    "github.com/zzopen/go-saber/timex"
)

func main() {
    rs := timex.NowDateTimeStr()
    fmt.Println(rs) // 2023-07-31 12:12:12
}
```

## Documentation

### <span id="index">Index<span>

- [Time](#Timex)

