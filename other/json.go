package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type ParamJson struct {
	param_json map[string]interface{}
}

func main() {

	// 例如: param_json={"product_id":123,"code":"HHK"} 需要调整为param_json={"code":"HHK","product_id":"123"}

	map1 := make(map[string]interface{})
	map2 := make(map[string]interface{})


	map1["product_id"] = 123
	map1["code"] = "HH<a></a>K"
	map1["a"] = "中国"
	map1["dd"] = "<span></span>####aaa\n\\\\k dd"

	map2["code"] = "HH<a></a>K"
	map2["product_id"] = 123
	map2["a"] = "aaaaaa"


	for k,v:=range map1 {
		fmt.Printf("key:%s,value:%v \n", k, v)
	}
	fmt.Printf("===============\n")
	for k,v:=range map2 {
		fmt.Printf("key:%s,value:%v \n", k, v)
	}
	bytes1, _ := json.Marshal(map1)
	bytes2, _ := json.Marshal(map2)

	fmt.Printf("json map1:%v \n", string(bytes1))
	fmt.Printf("json map2:%v \n", string(bytes2))


	// https://blog.csdn.net/wangzhezhilu001/article/details/91418069

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	//jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(map1)
	bf.String()
	fmt.Println("escapeHtml:", bf.String())

	// https://www.jianshu.com/p/2f7bdfe6f73d
	//第一种方法，替换'<', '>', '&'
	content := strings.Replace(string(string(bytes1)), "\\u003c", "<", -1)
	content = strings.Replace(content, "\\u003e", ">", -1)
	content = strings.Replace(content, "\\u0026", "&", -1)
	fmt.Println("escapeHtml2:", content)

}
