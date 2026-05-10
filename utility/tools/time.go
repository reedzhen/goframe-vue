package tools

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

// TrackTime 计算方法执行时间
func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed:", elapsed)

	return elapsed
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01 00:00:00 或者 2020-01-01
func GetBetweenDates(start, end string) []string {
	var dates []string

	// 解析开始和结束时间
	startTime, err := gtime.StrToTime(start)
	if err != nil {
		return dates
	}
	endTime, err := gtime.StrToTime(end)
	if err != nil {
		return dates
	}

	// 创建一个包含开始时间的日期切片
	dates = append(dates, startTime.Format("Y-m-d"))

	// 如果传入的是同一天，直接返回开始时间
	if startTime.Format("Y-m-d") == endTime.Format("Y-m-d") {
		return dates
	}

	// 逐天增加，直到达到或超过结束时间
	for startTime.Before(endTime) || startTime.Equal(endTime) {
		startTime = startTime.AddDate(0, 0, 1)
		dates = append(dates, startTime.Format("Y-m-d"))
	}

	return dates
}
