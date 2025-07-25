package main
import "fmt"

func main(){
	var bianliang int
	fmt.Println("abc =",bianliang)
	bianliang = 10
	fmt.Println("abc2=",bianliang)
    a,b := 10,20
	a,b = b,a
	fmt.Printf("a = %d,b =%d",a,b)
} 