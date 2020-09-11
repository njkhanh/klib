package times

import (
	"time"
)

const DayDuration time.Duration = time.Duration(24) * time.Hour

type RangeTime struct {
	Start time.Time
	End   time.Time
}

func GetRoundedTime(t time.Time, hour int, min int, sec int) time.Time {

	rounded := time.Date(t.Year(), t.Month(), t.Day(), hour, min, sec, 0, t.Location())

	return rounded
}

func IsToDay(ct int, dateFormat string) bool {
	currentTime := time.Now().Format(dateFormat)

	t, err := time.ParseInLocation(dateFormat, currentTime, time.Local)
	fromDate := int(t.Unix())
	return err == nil && fromDate <= ct && ct < fromDate+86400
}

func GetWeekDays(start time.Time, end time.Time) []RangeTime {
	currentWeekday := int(start.Weekday())
	weekDays := make([]RangeTime, 0, 6)
	if currentWeekday > 0 {
		for i := 1; i <= 6; i++ {
			rs := RangeTime{}
			if i < currentWeekday {
				rs.Start = start.Add(-DayDuration * time.Duration(currentWeekday-i))
				rs.End = end.Add(-DayDuration * time.Duration(currentWeekday-i))
			} else if i > currentWeekday {
				rs.Start = start.Add(DayDuration * time.Duration(i-currentWeekday))
				rs.End = end.Add(DayDuration * time.Duration(i-currentWeekday))
			} else {
				rs.Start = start
				rs.End = end
			}

			weekDays = append(weekDays, rs)

		}

	}

	return weekDays

}
