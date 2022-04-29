package json

import (
	"fmt"
	"github.com/json-iterator/go"
	"reflect"
)

func Parse() {
	str := "{\"type\": 7,\"info\":{\"trace_id\":\"new_jjjjjjjjjjjjjjjjjjjjjjjj\",\"span_id\":\"0\",\"depot_id\":\"\",\"module_type\":\"\",\"timestamp\":\"2022-04-13 10:08:18.448\",\"name\":\"ccccccccccccccc\",\"ip\":\"172.16.127.5:60050\",\"response_status\":0,\"version\":\"\",\"span_type\":1,\"server_name\":\"act.getFundaccountExt\",\"app_type\":\"jres-svr\",\"encode\":1,\"response_time\":30000.0,\"micro_service\":true,\"dlog_flag_chg\":0},\"mypackage\":{\"req\":\"\",\"rep\":\"Nzc9MQAxMT01NzQ1MjA0ADMzPXt9ADEzPTgAMz0xADE0PTE2ADM3PTIANT0xMTE0MDUANzU9cHVibGljIGFic3RyYWN0IGNvbS5odW5kc3VuLmhzZnVuZC50cmFkZS5iYXNlcHViLnByb3h5LmJhc2VwdWIuYmFzZXB1YmZ1bmN0aW9uLmR0by5PcGVyYXRvckxvZ1BhZ2VSZXNwb25zZSBjb20uaHVuZHN1bi5oc2Z1bmQudHJhZGUuYmFzZXB1Yi5wcm94eS5iYXNlcHViLmJhc2VwdWJmdW5jdGlvbi5zZXJ2aWNlLk9wZXJhdG9yQ2VudGVyU2VydmljZS5nZXRMb2dMaXN0QnlQYWdlKGNvbS5odW5kc3VuLmhzZnVuZC50cmFkZS5iYXNlcHViLnByb3h5LmJhc2VwdWIuYmFzZXB1YmZ1bmN0aW9uLmR0by5PcGVyYXRvckxvZ1BhZ2VSZXF1ZXN0KQA0Mz0xADI9NgAzNT1kOWJkNGE5MGYwZjE0MjMxYjE5OGQ3Y2MyYmNkY2JlNDJmNzcyMjBhNjA3YTQ0MDAAYTAAMAAAAAAxPTg2ADg9AA4hAgMAAAAAAwAAAAEAAABGAAAAAHRvdGFsAEkCAAAAZIwBA19jb3VudMQCAAVyb3dzAFICAAAAZDAAMAAAAAAUIQEDAC8AAAN0eXBlRHMvWAARAAAA\",\"error_no\":\"\",\"error_info\":\"\",\"type\":1,\"rep_clean\":0},\"refs\":[{\"req\":\"\",\"rep\":\"\",\"time\":\"10000\",\"type\":\"2\",\"span_id\":\"0.1\",\"span_type\":\"1\",\"dst_name\":\"duiduiduid\",\"charset\":1,\"index\":1}]}"
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data := new(map[string]interface{})
	// 方式一： 这种解析是解析全部json string,其中这里的data是map，也可以是一个具体的struct结构休
	err := json.UnmarshalFromString(str, &data)
	if err == nil {
		fmt.Printf("%v\n", data)
	}

	// 方式二: 这种就只解析这个字段 mypackage.rep，适用于 “定向解析”
	get := json.Get([]byte(str), "mypackage", "rep")
	b := get.GetInterface()
	if b != nil {
		fmt.Println(get.ToString())
	}
	fmt.Println(reflect.TypeOf(get))
}
