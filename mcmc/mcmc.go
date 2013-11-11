//这个程序用于实践一下我的go语言
//这个程序是求解pi的mcmc程序，可能写的比较简陋

package main

import ("fmt"
	"math/rand"
	"time"
       )

func main(){

	all := 0

	x := 0.0
	y := 0.0

	iter := 100000000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0 ; i < iter; i++ {
		x = r.Float64()*2
		y = r.Float64()*2
		if x*x + y*y <= 4{
			all += 1
		}
	}

	fmt.Println("Pi is ",float64(all)/float64(iter)*4)
}
