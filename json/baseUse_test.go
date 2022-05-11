package json

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "test1",
			input: "{\"cluster_uuid\": \"MdRxlC7RQgOjbrxFCaGZMg\",\n\"timestamp\": \"2022-05-10T02:18:49.846Z\",\n\"interval_ms\": 10000,\n\"type\": \"kibana_stats\",\n\"source_node\": {\n\"uuid\": \"2w7n7CJrQ56eCc8b9dPvPw\",\n\"host\": \"192.168.36.61\",\n\"transport_address\": \"192.168.36.61:9300\",\n\"ip\": \"192.168.36.61\",\n\"name\": \"node-192.168.36.61\",\n\"timestamp\": \"2022-05-10T02:18:49.847Z\"\n}}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Parse(tt.input)
		})
	}
}
