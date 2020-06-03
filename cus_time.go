package cus

import (
	"time"
)

//GetUTCTime 取UTC目前時間
func GetUTCTime() time.Time {
	t := time.Now()
	timeUTC := t.In(time.UTC)
	return timeUTC
}

//TimeToString Time轉String
func TimeToString(value time.Time, formatConfig string) string {
	result := value.Format(formatConfig)
	return result
}

//StringToTime String轉Time
func StringToTime(value string, formatConfig string) (time.Time, error) {
	result, err := time.Parse(formatConfig, value)
	return result, err
}

//UTCStringToTime UTC String轉Time
func UTCStringToTime(value string, formatConfig string) (time.Time, error) {
	result, err := time.Parse(formatConfig, value)
	return result, err
}
