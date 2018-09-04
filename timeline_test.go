package timeline

import (
	"testing"
	"time"
)

func TestTimeline(t *testing.T) {
	line := New()

	// Add all your callbacks
	line.Append(func() {
		// Profit here
	}, time.Now().Add(10*time.Second))

	line.Start()
}
