package main 
import (
	"fmt"
)

func main(){
	// 一维数组

	var arr_1 [5] int;
	fmt.Println(arr_1)

	var arr_2=[5] int {1,2,3,4,5}
	fmt.Println(arr_2)

	arr_3:=[5] int {1,2,3,4,5}
	fmt.Println(arr_3)
	arr_4:=[...] int {1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	arr_5:=[5] int {0:3,1:5,4:6}
	fmt.Println(arr_5)

	// 二维数组
	
}