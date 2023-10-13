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
