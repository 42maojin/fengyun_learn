// package main //循环猜拳游戏
// import (
// 	"bufio"
// 	. "fmt"
// 	"math/rand"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	reader := bufio.NewReader(os.Stdin) //带缓冲读取器，用于高效标准输入
// 	Println("剪刀石头布，请输入你的选择（剪刀，石头，布）：")
// 	for {
// 		input, _ := reader.ReadString('\n') //读取字符串，直到遇到\n
// 		input = strings.TrimSpace(input)    //去除首尾的空格\n,\t等
// 		choices := []string{"石头", "剪刀", "布"}
// 		index := rand.Intn(len(choices))
// 		comchoice := choices[index]

// 		switch {
// 		case input == comchoice:
// 			Println("平局！")
// 		case input == "剪刀":
// 			if comchoice == "布" {
// 				Println("你赢了！")
// 			} else {
// 				Println("你输了！")
// 			}
// 		case input == "石头":
// 			if comchoice == "剪刀" {
// 				Println("你赢了！")
// 			} else {
// 				Println("你输了！")
// 			}
// 		case input == "布":
// 			if comchoice == "石头" {
// 				Println("你赢了！")
// 			} else {
// 				Println("你输了！")
// 			}
// 		default:
// 			Println("请正确输入！")
// 			return
// 		}
// 		Printf("电脑出的是%s\n", choices[index])
// 		Printf("还玩吗？（yes/no):")
// 		a, _ := reader.ReadString('\n')
// 		a = strings.TrimSpace(a)
// 		if a == "yes" {
// 			Println("游戏继续，请输入你的选择（剪刀，石头，布）：")
// 		} else if a == "no" {
// 			Println("再见！")
// 			return
// 		} else {
// 			Println("请正确输入！已退出！")
// 			return
// 		}

// 	}
// }

//班级花名册题目
package main
import(
	"fmt"
	"os"
	"bufio"
	"strings"
)
func main(){
	fmt.Println("请输入要查询的学生姓名：")
	reader := bufio.NewReader(os.Stdin)
	input,_:=reader.ReadString('\n')
	input = strings.TrimSpace(input)
	map_name := map[string]int{"a":1,"ab":2,"小王":3,"小陈":4,"c":5,"sari":6,"apple":7,"123":8,"小高":9,"小刘":10,"老王":11,"lisa":12,"ci":13,"d":14,"e":15,"6":16,"老吴":17,"12":18,"33":19,"10":20}

	_,ok := map_name[input]
	if ok == true {
		fmt.Println("是我们班的同学")
	} else{
		fmt.Println("查无此人")
	}
	
}
