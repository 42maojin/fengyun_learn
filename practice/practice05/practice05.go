// package main//数组最大最小元素问题
// import(
// 	"fmt"
// 	"os"
// 	"bufio"
// 	"strings"
// 	"strconv"
// )
// func main(){
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Println("请输入数组A元素：")
//     scanner.Scan() // 读取一行
//     inputA := scanner.Text()
//     // 按空格分割字符串
//     fieldsA := strings.Fields(inputA)
//     arrA := make([]int, len(fieldsA))
//     // 转换并填充数组
//     for i, s := range fieldsA {
//         num, err1 := strconv.Atoi(s)
// 		if err1 != nil{
// 			fmt.Printf("转换失败：%v\n",err1)
// 			return
// 		}
// 		arrA[i] = num
//     }

// 	fmt.Println("请输入字符串B:")
//     scanner.Scan() // 读取一行
//     inputB := scanner.Text()
//     // 按空格分割字符串
//     fieldsB := strings.Fields(inputB)
//     arrB := make([]int, len(fieldsB))
//     // 转换并填充数组
//     for i, s := range fieldsB {
//         num, err2 := strconv.Atoi(s)
// 		if err2 != nil{
// 			fmt.Printf("转换失败：%v\n",err2)
// 			return
// 		}
//         arrB[i] = num
//     }
//     maxA := arrA[0]
// 	returnA := 0
// 	minB := arrB[0]
// 	returnB := 0
// 	for i:=1;i<len(arrA);i++{
// 		if maxA < arrA[i]{
// 			maxA = arrA[i]
// 			returnA = i
// 		}
// 	}

// 	for i:=1;i<len(arrB);i++{
// 		if minB > arrB[i]{
// 			minB = arrB[i]
// 			returnB = i
// 		}
// 	}
// 	sum := maxA+minB
// 	fmt.Printf("A中最大元素下标：%d\nB中最小元素下标：%d\n总和：%d",returnA,returnB,sum)
// }

// 怎么应对可能的多答案问题？输出第一个实现的答案
// 还要注意防备根本没有答案的情况
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func main(){
	var target int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入数组(用空格分隔）：")
    scanner.Scan() // 读取一行
    input := scanner.Text()
    // 按空格分割字符串
    fieldsA := strings.Fields(input)
    arrA := make([]int, len(fieldsA))
    // 转换并填充数组
    for i, s := range fieldsA {
        num, err1 := strconv.Atoi(s)
		if err1 != nil{
			fmt.Printf("转换失败：%v\n",err1)
			return
		}
		arrA[i] = num
    }
	fmt.Println("请输入目标值：")
	fmt.Scanf("%d",&target)
	for i:=0;i<len(arrA);i++{
		for j:=i+1;j<len(arrA);j++{
			if arrA[i]+arrA[j] == target{
				fmt.Printf("[%d,%d]",i,j)
				return
			}
		}
	}
	fmt.Println("没有符合要求的答案！")
}