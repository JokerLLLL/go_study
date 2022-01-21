package main

import (
	"encoding/json"
	"fmt"
)

func interfaceToString(src interface{}) string {
	if src == nil {
		return ""
	}
	switch src.(type) {
	case string:
		fmt.Println("get string")
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return fmt.Sprint(src)
	}
	data, err := json.Marshal(src)
	fmt.Println("get struct")

	if err != nil {
		panic(err)
	}
	return string(data)
}

type StructString string

func main() {
	var s string = "test"
	x := StructString(s)
	fmt.Println(x,interfaceToString(x))

}
