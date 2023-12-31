package utility

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func init() {
	time.Local = time.FixedZone("CST", 8*3600)
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t Time) AsTime() time.Time {
	return time.Time(t)
}

func (t *Time) Scan(v interface{}) error {
	switch v.(type) {
	case time.Time:
		*t = Time(v.(time.Time))
		return nil
	default:
		return errors.New("failed to scan Time")
	}
}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func AsTime(str string) Time {
	t, _ := time.ParseInLocation(timeFormart, str, time.Local)
	return Time(t)
}

func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}
