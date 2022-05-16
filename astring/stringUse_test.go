package astring

import "testing"

func TestStringUsage(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试string用法"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringUsage()
		})
	}
}
