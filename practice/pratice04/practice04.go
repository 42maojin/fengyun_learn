// package main//日期转换为时间戳（北京时间）,同时指出所在季节
// import (
// 	"fmt"
// 	"bufio"
// 	"os"
// 	"strings"
// 	"time"
// 	"strconv"
// )


// func main(){
// 		fmt.Println("请输入日期：（如2025-07-24,仅公元后）：")
// 		reader := bufio.NewReader(os.Stdin)
// 		input,_ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		t,err := time.Parse("2006-01-02",input)//神圣日期，数字如01，2006什么不能变，但是2006/01/02可以改变函数的解析格式
// 		if err!= nil{
// 			fmt.Println("解析失败",err)
// 			return
// 		}
// 		timestamp:=t.Unix()
// 		fmt.Printf("对应的时间戳为：%d\n",timestamp)
// 		parts := strings.Split(input,"-")
// 		part,_:= strconv.Atoi(parts[1])
// 		switch {
// 			case part>=3 && part <= 5:
// 				fmt.Println("这是在春季")
// 			case part>=6 && part <= 8:
// 				fmt.Println("这是在夏季")
// 			case part>=9 && part <= 11:
// 				fmt.Println("这是在秋季")
// 			case part==12 || part == 1 || part == 2:
// 				fmt.Println("这是在冬季")
// 			default:
// 				fmt.Println("出错")
// 		}
// }




//数组相同元素下标问题
//读取两个字符串，第一个字符串A用map存起来，键是对应字符，值是序号,序号可以有多个，因为一个字符不一定只出现一次
//B挨个解析，比较对应的键，一旦涉及到一样的，添加在对应键的值，B不一样的地方在于用不同方式计数，如用a,b,c,d来计数
package main
import(
	"fmt"
	"bufio"
	"os"
	"strings"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入字符串A:")
	inputA,_ := reader.ReadString('\n')
	inputA = strings.TrimSpace(inputA)
	fmt.Println("请输入字符串B:")
	inputB,_ := reader.ReadString('\n')
	inputB = strings.TrimSpace(inputB)
	fmt.Printf("A:%s\nB:%s",inputA,inputB)
}