package main

import "fmt"

func TestQueryType(prm interface{}) {
	switch prm.(type) {
	case string:
		{
			fmt.Println("string type")
		}
	case int:
		{
			fmt.Println("int type")
		}
	case map[string]string:
		{
			fmt.Println("map[string]string type")
		}
	default:
		{
			fmt.Println("unknow type")
		}
	}
}

func main() {

	var str interface{} = "hello,world"
	var num interface{} = 1
	var mp map[string]string = map[string]string{"a": "1", "b": "2"}
	var fVal = 1.3
	TestQueryType(str)
	TestQueryType(num)
	TestQueryType(mp)
	TestQueryType(fVal)
}
