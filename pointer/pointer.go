//测试指针

package main

import ("fmt"
       )


type composer struct{
	name string
	birth int
}

func main(){

	//新建一个结构体
	fmt.Println("com1")
	com1 := composer{"哈尔滨",1990}
	fmt.Println(com1)
	
	fmt.Println("com2")
	com2 := new(composer)
	com2.name,com2.birth = "南京",1992
	fmt.Println(com2)

	array := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(len(array))

	a2 := make([]composer,2,5)
	fmt.Println(a2)
	
}

