package main//完成日期一系列功能，包括输入日期判断星期几，判断年份等
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
)

func isDataValid(data string) bool{
	_,err := time.Parse("2006-01-02",data)//神圣日期，数字如01，2006什么不能变，但是2006/01/02可以改变函数的解析格式
	return err == nil
}


func main(){
	//输入日期判断星期几使用基姆拉尔森公式
	var a int
	var ayear int
	fmt.Println("请输入想要的功能：\n输入“1”，查询年份是否为闰年\n输入“2”,查询日期对应的星期数\n请输入：")
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d",&a)
	reader.ReadString('\n') // 清除残留的换行符



	if a==1 {//闰年判断功能
		fmt.Println("请输入年份：（如2025,仅公元后）：")
		fmt.Scanf("%d",&ayear)
		if (ayear%4==0&&ayear%100!=0) || (ayear%400 == 0){
			fmt.Printf("%d年是闰年",ayear)
		}else{
			fmt.Printf("%d年不是闰年",ayear)
		}

	} else if a==2{//对应星期功能
		fmt.Println("请输入日期：（如2025-07-24,仅公元后）：")
		reader := bufio.NewReader(os.Stdin)
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input,"-")

		valid := isDataValid(input)
		if !valid {
			fmt.Println("不存在该日期，请按格式正确输入！")
			return
		}

		// if len(parts) != 3 {
		// 	fmt.Println("非正确日期，请按格式正确输入")
		// 	return
		// }
		year := parts[0]
		month := parts[1]
		day := parts[2]


		year1,_ := strconv.Atoi(year)//字符串转数字模块
		month1,_ := strconv.Atoi(month)
		day1,_ := strconv.Atoi(day)


		if month1 == 1 || month1 == 2	{//1，2月份的特殊处理
			month1 += 12
			year1--	
		}
		
		//核心公式
		w := (day1 + 2*month1 + 3*(month1+1)/5 + year1 + year1/4 - year1/100 + year1/400 + 1)%7
		

		//输出
		switch w{
			case 0:
				fmt.Println("当天是星期日")
			case 1:
				fmt.Println("当天是星期一")
			case 2:
				fmt.Println("当天是星期二")
			case 3:
				fmt.Println("当天是星期三")
			case 4:
				fmt.Println("当天是星期四")
			case 5:
				fmt.Println("当天是星期五")
			default:
				fmt.Println("当天是星期六")
		}
	} else{//保护
		fmt.Println("请正确输入！")
	}
		
}