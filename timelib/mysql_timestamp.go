package timelib

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"fmt"
	"strconv"
	"time"
)

var (
	MySQLTimestampZero     = MySQLTimestamp(time.Time{})
	MySQLTimestampUnixZero = MySQLTimestamp(time.Unix(0, 0))
)

// swagger:strfmt date-time
type MySQLTimestamp time.Time

func ParseMySQLTimestampFromString(s string) (dt MySQLTimestamp, err error) {
	var t time.Time
	t, err = time.Parse(time.RFC3339, s)
	dt = MySQLTimestamp(t)
	return
}

var _ interface {
	sql.Scanner
	driver.Valuer
} = (*MySQLTimestamp)(nil)

func (dt *MySQLTimestamp) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		n, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return fmt.Errorf("sql.Scan() strfmt.MySQLTimestamp from: %#v failed: %s", v, err.Error())
		}
		*dt = MySQLTimestamp(time.Unix(n, 0))
	case int64:
	case float64:
		if v < 0 {
			*dt = MySQLTimestamp{}
		} else {
			*dt = MySQLTimestamp(time.Unix(int64(v), 0))
		}
	case nil:
		*dt = MySQLTimestampZero
	default:
		return fmt.Errorf("cannot sql.Scan() strfmt.MySQLTimestamp from: %#v", v)
	}
	return nil
}

func (dt MySQLTimestamp) Value() (driver.Value, error) {
	return (time.Time)(dt).Unix(), nil
}

func (dt MySQLTimestamp) String() string {
	return time.Time(dt).In(CST).Format(time.RFC3339)
}

func (dt MySQLTimestamp) Format(layout string) string {
	return time.Time(dt).In(CST).Format(layout)
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*MySQLTimestamp)(nil)

func (dt MySQLTimestamp) MarshalText() ([]byte, error) {
	if dt.IsZero() {
		return []byte(""), nil
	}
	str := dt.String()
	return []byte(str), nil
}

func (dt *MySQLTimestamp) UnmarshalText(data []byte) (err error) {
	str := string(data)
	if len(str) > 1 {
		if str[0] == '"' && str[len(str)-1] == '"' {
			str = str[1 : len(str)-1]
		}
	}
	if len(str) == 0 || str == "0" {
		str = MySQLDatetimeZero.String()
	}
	*dt, err = ParseMySQLTimestampFromString(str)
	return
}

func (dt MySQLTimestamp) Unix() int64 {
	return time.Time(dt).Unix()
}

func (dt MySQLTimestamp) IsZero() bool {
	unix := dt.Unix()
	return unix == 0 || unix == MySQLTimestampZero.Unix()
}

func (dt MySQLTimestamp) In(loc *time.Location) MySQLTimestamp {
	return MySQLTimestamp(time.Time(dt).In(loc))
}

// 获取当天最后一秒（东8区）
func (dt MySQLTimestamp) GetTodayLastSecCST() MySQLTimestamp {
	return MySQLTimestamp(GetTodayLastSecInLocation(time.Time(dt), CST))
}

// 添加 N 个工作日（东8区）
func (dt MySQLTimestamp) AddWorkingDaysCST(days int) MySQLTimestamp {
	return MySQLTimestamp(AddWorkingDaysInLocation(time.Time(dt), days, CST))
}

// 获取当天0点（东8区）
func (dt MySQLTimestamp) GetTodayFirstSecCST() MySQLTimestamp {
	return MySQLTimestamp(GetTodayFirstSecInLocation(time.Time(dt), CST))
}
