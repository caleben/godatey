package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	message := "{ \"@timestamp\": \"18/May/2021:08:14:03 +0800\", \"remote_addr\": \"10.20.146.127\", \"remote_user\": \"-\", \"body_bytes_sent\": \"437\", \"request_time\": \"0.000\", \"request_length\": \"171\", \"request\": \"GET /manage/get_hsiar_list HTTP/1.1\", \"request_method\": \"GET\", \"http_content_type\": \"-\", \"upstream_addr\" : \"-\" ,\"upstream_response_time\": \"-\", \"status\": \"200\", \"client_id\": \"-\", \"http_referrer\": \"-\", \"http_x_forwarded_for\": \"-\", \"tid\": \"-\", \"http_user_agent\": \"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1;SV1)\" }"
	m := make(map[string]interface{})
	of := reflect.TypeOf(message)
	fmt.Printf("%v\n", of)
	err := json.Unmarshal([]byte(message), &m)
	if err != nil {
		fmt.Print(err)
	} else {
		for k, v := range m {
			fmt.Printf("%v    %v\n", k, v)
		}
	}
}
