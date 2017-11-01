package clock_test

import (
	"testing"

	"fmt"
	"time"

	"github.com/goph/clock"
	"github.com/stretchr/testify/assert"
)

func TestSystemClock(t *testing.T) {
	ti := time.Now()

	time.Sleep(time.Nanosecond)

	ti = ti.Add(time.Second)
	assert.True(t, clock.SystemClock.Now().Before(ti), "expected the clock's current time to be before %v", ti)
}

func TestStoppedClock(t *testing.T) {
	ti := time.Date(2017, time.May, 10, 22, 52, 0, 0, time.UTC)

	c := clock.StoppedAt(ti)

	assert.Equal(t, ti, c.Now())
}

func ExampleStoppedAt() {
	t := time.Date(2017, time.May, 10, 22, 52, 0, 0, time.UTC)
	c := clock.StoppedAt(t)

	fmt.Println(c.Now())
	// Output: 2017-05-10 22:52:00 +0000 UTC
}
