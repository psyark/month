package month

import (
	"fmt"
	"time"
)

// Period は、年月の区間を表します
type Period struct {
	Start YearMonth
	End   YearMonth
}

func (p Period) ContainsTime(date time.Time) bool {
	return p.Contains(Get(date))
}
func (p Period) Contains(ym YearMonth) bool {
	return ym >= p.Start && ym <= p.End
}

func (p Period) StartDate() time.Time {
	return p.Start.StartDate()
}
func (p Period) EndDate() time.Time {
	return p.End.EndDate()
}

func (p Period) Length() int {
	return int(p.End) - int(p.Start) + 1
}

// ISO 8601
func (p Period) String() string {
	return fmt.Sprintf("%s/%s", p.Start, p.End)
}
