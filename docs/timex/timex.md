# Timex

Package timex supports date and time format and compare.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/zzopen/zgo/blob/main/timex/datetime.go](https://github.com/zzopen/zgo/blob/main/timex/datetime.go)


<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/zzopen/zgo"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [FormatTimeToStr](#FormatTimeToStr)


<div STYLE="page-break-after: always;"></div>

## Documentation

## Note:

1. 'format' string param in func FormatTimeToStr and FormatStrToTime function should be one of flows:

- YYYY-MM-DD HH:mm:ss
- YYYY-MM-DD HH:mm
- YYYY-MM-DD HH
- YYYY-MM-DD
- HH:mm:ss
- HH:mm
- HH


### <span id="Equal">FormatTimeToStr</span>

<p>Format time to string, `format` param specification see note 1.</p>

<b>Signature:</b>

```go
func FormatTimeToStr(t time.Time, format string) string
```

<b>Example:</b>

```go
package main

import (
	"fmt"
	"time"

	"github.com/zzopen/zgo/timex"
)

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2023-07-31 16:04:08")

	result1 := timex.FormatTimeToStr(t, "YYYY-MM-DD HH:mm:ss")
	result2 := timex.FormatTimeToStr(t, "YYYY-MM-DD")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 2023-07-31 16:04:08
	// 2023-07-31

}
```