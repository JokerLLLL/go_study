package main

import (
	"encoding/xml"
	"fmt"
)

/*
```xml
<sites>
    <site>
        <name>xx</name>
        <url>yy</url>
    </site>
    <site>
        <se>111</se>
        <se>222</se>
		<se2>test2</se2>
		<se2>test2</se2>
    </site>
</sites>
```
翻译成golang的结构体
 */

// 定义结构体来表示 XML 结构
type Sites struct {
	XMLName xml.Name `xml:"sites"`
	Sites   []Site   `xml:"site"`
}

type Site struct {
	Name string   `xml:"name,omitempty"`
	URL  string   `xml:"url,omitempty"`
	Se   []string `xml:"se,omitempty"`
	Se2  []string `xml:"se2,omitempty"`
}

// - `omitempty` 标签表示如果字段为空，则在生成 XML 时忽略该字段。

func main() {
	// 示例 XML 数据
	data := `
<sites>
    <site>
        <name>xx</name>
        <url>yy</url>
    </site>
    <site>
        <se>111</se>
        <se>222</se>
        <se2>test2</se2>
        <se2>test2</se2>
    </site>
</sites> 
`

	// 创建一个 Sites 结构体实例
	var sites Sites

	// 解析 XML 数据
	err := xml.Unmarshal([]byte(data), &sites)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 输出解析结果
	fmt.Printf("Parsed XML: %+v\n", sites)
	var re = Sites{}
	var ss = []Site{}
	ss = append(ss, Site{Name: "abc", URL: "abc.com"})
	ss = append(ss, Site{Se: []string{"ee1", "ee1"}, Se2: []string{"ee1", "ee2"}})
	ss = append(ss, Site{Se2: []string{"ee2"}})
	re.Sites = ss
	var sXml, _ = xml.Marshal(re)
	fmt.Println(string(sXml))

	var x = Sites{
		Sites: []Site{
			{Name: "123", URL: "123.com"},
			{Se: []string{"ee1", "ee2"}, Se2: []string{"ee3", "ee4"}},
		},
	}
	s, _ := xml.Marshal(x)
	fmt.Println(string(s))
}