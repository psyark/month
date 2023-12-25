package month

import (
	"fmt"
	"time"
)

func ExampleGet() {
	var ym YearMonth
	ym = Get(time.Date(2023, time.November, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(ym, ym.Year(), ym.Month())
	ym = Get(time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(ym, ym.Year(), ym.Month())
	ym = Get(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(ym, ym.Year(), ym.Month())
	// output:
	// 2023-11 2023 November
	// 2023-12 2023 December
	// 2024-01 2024 January
}
