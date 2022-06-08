package aencoding

import "testing"

func Test_encodingUsage(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{
			name: "测试1",
			str:  "中国字",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodingUsage(tt.str)
		})
	}
}

func Test_chineseIndex(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chineseIndex()
		})
	}
}
