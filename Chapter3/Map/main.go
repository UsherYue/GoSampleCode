package main

import "fmt"

func main() {
	mp := map[string]string{} //通过字面的方式来初始化
	//mp:=make(map[string],string)
	mp["zhangsan"] = "15"
	fmt.Println(mp)
	delete(mp, "zhangsan")
	fmt.Println(mp)
	mp["zhangsan"] = "15"
	mp["lisi"] = "16"
	mp["wangwu"] = "17"
	fmt.Println(mp)
	//我们来判断一下指定的key是否存在于map中
	if v, ok := mp["lisi"]; ok {
		fmt.Println("lisi 存在,值为:", v)
	} else {
		fmt.Println("lisi不存在")
	}
	if v, ok := mp["lisi1"]; ok {
		fmt.Println("lisi1 存在,值为:", v)
	} else {
		fmt.Println("lisi1不存在")
	}
	//遍历map
	for k, v := range mp {
		fmt.Println("k=", k, " v=", v)
	}
	//我们将map当做函数的参数来使用
	testFunc := func(mp map[string]string) {
		mp["testkey"] = "testvalue"
	}
	fmt.Println(mp)
	testFunc(mp)
	fmt.Println(mp)

}
