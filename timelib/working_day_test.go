package timelib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTodayFirstSecInLocation(t *testing.T) {
	tt := assert.New(t)
	t0, _ := time.Parse(time.RFC3339, "2017-03-27T23:58:59+08:00")
	mtl0 := GetTodayFirstSecInLocation(t0, CST)
	tt.Equal("2017-03-27T00:00:00+08:00", mtl0.Format(time.RFC3339))
	tt.Equal(int64(1490544000), time.Time(mtl0).Unix())
}

func TestGetTodayLastSecInLocation(t *testing.T) {
	tt := assert.New(t)
	t0, _ := time.Parse(time.RFC3339, "2017-03-27T23:58:59+08:00")
	mtl0 := GetTodayLastSecInLocation(t0, CST)
	tt.Equal("2017-03-27T23:59:59+08:00", mtl0.Format(time.RFC3339))
	tt.Equal(int64(1490630399), time.Time(mtl0).Unix())
}

func TestAddWorkingDaysInLocation(t *testing.T) {
	tt := assert.New(t)
	t0, _ := time.Parse(time.RFC3339, "2017-03-27T23:58:59+08:00")
	mtl0 := AddWorkingDaysInLocation(t0, 10, CST)
	tt.Equal("2017-04-10T23:58:59+08:00", mtl0.Format(time.RFC3339))
	tt.Equal(int64(1491839939), time.Time(mtl0).Unix())
}

func TestTime(t *testing.T) {
	//tt := assert.New(t)
	fomat := "20060102150405"
	t0, _ := time.ParseInLocation(fomat, "20160312104325", CST)
	datetime := MySQLDatetime(t0)
	t.Logf("====%s\n", datetime.String())
}

func TestTime2(t *testing.T) {
	//tt := assert.New(t)
	fomat := "20060102150405"
	datetime, _ := ParseMySQLDatetimeFromStringWithCustomFormatterInCST("20160312104325", fomat)
	t.Logf("====%s\n", datetime.String())
}
