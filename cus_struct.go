package cus

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

//JSONToArrayObject JSON轉ArrayObject
func JSONToArrayObject(payload string) (result []interface{}) {
	json.Unmarshal([]byte(payload), &result)
	return
}

//MapToJSON Map轉JSON
func MapToJSON(payload map[string]interface{}) (result string) {
	jsonByte, _ := json.Marshal(payload)
	result = string(jsonByte)

	return
}

//JSONToMap JSON轉Map
func JSONToMap(payload string) (result map[string]interface{}) {
	result = make(map[string]interface{})
	json.Unmarshal([]byte(payload), &result)

	return
}

//StructToJSON Struct轉JSON
func StructToJSON(value interface{}) (result string) {
	jsonStr, _ := json.Marshal(value)
	result = string(jsonStr)

	return
}

//JSONToStruct JSON轉Struct
func JSONToStruct(value string, result interface{}) {
	jsonMap := JSONToMap(value)
	mapstructure.Decode(jsonMap, &result)

	return
}

//StringToJSON String轉JSON
func StringToJSON(payload string) (result string) {
	json.Unmarshal([]byte(payload), &result)

	return
}

//JSONBToString JSONB([]Uint8)轉String
func JSONBToString(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

//JSONBToStruct JSONB轉Struct
func JSONBToStruct(value []uint8, result interface{}) {
	jsonMap := JSONToMap(JSONBToString(value))
	mapstructure.Decode(jsonMap, &result)

	return
}

//StructToMap Struct轉Map String
func StructToMap(value interface{}) (result map[string]interface{}) {
	jsonStr := StructToJSON(value)
	result = JSONToMap(jsonStr)

	return
}

//MapToStruct Map轉Struct
func MapToStruct(value interface{}, result interface{}) {
	mapstructure.Decode(value, &result)

	return
}

//BooleanToInteger Boolean 轉 Integer
func BooleanToInteger(b bool) int {
	if b {
		return 1
	}
	return 0
}
