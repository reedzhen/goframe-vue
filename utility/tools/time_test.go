package tools

import (
	"testing"
	"time"
)

func TestTrackTime(t *testing.T) {
	defer TrackTime(time.Now())

	time.Sleep(500 * time.Millisecond)
}

// 输出：
// elapsed: 500.310637ms
