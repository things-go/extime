package extime

import (
	"fmt"
	"time"
)

// Valid 检查是否正常的日期.
// 1<= year <= 9999
// 1<= month < 12
// 1<= day <=31 && day <= monthDays(month)
func Valid(year, month, day int) bool {
	return ValidYear(year) &&
		ValidMonth(month) &&
		day >= 1 && day <= 31 && day <= MonthDays(year, time.Month(month))
}

// ValidYear 检查是否正常的年.
func ValidYear(year int) bool { return year >= 1 && year <= 9999 }

// ValidMonth 检查是否正常的月.
func ValidMonth(month int) bool { return month >= 1 && month <= 12 }

// IsLeapYear 是否闰年
func IsLeapYear(year int) bool { return (year%4 == 0 && year%100 != 0) || year%400 == 0 }

// Days time.Duration转化为天数
func Days(d time.Duration) float64 {
	return float64(d/(24*time.Hour)) + float64(d%(24*time.Hour))/(24*60*60*1e9)
}

// YearDays 所在年份总天数
func YearDays(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

// MonthDays 所在年份月份的天数
func MonthDays(year int, month time.Month) int {
	switch month {
	case time.January, time.March, time.May, time.July,
		time.August, time.October, time.December:
		return 31
	case time.February:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	case time.April, time.June, time.September, time.November:
		return 30
	default:
		panic(fmt.Errorf("invalid month %v", month))
	}
}

// MonthDays2 t 所在时间月份的天数
func MonthDays2(t time.Time) int { return MonthDays(t.Year(), t.Month()) }

// StartOfDay 获取时间中当天的开始时间.
// 2022-02-18 00:00:00 +0800 CST
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// PreviousStartOfDay 获取时间中上一天的开始时间.
// 2022-02-17 00:00:00 +0800 CST
func PreviousStartOfDay(t time.Time) time.Time {
	return StartOfDay(t).AddDate(0, 0, -1)
}

// NextStartOfDay 获取时间中下一天的开始时间.
// 2022-02-19 00:00:00 +0800 CST
func NextStartOfDay(t time.Time) time.Time {
	return StartOfDay(t).AddDate(0, 0, +1)
}

// EndOfDay 获取时间中当天的结束时间.
// 2022-02-18 23:59:59.999999999 +0800 CST
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// PreviousEndOfDay 获取时间中前一天的结束时间.
// 2022-02-17 23:59:59.999999999 +0800 CST
func PreviousEndOfDay(t time.Time) time.Time {
	return EndOfDay(t).AddDate(0, 0, -1)
}

// NextEndOfDay 获取时间中下一天的结束时间.
// 2022-02-19 23:59:59.999999999 +0800 CST
func NextEndOfDay(t time.Time) time.Time {
	return EndOfDay(t).AddDate(0, 0, 1)
}

// StartOfMonth 获取时间中当月的开始时间.
// 2022-02-01 00:00:00 +0800 CST
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// PreviousStartOfMonth 获取时间中上月的开始时间.
// 2022-01-01 00:00:00 +0800 CST
func PreviousStartOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, -1, 0)
}

// NextStartOfMonth 获取时间中下月的开始时间.
// 2022-03-01 00:00:00 +0800 CST
func NextStartOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0)
}

// EndOfMonth 获取时间中当月的结束时间.
// 2022-02-28 23:59:59.999999999 +0800 CST
func EndOfMonth(t time.Time) time.Time {
	return NextStartOfMonth(t).Add(-time.Nanosecond)
}

// PreviousEndOfMonth 获取时间中上月的结束时间.
// 2022-01-31 23:59:59.999999999 +0800 CST
func PreviousEndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).Add(-time.Nanosecond)
}

// NexEndOfMonth 获取时间中下月的结束时间.
// 2022-03-31 23:59:59.999999999 +0800 CST
func NexEndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 2, 0).Add(-time.Nanosecond)
}

// StartOfYear 获取时间当年的开始时间.
// 2022-01-01 00:00:00 +0800 CST
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// PreviousStartOfYear 获取时间上一年的开始时间.
// 2023-01-01 00:00:00 +0800 CST
func PreviousStartOfYear(t time.Time) time.Time {
	return StartOfYear(t).AddDate(-1, 0, 0)
}

// NextStartOfYear 获取时间下一年的开始时间.
// 2023-01-01 00:00:00 +0800 CST
func NextStartOfYear(t time.Time) time.Time {
	return StartOfYear(t).AddDate(1, 0, 0)
}

// EndOfYear 获取时间当年的结束时间.
// 2022-12-31 23:59:59.999999999 +0800 CST
func EndOfYear(t time.Time) time.Time {
	return NextStartOfYear(t).Add(-time.Nanosecond)
}

// PreviousEndOfYear 获取时间上一年的结束时间.
// 2021-12-31 23:59:59.999999999 +0800 CST
func PreviousEndOfYear(t time.Time) time.Time {
	return StartOfYear(t).Add(-time.Nanosecond)
}

// NextEndOfYear 获取日期中下一年的结束时间.
// 2023-12-31 23:59:59.999999999 +0800 CST
func NextEndOfYear(t time.Time) time.Time {
	return StartOfYear(t).AddDate(2, 0, 0).Add(-time.Nanosecond)
}

// StartOfWeek 获取时间当周的开始时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func StartOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	weekStart := time.Monday
	if len(weekStartDay) > 0 {
		weekStart = weekStartDay[0]
	}

	// 当前是周几
	weekday := int(date.Weekday())
	if weekStart != time.Sunday {
		weekStartDayInt := int(weekStart)

		if weekday < weekStartDayInt {
			weekday += 7 - weekStartDayInt
		} else {
			weekday -= weekStartDayInt
		}
	}

	return time.Date(date.Year(), date.Month(), date.Day()-weekday, 0, 0, 0, 0, date.Location())
}

// PreviousStartOfWeek 获取时间中上一周的开始时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func PreviousStartOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	return StartOfWeek(date, weekStartDay...).AddDate(0, 0, -7)
}

// NextStartOfWeek 获取时间下一周的开始时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func NextStartOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	return StartOfWeek(date, weekStartDay...).AddDate(0, 0, 7)
}

// EndOfWeek 获取时间当周的结束时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func EndOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	return NextStartOfWeek(date, weekStartDay...).Add(-time.Nanosecond)
}

// PreviousEndOfWeek 获取时间上周的结束时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func PreviousEndOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	return StartOfWeek(date, weekStartDay...).Add(-time.Nanosecond)
}

// NextEndOfWeek 获取时间当周的结束时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func NextEndOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	return StartOfWeek(date, weekStartDay...).AddDate(0, 0, 14).Add(-time.Nanosecond)
}

// A Quarter specifies a quarter of the year (1, 2 , 3 ,4).
type Quarter int

const (
	Quarter1th Quarter = iota + 1 // 1, 2 ,3
	Quarter2th                    // 4, 5, 6
	Quarter3th                    // 7, 8, 9
	Quarter4th                    // 10, 11, 12
)

// ValidQuarter 是否为有效的季度
func ValidQuarter(q int) bool { return q >= 1 && q <= 4 }

// GetQuarter 获取季度
func GetQuarter(month time.Month) Quarter {
	switch m := int(month); {
	case m >= 1 && m <= 3:
		return 1
	case m >= 4 && m <= 6:
		return 2
	case m >= 7 && m <= 9:
		return 3
	case m >= 10 && m <= 12:
		return 4
	}
	return 0
}

// GetQuarter2 通过时间获取季度
func GetQuarter2(t time.Time) Quarter { return GetQuarter(t.Month()) }

// StartOfQuarter 获取季度开始的时间
// q的值不为范围内的值, 返回第一季度值
// 2022-01-01 00:00:00 +0800 CST
func StartOfQuarter(year int, q Quarter) time.Time {
	switch q {
	case Quarter1th:
		return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	case Quarter2th:
		return time.Date(year, 4, 1, 0, 0, 0, 0, time.Local)
	case Quarter3th:
		return time.Date(year, 7, 1, 0, 0, 0, 0, time.Local)
	case Quarter4th:
		return time.Date(year, 10, 1, 0, 0, 0, 0, time.Local)
	}
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
}

// PreviousStartOfQuarter 获取上个季度开始的时间
// 2022-01-01 00:00:00 +0800 CST
func PreviousStartOfQuarter(year int, q Quarter) time.Time {
	return StartOfQuarter(year, q).AddDate(0, -3, 0)
}

// NextStartOfQuarter 获取下个季度开始的时间
// 2022-04-01 00:00:00 +0800 CST
func NextStartOfQuarter(year int, q Quarter) time.Time {
	return StartOfQuarter(year, q).AddDate(0, 3, 0)
}

// EndOfQuarter 获取季度结束的时间
// 2022-03-31 23:59:59.999999999 +0800 CST
func EndOfQuarter(year int, q Quarter) time.Time {
	return NextStartOfQuarter(year, q).Add(-time.Nanosecond)
}

// PreviousEndOfQuarter 获取上一季度结束的时间
// 2022-12-31 23:59:59.999999999 +0800 CST
func PreviousEndOfQuarter(year int, q Quarter) time.Time {
	return StartOfQuarter(year, q).Add(-time.Nanosecond)
}

// NextEndOfQuarter 获取下一季度结束的时间
// 2022-06-31 23:59:59.999999999 +0800 CST
func NextEndOfQuarter(year int, q Quarter) time.Time {
	return StartOfQuarter(year, q).AddDate(0, 6, 0).Add(-time.Nanosecond)
}

// StartOfQuarter2 获取时间季度开始的时间
// 2022-01-01 00:00:00 +0800 CST
func StartOfQuarter2(t time.Time) time.Time {
	return StartOfQuarter(t.Year(), GetQuarter(t.Month()))
}

// PreviousStartOfQuarter 获取时间的上个季度开始的时间
// 2022-01-01 00:00:00 +0800 CST
func PreviousStartOfQuarter2(t time.Time) time.Time {
	return PreviousStartOfQuarter(t.Year(), GetQuarter(t.Month()))
}

// NextStartOfQuarter 获取时间的下个季度开始的时间
// 2022-04-01 00:00:00 +0800 CST
func NextStartOfQuarter2(t time.Time) time.Time {
	return NextStartOfQuarter(t.Year(), GetQuarter(t.Month()))
}

// EndOfQuarter 获取时间的季度结束的时间
// 2022-03-31 23:59:59.999999999 +0800 CST
func EndOfQuarter2(t time.Time) time.Time {
	return EndOfQuarter(t.Year(), GetQuarter(t.Month()))
}

// PreviousEndOfQuarter 获取时间的上一季度结束的时间
// 2022-12-31 23:59:59.999999999 +0800 CST
func PreviousEndOfQuarter2(t time.Time) time.Time {
	return PreviousEndOfQuarter(t.Year(), GetQuarter(t.Month()))
}

// NextEndOfQuarter 获取时间的下一季度结束的时间
// 2022-06-31 23:59:59.999999999 +0800 CST
func NextEndOfQuarter2(t time.Time) time.Time {
	return NextEndOfQuarter(t.Year(), GetQuarter(t.Month()))
}
