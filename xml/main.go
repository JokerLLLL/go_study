package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type SaleOrderStatusPushServiceResponse struct {
	XMLName xml.Name `xml:"Response"`  //  XMLName 区分大小写
	Head    string   `xml:"Head"`
	Error   string   `xml:"Error"`
}



type SaleOrderOutboundDetailPushServiceResponse struct {
	XmlName xml.Name   `xml:"Response"`
	Head    string     `xml:"Head"`
	Error   string     `xml:"Error"`
	Body    DetailBody `xml:"Body"`
}

type DetailBody struct {
	SaleOrderOutboundDetailResponse SaleOrderOutboundDetailResponse `xml:"SaleOrderOutboundDetailResponse"`
}

type SaleOrderOutboundDetailResponse struct {
	Result int64  `xml:"Result"`
	Note   string `xml:"Note"`
}

func main() {
	// 创建结构体实例并填充数据
	response := SaleOrderStatusPushServiceResponse{
		Head:  "OK",
		Error: "None",
	}

	// 设置 XML 编码器
	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "")

	xmlStr, _ := xml.Marshal(response)

	fmt.Println(string(xmlStr))

	// 输出 XML 头部
	fmt.Println(`<?xml version="1.0" encoding="UTF-8"?>`)

	// 编码结构体为 XML 并输出到标准输出
	if err := encoder.Encode(response); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("===================================>>>>>>>>>>>>")

	x := SaleOrderOutboundDetailPushServiceResponse{
		Head:  "OKK",
		Error: "IS_ERR",
	}
	xmlStr2, _ := xml.Marshal(x)
	fmt.Println(string(xmlStr2))
}