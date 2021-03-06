/*
使用说明：这个程序用来给词频计数
输入说明:
    --import-file 表明输入的是单个文件，每个文件以\n 结尾
    --total-file 表明只计算整个文件的词频统计
    --import-dir 表明输入是目录，统计目录下的所有文件，只支持一级目录
    --help 打印说明
*/

package main

import ("fmt"
	//"strings"
	"os"
	"bufio"
	"io"
)

//这个文件将单个的文件打开，并计数
func count_file(import_file_name string,total_file bool){
	import_file,err := os.Open(import_file_name)
	if err!= nil{
		fmt.Println("读文件有错误",err)
		os.Exit(3)
	}

	defer import_file.Close()

	reader := bufio.NewReader(import_file)

	eof := false
	count := 0
	for !eof {
		var line string
		line,err = reader.ReadString('\n')
		if err == io.EOF {
			eof = true
		}
		if len(line) >0 {
			fmt.Printf("%d %s",count,line)
			count ++
		}
	}
	
}

func main(){
	begin := 1
	judge := true
	import_file_name := ""
	total_file := false
	for begin<= len(os.Args) -1 && judge {
		if os.Args[begin] == "--help"{
			fmt.Println("使用说明：这个程序用来给词频计数\n"+
                                    "输入说明:\n"+
                                    "     --import-file 表明输入的是单个文件，每个文件以\\n 结尾\n"+
                                    "     --total-file 表明只计算整个文件的词频统计\n"+
                                    "     --import-dir 表明输入是目录，统计目录下的所有文件，只支持一级目录\n"+
                                    "     --help 打印说明\\n"+
				    "一个初级的例子是====>:\n"+
				    "                  go run wc --import-file test.txt --total-file")
			begin++
			os.Exit(1)
		} else if os.Args[begin] == "--import-file"{
			if begin >= len(os.Args)-1 {
				fmt.Println("未找到输入文件")
				os.Exit(1)
			}
			import_file_name = os.Args[begin+1]
			begin += 2
			fmt.Println("输入文件:",import_file_name)
		} else if os.Args[begin] == "--total-file"{
			begin++
			fmt.Println("使用\\n作为每个文件结尾",total_file)
		} else if os.Args[begin] == "--import-dir"{
			begin++
			fmt.Println("程序还未开发完成，敬请期待")
			os.Exit(1)
		} else{
			fmt.Println("输入的命令有错误")
			os.Exit(2)
		}
	}

	//下一步就是开文件，然后将具体的参数传进去了。
	
	count_file(import_file_name,total_file)
	

}
