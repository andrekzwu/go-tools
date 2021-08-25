package constant

import "time"

type TimeLayout string

const (
	DEFAULT                TimeLayout = time.RFC3339
	FULL_DATETIME          TimeLayout = "2006-01-02 15:04:05"
	SHORT_DATETIME         TimeLayout = "2006-01-02 15:04"
	FULL_TIME              TimeLayout = "15:04:05"
	SHORT_TIME             TimeLayout = "15:04"
	DATE                   TimeLayout = "2006-01-02"
	YEAR_MONTH             TimeLayout = "2006-01"
	YEAR                   TimeLayout = "2006"
	MONTH                  TimeLayout = "01"
	DAY                    TimeLayout = "02"
	COMPACT_FULL_DATETIME  TimeLayout = "20060102150405"
	COMPACT_SHORT_DATETIME TimeLayout = "200601021504"
	COMPACT_FULL_TIME      TimeLayout = "150405"
	COMPACT_SHORT_TIME     TimeLayout = "1504"
	COMPACT_DATE           TimeLayout = "20060102"
	COMPACT_YEAR_MONTH     TimeLayout = "200601"
)
