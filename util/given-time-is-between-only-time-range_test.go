package util

import (
	"testing"
	"time"
)

func TestGivenTimeIsBetweenOnlyTimeRange(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name      string
		givenTime time.Time
		start     time.Time
		end       time.Time
		want      bool
	}{
		{
			name:      "within range",
			givenTime: now.Add(30 * -time.Minute),
			start:     now.Add(-time.Hour),
			end:       now.Add(time.Hour),
			want:      true,
		},
		{
			name:      "Start time greater than given time",
			givenTime: now.Add(30 * -time.Minute),
			start:     now,
			end:       now.Add(time.Millisecond),
			want:      false,
		},
		{
			name:      "End time greater than given time",
			givenTime: now.Add(30 * time.Minute),
			start:     now,
			end:       now.Add(time.Millisecond),
			want:      false,
		},
		{
			name:      "given time and start time exact",
			givenTime: now,
			start:     now,
			end:       now.Add(time.Millisecond),
			want:      true,
		},
		{
			name:      "given time and end time exact",
			givenTime: now,
			start:     now.Add(-time.Hour),
			end:       now,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GivenTimeIsBetweenOnlyTimeRange(tt.givenTime, tt.start, tt.end); got != tt.want {
				t.Errorf("GiventimeIsBetweenOnlyTimeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
