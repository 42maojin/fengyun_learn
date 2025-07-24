package main
import(
	."fmt"
	"strings"
	"bufio"
	"os"
	"math/rand"
	"time"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	choices := []string{"石头","剪刀","布"}

	
	Println("剪刀石头布，请输入你的选择（剪刀，石头，布）：")
	reader := bufio.NewReader(os.Stdin)//带缓冲读取器，用于高效标准输入
	input,_:=reader.ReadString('\n')//读取字符串，直到遇到\n
	input = strings.TrimSpace(input)//去除首尾的空格\n,\t等

}