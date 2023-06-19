package util

import "time"

var layout = "2006-01-02"

func ConvertToDate(str string) (time.Time, error) {
	return time.Parse(layout, str)
}
