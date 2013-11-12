//在这个程序里，主要演示go的字符串操作，都是基本必会的那种，特别高端的等着慢慢学

package main

import ("fmt"
	"strings"
       )

func main(){
	//演示遍历字符串
	fmt.Println("测试打印字符串，索引")
	hello := "hello world"
	for index,value := range hello{
		fmt.Print(index,value,"|")
	}
	fmt.Print("\n")

	//注意这里由于是utf-8的所以三位一个码
	fmt.Println("测试汉语")
	hello = "我在哈尔滨工业大学上学"
	for index,value := range hello{
		fmt.Print(index,string(value),strings.Index(hello,string(value)),"|")
	}
	fmt.Print("\n")

	hello = "I study in harbin institute of tech"
	fmt.Println("测试string split")
	//for index,value := range strings.Split(hello," "){
	for index,value := range strings.Fields(hello){
		fmt.Printf("%d %s|",index,value)
	}
	fmt.Print("\n")
}
