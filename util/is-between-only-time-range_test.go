package util

import (
	"testing"
	"time"
)

func TestIsBetweenOnlyTimeRange(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected bool
	}{
		{
			name:     "within range",
			start:    convertToTimeOnly(now.Add(-time.Hour)),
			end:      convertToTimeOnly(now.Add(time.Hour)),
			expected: true,
		},
		{
			name:     "exact start time",
			start:    convertToTimeOnly(now),
			end:      convertToTimeOnly(now.Add(time.Hour)),
			expected: true,
		},
		{
			name:     "exact end time",
			start:    convertToTimeOnly(now.Add(-time.Hour)),
			end:      convertToTimeOnly(now.Add(time.Millisecond)),
			expected: true,
		},
		{
			name:     "before range",
			start:    convertToTimeOnly(now.Add(time.Hour)),
			end:      convertToTimeOnly(now.Add(2 * time.Hour)),
			expected: false,
		},
		{
			name:     "after range",
			start:    convertToTimeOnly(now.Add(-2 * time.Hour)),
			end:      convertToTimeOnly(now.Add(-time.Hour)),
			expected: false,
		},
		{
			name:     "start after end (invalid range)",
			start:    convertToTimeOnly(now.Add(time.Hour)),
			end:      convertToTimeOnly(now.Add(-time.Hour)),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			current := time.Now().UTC()
			if got := IsBetweenOnlyTimeRange(current, tt.start, tt.end); got != tt.expected {
				t.Errorf("isBetweenOnlyTimeRange() = %v, want %v", got, tt.expected)
			}
		})
	}
}
