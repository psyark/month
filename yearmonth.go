package month

import (
	"fmt"
	"time"
)

// YearMonth　は、年月のみを表し、日付以降の解像度を持たないデータ構造です
type YearMonth int

func Get(t time.Time) YearMonth {
	return YearMonth(t.Year()*12 + int(t.Month()))
}

func (ym YearMonth) Year() int {
	return int(ym) / 12
}
func (ym YearMonth) Month() time.Month {
	return time.Month(ym % 12)
}

// ISO 8601
func (ym YearMonth) String() string {
	return fmt.Sprintf("%04d-%02d", ym.Year(), ym.Month())
}
func (ym YearMonth) StartDate() time.Time {
	return time.Date(ym.Year(), ym.Month(), 1, 0, 0, 0, 0, time.UTC)
}
func (ym YearMonth) EndDate() time.Time {
	return ym.StartDate().AddDate(0, 1, -1)
}

func (ym YearMonth) Next() YearMonth {
	return Get(ym.StartDate().AddDate(0, 1, 0))
}
func (ym YearMonth) Prev() YearMonth {
	return Get(ym.StartDate().AddDate(0, -1, 0))
}
