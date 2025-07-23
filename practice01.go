// package main
// import "fmt"
// func main(){
// 	// var num int
// 	// fmt.Println("请输入数值：")
// 	// fmt.Scanf("%d",&num)
// 	// for i:=1;i<=num;i++{
// 	// 	for j:= 1; j<=i;j++ {
// 	// 		fmt.Printf("%d ",i*j)
// 	// 	}
// 	// 	fmt.Printf("\n")
// 	// }

// }



package main
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

)

func main(){
	fmt.Println("请输入分数（整数）：")
	reader := bufio.NewReader(os.Stdin)//带缓冲读取器，用于高效读取标准输入
	input,_:=reader.ReadString('\n')//读取字符串，直到遇到\n
	input = strings.TrimSpace(input)//去除收尾的空格和\t,\n等  *此时input是字符串

	if strings.Contains(input,"."){
		fmt.Println("本成绩不含小数!请输入整数！")
		return
	}
	
	if match,_ := regexp.MatchString(`^\d+$`,input) //注意不是单引号
	!match {
		fmt.Println("请输入正确成绩！仅包含数字！")
		return
	}

	score,err := strconv.Atoi(input)
	if err != nil{
		fmt.Println("错误！无效整数格式！")
	}
	
	switch {
	   case score >= 0 && score <= 59:
			fmt.Println("不及格")
		case score >= 60 && score <= 80:
			fmt.Println("中等")
		case score >= 81 && score <= 90:
			fmt.Println("良好")
		case score >= 91 && score <= 99:
			fmt.Println("优秀")
		case score == 100:
			fmt.Println("完美")
		default:
			fmt.Println("请输入正确的分数！")

	}
}