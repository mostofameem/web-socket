package util

import (
	"testing"
)

func TestPaginateSlice(t *testing.T) {
	tests := []struct {
		name     string
		sliceLen int
		page     int
		pageSize int
		start    int
		end      int
	}{
		{
			name:     "Should Success",
			sliceLen: 100,
			page:     2,
			pageSize: 50,
			start:    50,
			end:      100,
		},
		{
			name:     "Should Success",
			sliceLen: 0,
			page:     2,
			pageSize: 50,
			start:    0,
			end:      0,
		},
		{
			name:     "Should Success",
			sliceLen: 7,
			page:     2,
			pageSize: 5,
			start:    5,
			end:      7,
		},
		{
			name:     "Should Success",
			sliceLen: 7,
			page:     0,
			pageSize: 0,
			start:    0,
			end:      7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start, end := PaginateSlice(test.sliceLen, test.page, test.pageSize)

			if start != test.start {
				t.Errorf("Expected start: %d. But got %d\n", test.start, start)
			}

			if end != test.end {
				t.Errorf("Expected end: %d. But got %d\n", test.end, end)
			}
		})
	}
}
