package atime

import "testing"

func Test_usage(t *testing.T) {
	tests := []struct {
		give string
	}{
		{give: "2022-04-13 10:08:18.448"},
		{give: "2022-06-01 20:20:20.123"},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			usage(tt.give)
		})
	}
}
