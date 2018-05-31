package duration_test

import (
	"fmt"
	"testing"
	"time"

	"golib/tools/duration"
)

func TestNewCost(t *testing.T) {
	cost := duration.NewDuration()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(cost.Get())
}
