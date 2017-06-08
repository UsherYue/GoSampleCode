package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Names       []string `json:"NameList"`
	Test        string   `json:",omitempty"`
	Int64String int64    `json:",string"`
}

func main() {
	data := &Data{[]string{"zhangsan", "lisi", "wangwu"}, "1122", 22}
	//序列化结构体
	bts, _ := json.Marshal(&data)
	fmt.Println(string(bts))
	//反序列化json字符串到结构体
	var tmpData Data
	json.Unmarshal(bts, &tmpData)
	fmt.Println(tmpData)
	//解码未知类型的数据
	var jsonData = `{"a":1}`
	var v interface{}
	json.Unmarshal([]byte(jsonData), &v)
	fmt.Println(v)

}
