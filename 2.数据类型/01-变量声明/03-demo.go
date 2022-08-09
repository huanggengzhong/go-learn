package main

import (
	"fmt"
)

func main(){
	fmt.Print("输出到控制台不换行")
	fmt.Println("------")
	fmt.Println("输出到控制台并换行")
	fmt.Printf("name=%s,age=%d\n","Tom",30)
	fmt.Printf("name=%s,age=%d,height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.567))
}

/* 
%s	直接输出字符串
%d	表示为十进制
%v	值的默认格式表示
%.2f   默认宽度，精度2
\n   换行
*/