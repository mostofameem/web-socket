package util

import (
	"testing"
	"time"
)

func TestIsBetweenTimeRange(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name  string
		start time.Time
		end   time.Time
		want  bool
	}{
		{
			name:  "within range",
			start: now.Add(-time.Hour),
			end:   now.Add(time.Hour),
			want:  true,
		},
		{
			name:  "exact start time",
			start: now,
			end:   now.Add(time.Hour),
			want:  true,
		},
		{
			name:  "exact end time",
			start: now.Add(-time.Hour),
			end:   now.Add(time.Millisecond),
			want:  true,
		},
		{
			name:  "before range",
			start: now.Add(time.Hour),
			end:   now.Add(2 * time.Hour),
			want:  false,
		},
		{
			name:  "after range",
			start: now.Add(-2 * time.Hour),
			end:   now.Add(-time.Hour),
			want:  false,
		},
		{
			name:  "start after end (invalid range)",
			start: now.Add(time.Hour),
			end:   now.Add(-time.Hour),
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			current := time.Now().UTC()
			if got := IsBetweenTimeRange(current, tt.start, tt.end); got != tt.want {
				t.Errorf("isBetweenTimeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
