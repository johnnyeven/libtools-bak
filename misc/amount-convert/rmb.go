package amount

import (
	"fmt"
	"strconv"
)

// FenToYuan from fen to yuan
func FenToYuan(a int64) string {
	head := a / 100
	tail := a % 100
	return fmt.Sprintf("%d.%02d", head, tail)
}

// YuanToFen from yun to fen
func YuanToFen(a float64) int64 {
	b := fmt.Sprintf("%.0f", a*100.0)
	m, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		panic("convert err")
	}
	return m
}
