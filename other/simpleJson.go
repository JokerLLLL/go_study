package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main()  {
	s := `{
	    "data":{
	        "custom_err":{"err_code":300001,"err_msg":"解密错误: 批量解密存在部分失败"},
	        "decrypt_infos":[
	            {"auth_id":"123","cipher_text":"abc","decrypt_text":"","err_msg":"","err_no":0},
	            {"auth_id":"123","cipher_text":"错误密文","decrypt_text":"","err_msg":"KMS解密错误，请联系相对应的开发同学","err_no":300003}
	         ]
	    },
	    "err_no":0,
	    "log_id":"202107231820180101501010795D05A13E",
	    "message":"success"
	}`
	json, _ := simplejson.NewJson([]byte(s))
	jsonAll := json.GetPath("data", "decrypt_infos")
	for i:=range jsonAll.MustArray(){
		newJ := jsonAll.GetIndex(i)
		fmt.Println(newJ)
	}
}
