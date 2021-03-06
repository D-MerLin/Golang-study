## Go的声明语法
Go声明变量或者函数时，变量的类型在变量名或函数名之后，这与C语言不同。

---

### C语言
C语言在声明变量时，没有使用特殊的语法，左侧是基本类型，右侧是一个包含变量的表达式。写一个新变量时，右侧的表达式与左侧的基本类型存在关联，即需要通过两者共同确定变量的类型：
```c
int x;
int *p;
int a[3];
```
`p`是指向int的一个指针，`a`是int的数组。在声明函数时，基本类型也在函数名的左侧，指定了返回值的类型：
```c
int main(int argc, char *argv[]) { /* ... */ }
```
这种方法使用起来很简单，但对于一些例子就会容易难以理解，比如声明一个函数指针：
```c
int (*fp)(int a, int b)
```
`fp`是一个函数的指针，它的类型是`int (*)(int, int)`，表达式`(*fp)(a,b)`会返回一个int。但当`fp`的一个参数本身就是一个函数时，就难以阅读：
```c
int (*fp)(int (*ff)(int x, int y), int b)
```
或者省略参数名：
```c
int (*fp)(int (*)(int, int), int)
```
甚至更加难读懂。
又或者返回的类型是函数指针时：
```c
int (*(*fp)(int (*)(int, int), int))(int, int)
```
除此之外类型和声明的语法是一样的，因此解析中间有类型的表达式会很困难，这也C在类型转换时要用括号原因：
```c
(int)PI
```

### Go语言
有其他的一些语言在声明时使用不同的语法，例如（假设有一种语言）名字后面跟冒号和类型：
```c
x: int
p: pointer to int
a: array[3] of int
```
这读起来很清楚，Go便使用了如下的形式:
```go
x int
p *int
a [3]int
```
与C不同，左侧的变量名与右侧的类型并不存在直接关联，只需要右侧部分就可以知道变量的类型。也就是说Go用单独的语法为代价，获取了更好的可读性。函数在声明的时候表面上和C的差别不大：
```go
func main(argc int, argv []string) int
func main(int, []string) int
```
这种从左到右的样式的一个优点是，当类型变得更复杂时，它的工作效果很好。下面是一个函数变量的声明(类似于C中的函数指针):包括参数是函数：
```go
f func(func(int,int) int, int) int
```
或者返回类型是函数时
```go
f func(func(int,int) int, int) func(int, int) int
```

### 指针
对于数组Go把括号放在类型的左边:
```go
var a []int
x = a[1]
```
并且Go中的指针沿用了C中的符号`*`:
```go
var p *int
x = *p
```
但不能把`*`放在后面，因为会和乘法混淆：
```go
var p *int
x = p*  // 错误用法
```
总之，Go在指针方面与C语言形似，这也就无法避免使用括号来消除歧义，比如类型转换：
```go
（*int)(nil)
```

---

**参考资料**

[Go's Declaration Syntax](https://blog.go-zh.org/gos-declaration-syntax)