### 1.变量声明的几个方法

1.不赋值有默认零值
```js
var var1 init
/*
常量默认值：
int 0
int8 0
int32 0
int64 0
rune 0 // rune 的实际类型是 int32
float32 0 // 长度为 4 byte
float64 0 // 长度为 8 byte
bool false
string “”
*/
```
2.简短声明
```js
var2:="我是简洁声明字符串"
```

3.var 变量名=value形式
```js
var var3=100//会自动类型推导
```

4.一次普通多个赋值
```js
var var4,var5,var6,var7=1,"haha",3.14,true
```

5.一次简洁多个赋值
```js
var4,var5,var6,var7:=1,"haha",3.14,true
```

6.在函数体外部声明多个全局变量
```js
var (
    variablesx int
    slicex []int
    interfacex interface{}
)
```
7.变量交换
```js
var i=100
var j=200
i,j=j,i
```

### 注意事项
1.Go 对于已声明但未使用的变量会在编译阶段报错，比如下面的代码就会产生一个错误：声明了 i 但未使用。

2._（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。
```js
_,x:=122,211;
fmt.Println(_)//报错,cannot use _ as value

```
3.访问权限
大写字母开头的变量是可导出的，即其它包（后面会讲到）可以读取的，是公有变量（相当于传统编程语言中class的public权限修饰符）；小写字母开头的就是不可导出的，是私有变量，仅本包可以使用（相当于传统编程语言中class的private权限修饰符）。
