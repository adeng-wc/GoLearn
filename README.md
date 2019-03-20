# Go语言的学习之旅

# 1 安装

[golang官网自行下载](https://golang.org/dl/) 

Golang 的主要关注点是使得高可用性和可扩展性的 Web 应用的开发变得简便容易。（译注：Go 的定位是系统编程语言，只是对 Web 开发支持较好）



## Mac OS

在 <https://golang.org/dl/> 下载安装程序。双击开始安装并且遵循安装提示，会将 Golang 安装到 `/usr/local/go` 目录下，同时 `/usr/local/go/bin` 文件夹也会被添加到 `PATH` 环境变量中。



## 建立 Go 工作区

在编写代码之前，我们首先应该建立 Go 的工作区（Workspace）。

在 **Mac 或 Linux** 操作系统下，Go 工作区应该设置在 **$HOME/go**。所以我们要在 **$HOME** 目录下创建 **go** 目录。

也可以通过设置 GOPATH 环境变量，用其他目录来作为工作区。

所有 Go 源文件都应该放置在工作区里的 **src** 目录下。请在刚添加的 **go** 目录下面创建目录 **src**。

所有 Go 项目都应该依次在 src 里面设置自己的子目录。我们在 src 里面创建一个目录 **hello** 来放置整个 hello world 项目。

创建上述目录之后，其目录结构如下：

```xml
go
  src
    hello
	  helloworld.go
```



## 运行 Go 程序

1. 使用 `go run` 命令，输入 `go run workspacepath/src/hello/helloworld.go`。
2. 使用 `go install` 命令，接着可以用 `workspacepath/bin/hello` 来运行该程序。

**package main - 每一个 Go 文件都应该在开头进行 package name 的声明**（译注：只有可执行程序的包名应当为 main）。包（Packages）用于代码的封装与重用，这里的包名称是`main`。

**import "fmt"** - 我们引入了 fmt 包，用于在 main 函数里面打印文本到标准输出。

**func main()** - main 是一个特殊的函数。整个程序就是从 main 函数开始运行的。**main 函数必须放置在 main 包中**。`{` 和 `}` 分别表示 main 函数的开始和结束部分。

**fmt.Println("Hello World")** - **fmt** 包中的 **Println** 函数用于把文本写入标准输出。



# 2 基础语法

## 2.1 变量（Variables）

变量指定了某存储单元（Memory Location）的名称，该存储单元会存储特定类型的值。

### 声明单个变量

**var name type** 是声明单个变量的语法。

语句 `var age int` 声明了一个 int 类型的变量，名字为 age。我们还没有给该变量赋值。如果变量未被赋值，Go 会自动地将其初始化，赋值该变量类型的零值（Zero Value）

### 声明变量并初始化

 **var name type = initialvalue** 的语法用于声明变量并初始化。

### 类型推断（Type Inference）

如果变量有初始值，那么 Go 能够自动推断具有初始值的变量的类型。因此，如果变量有初始值，就可以在变量声明中省略 `type`。

如果变量声明的语法是 **var name = initialvalue**，Go 能够根据初始值自动推断变量的类型。

### 声明多个变量

声明多个变量的语法是 **var name1, name2 type = initialvalue1, initialvalue2**。

### 简短声明

Go 也支持一种声明变量的简洁形式，称为简短声明（Short Hand Declaration），该声明使用了 **:=** 操作符。

声明变量的简短语法是 **name := initialvalue**。

## 2.2 类型

下面是 Go 支持的基本类型：

- bool
- 数字类型
  - int8, int16, int32, int64, int
  - uint8, uint16, uint32, uint64, uint
  - float32, float64
  - complex64, complex128
  - byte
  - rune
- string

### bool

bool 类型表示一个布尔值，值为 true 或者 false。

### 有符号整型

**int8**：表示 8 位有符号整型
**大小**：8 位
**范围**：-128～127

**int16**：表示 16 位有符号整型
**大小**：16 位
**范围**：-32768～32767

**int32**：表示 32 位有符号整型
**大小**：32 位
**范围**：-2147483648～2147483647

**int64**：表示 64 位有符号整型
**大小**：64 位
**范围**：-9223372036854775808～9223372036854775807

**int**：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 *int*表示整型。
**大小**：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
**范围**：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807。

### 无符号整型

**uint8**：表示 8 位无符号整型
**大小**：8 位
**范围**：0～255

**uint16**：表示 16 位无符号整型
**大小**：16 位
**范围**：0～65535

**uint32**：表示 32 位无符号整型
**大小**：32 位
**范围**：0～4294967295

**uint64**：表示 64 位无符号整型
**大小**：64 位
**范围**：0～18446744073709551615

**uint**：根据不同的底层平台，表示 32 或 64 位无符号整型。
**大小**：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
**范围**：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。

### 浮点型

**float32**：32 位浮点数
**float64**：64 位浮点数

### 复数类型

**complex64**：实部和虚部都是 float32 类型的的复数。
**complex128**：实部和虚部都是 float64 类型的的复数。

内建函数 [**complex**](https://golang.org/pkg/builtin/#complex) 用于创建一个包含实部和虚部的复数。complex 函数的定义如下：

```go
func complex(r, i FloatType) ComplexType
```

该函数的参数分别是实部和虚部，并返回一个复数类型。实部和虚部应该是相同类型，也就是 float32 或 float64。如果实部和虚部都是 float32 类型，则函数会返回一个 complex64 类型的复数。如果实部和虚部都是 float64 类型，则函数会返回一个 complex128 类型的复数。

还可以使用简短语法来创建复数：

```go
c := 6 + 7i
```

说明： [Wiki 复数详解](https://zh.wikipedia.org/wiki/%E5%A4%8D%E6%95%B0_(%E6%95%B0%E5%AD%A6)) 至今没有搞明白 复数的实际应用场景。

### 其他数字类型

**byte** 是 uint8 的别名。
**rune** 是 int32 的别名。

在学习字符串的时候，我们会详细讨论 byte 和 rune。

### string 类型

在 Golang 中，字符串是字节的集合。如果你现在还不理解这个定义，也没有关系。我们可以暂且认为一个字符串就是由很多字符组成的。

### 类型转换

Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换。没有隐式转换。

## 2.3 常量

在 Go 语言中，术语"常量"用于表示固定的值。

**变量 a 和 b 分别被赋值为常量 50 和 I love GO**。关键字 `const` 被用于表示常量，比如 `50` 和 `I love Go`。即使在上面的代码中我们没有明确的使用关键字 `const`，但是在 Go 的内部，它们是常量。

顾名思义，常量不能再重新赋值为其他的值。

常量仍然没有类型。

Go 是一门强类型语言，所有的变量必须有明确的类型。

Go 是一个强类型的语言，在分配过程中混合类型是不允许的。

**无类型的常量有一个与它们相关联的默认类型，并且当且仅当一行代码需要时才提供它。在声明中 var name = "Sam" ， name需要一个类型，它从字符串常量 Sam 的默认类型中获取。**

## 2.4 函数

函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出。

### 函数的声明

在 Go 语言中，函数声明通用语法如下：

```go
func functionname(parametername type) returntype {  
    // 函数体（具体实现的功能）
}
```

函数的声明以关键词 `func` 开始，后面紧跟自定义的函数名 `functionname (函数名)`。函数的参数列表定义在 `(` 和 `)` 之间，返回值的类型则定义在之后的 `returntype (返回值类型)`处。声明一个参数的语法采用 **参数名** **参数类型** 的方式，任意多个参数采用类似 `(parameter1 type, parameter2 type) 即(参数1 参数1的类型,参数2 参数2的类型)`的形式指定。之后包含在 `{` 和 `}` 之间的代码，就是函数体

### 多返回值

Go 语言支持一个函数可以有多个返回值。



### 命名返回值

从函数中可以返回一个命名值。一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了。

```go
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
```

请注意, 函数中的 return 语句没有显式返回任何值。由于 **area** 和 **perimeter** 在函数声明中指定为返回值, 因此当遇到 return 语句时, 它们将自动从函数返回。



### 空白符

**_** 在 Go 中被用作空白符，可以用作表示任何类型的任何值。

>  在程序的 `area, _ := rectProps(10.8, 5.6)` 这一行，我们看到空白符 `_` 用来跳过不要的计算结果。



## 2.5 包

**包用于组织 Go 源代码，提供了更好的可重用性与可读性**。由于包提供了代码的封装，因此使得 Go 应用程序易于维护。



### main 函数和 main 包

所有可执行的 Go 程序都必须包含一个 main 函数。这个函数是程序运行的入口。main 函数应该放置于 main 包中。

**package packagename 这行代码指定了某一源文件属于一个包。它应该放在每一个源文件的第一行。**



工作区的目录结构会是这样:

```
src
    geometry
        gemometry.go
bin
    geometry
```

### 创建自定义的包

**属于某一个包的源文件都应该放置于一个单独命名的文件夹里。按照 Go 的惯例，应该用包名命名该文件夹。**

因此，我们在 `geometry` 文件夹中，创建一个命名为 `rectangle` 的文件夹。在 `rectangle` 文件夹中，所有文件都会以 `package rectangle` 作为开头，因为它们都属于 rectangle 包。



### 导入自定义包

为了使用自定义包，我们必须要先导入它。导入自定义包的语法为 `import path`。我们必须指定自定义包相对于工作区内 `src` 文件夹的相对路径。我们目前的文件夹结构是：

```go
src
    geometry
        geometry.go
        rectangle
            rectprops.go
```

`import "geometry/rectangle"` 这一行会导入 rectangle 包。



### 导出名字（Exported Names）

我们将 rectangle 包中的函数 Area 和 Diagonal 首字母大写。在 Go 中这具有特殊意义。在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。在这里，我们需要在 main 包中访问 Area 和 Diagonal 函数，因此会将它们的首字母大写。

在 `rectprops.go` 中，如果函数名从 `Area(len, wid float64)` 变为 `area(len, wid float64)`，并且在 `geometry.go` 中， `rectangle.Area(rectLen, rectWidth)` 变为 `rectangle.area(rectLen, rectWidth)`， 则该程序运行时，编译器会抛出错误 `geometry.go:11: cannot refer to unexported name rectangle.area`。因为如果想在包外访问一个函数，它应该首字母大写。



### init 函数

所有包都可以包含一个 `init` 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。init 函数的形式如下：

```go
func init() {  
}
```

init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性。

包的初始化顺序如下：

1. 首先初始化包级别（Package Level）的变量
2. 紧接着调用 init 函数。包可以有多个 init 函数（在一个文件或分布于多个文件中），它们按照编译器解析它们的顺序进行调用。

如果一个包导入了另一个包，会先初始化被导入的包。

尽管一个包可能会被导入多次，但是它只会被初始化一次。

### 使用空白标识符（Blank Identifier）

导入了包，却不在代码中使用它，这在 Go 中是非法的。当这么做时，编译器是会报错的。其原因是为了避免导入过多未使用的包，从而导致编译时间显著增加。

然而，在程序开发的活跃阶段，又常常会先导入包，而暂不使用它。遇到这种情况就可以使用空白标识符 `_`。

`var _ = rectangle.Area` 这一行屏蔽了错误。我们应该了解这些错误屏蔽器（Error Silencer）的动态，在程序开发结束时就移除它们，包括那些还没有使用过的包。由此建议在 import 语句下面的包级别范围中写上错误屏蔽器。



## 2.6 if-else 语句

if 是条件语句。if 语句的语法是

```go
if condition {  
}
```

如果 `condition` 为真，则执行 `{` 和 `}` 之间的代码。

不同于其他语言，例如 C 语言，Go 语言里的 `{ }` 是必要的，即使在 `{ }` 之间只有一条语句。

if 语句还有可选的 `else if` 和 `else` 部分。

```go
if condition {  
} else if condition {
} else {
}
```

`if` 还有另外一种形式，它包含一个 `statement` 可选语句部分，该组件在条件判断之前运行。它的语法是

```go
if statement; condition {  
}
```

### 一个注意点

`else` 语句应该在 `if` 语句的大括号 `}` 之后的同一行中。如果不是，编译器会不通过。



## 2.7 循环

`for` 是 Go 语言唯一的循环语句。Go 语言中并没有其他语言比如 C 语言中的 `while` 和 `do while` 循环。

### for 循环语法

```go
for initialisation; condition; post {  
}
```

初始化语句只执行一次。循环初始化后，将检查循环条件。如果条件的计算结果为 `true` ，则 `{}` 内的循环体将执行，接着执行 post 语句。post 语句将在每次成功循环迭代后执行。在执行 post 语句后，条件将被再次检查。如果为 `true`, 则循环将继续执行，否则 for 循环将终止。

### break

`break` 语句用于在完成正常执行之前突然终止 for 循环，之后程序将会在 for 循环下一行代码开始执行。

### continue

`continue` 语句用来跳出 `for` 循环中当前循环。在 `continue` 语句后的所有的 `for` 循环语句都不会在本次循环中执行。循环体会在一下次循环中继续执行。

### 无限循环

无限循环的语法是：

```go
for {  
}
```

## 2.8 switch 语句

switch 是一个条件语句，用于将表达式的值与可能匹配的选项列表进行比较，并根据匹配情况执行相应的代码块。它可以被认为是替代多个 `if else` 子句的常用方式。



### 默认情况（Default Case）

当其他情况都不匹配时，将运行默认情况。

### 多表达式判断

通过用逗号分隔，可以在一个 case 中包含多个表达式。

在 `case "a","e","i","o","u":` 这一行中，列举了所有的元音。只要匹配该项，则将输出 `vowel`。



### 无表达式的 switch

在 switch 语句中，表达式是可选的，可以被省略。如果省略表达式，则表示这个 switch 语句等同于 `switch true`，并且每个 `case`表达式都被认定为有效，相应的代码块也会被执行。



### Fallthrough 语句

在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。使用 `fallthrough` 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。

**fallthrough 语句应该是 case 子句的最后一个语句。如果它出现在了 case 语句的中间，编译器将会报错：fallthrough statement out of place**



# 3 内建容器



## 3.1 数组

### 数组

数组是同一类型元素的集合。例如，整数集合 5,8,9,79,76 形成一个数组。Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。（当然，如果是 interface{} 类型数组，可以包含任意类型）

### 数组的声明

一个数组的表示形式为 `[n]T`。`n` 表示数组中元素的数量，`T` 代表每个元素的类型。元素的数量 `n` 也是该类型的一部分。

**var a[3]int** 声明了一个长度为 3 的整型数组。**数组中的所有元素都被自动赋值为数组类型的零值。** 在这种情况下，`a` 是一个整型数组，因此 `a` 的所有元素都被赋值为 `0`，即 int 型的零值。运行上述程序将 **输出** `[0 0 0]`。

数组的索引从 `0` 开始到 `length - 1` 结束。

```go
 var a [3]int //i
 b := [3]int{12, 78, 50} 
```

`a := [3]int{12}` 声明一个长度为 3 的数组，但只提供了一个值 `12`，剩下的 2 个元素自动赋值为 `0`。你甚至可以忽略声明数组的长度，并用 `...` 代替，让编译器为你自动计算长度。

**数组的大小是类型的一部分**。因此 `[5]int` 和 `[25]int` 是不同类型。数组不能调整大小，不要担心这个限制，因为 `slices` 的存在能解决这个问题。

### 数组是值类型

Go 中的数组是值类型而不是引用类型。这意味着当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。如果对新变量进行更改，则不会影响原始数组。



### 数组的长度

通过将数组作为参数传递给 `len` 函数，可以得到数组的长度。

### 使用 range 迭代数组

`for` 循环可用于遍历数组中的元素。

通过使用 `for` 循环的 **range** 方法来遍历数组。`range` 返回索引和该索引处的值。

`for i, v := range a` 利用的是 for 循环 range 方式。 它将返回索引和该索引处的值。 

如果你只需要值并希望忽略索引，则可以通过用 `_` 空白标识符替换索引来执行。

### 多维数组

Go 语言可以创建多维数组。

```go
func printarray(a [3][2]string) {
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}
```

## 3.2 切片

切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用。

### 创建一个切片

带有 T 类型元素的切片由 `[]T` 表示

使用语法 `a[start:end]` 创建一个从 `a` 数组索引 `start` 开始到 `end - 1` 结束的切片。

### 切片的修改

切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。

`numa [:]` 缺少开始和结束值。开始和结束的默认值分别为 `0` 和 `len (numa)`。两个切片 `nums1` 和 `nums2` 共享相同的数组。

### 切片的长度和容量

切片的长度是切片中的元素数。**切片的容量是从创建切片索引开始的底层数组中元素数。**

### 使用 make 创建一个切片

`func make（[]T，len，cap）[]T `通过传递类型，长度和容量来创建切片。容量是可选参数, 默认值为切片长度。make 函数创建一个数组，并返回引用该数组的切片。

### 追加切片元素

正如我们已经知道数组的长度是固定的，它的长度不能增加。 切片是动态的，使用 `append` 可以将新元素追加到切片上。append 函数的定义是 `func append（s[]T，x ... T）[]T`。

**x ... T** 在函数定义中表示该函数接受参数 x 的个数是可变的。这些类型的函数被称为[可变函数](https://golangbot.com/variadic-functions/)。

有一个问题可能会困扰你。如果切片由数组支持，并且数组本身的长度是固定的，那么切片如何具有动态长度。以及内部发生了什么，当新的元素被添加到切片时，会创建一个新的数组。现有数组的元素被复制到这个新数组中，并返回这个新数组的新切片引用。**现在新切片的容量是旧切片的两倍**。

切片类型的零值为 `nil`。一个 `nil` 切片的长度和容量为 0。可以使用 append 函数将值追加到 `nil` 切片。

也可以使用 `...` 运算符将一个切片添加到另一个切片。

```go
func main() {
    veggies := []string{"potatoes", "tomatoes", "brinjal"}
    fruits := []string{"oranges", "apples"}
    food := append(veggies, fruits...)
    fmt.Println("food:",food)
}
```



### 切片的函数传递

我们可以认为，切片在内部可由一个结构体类型表示。这是它的表现形式，

```go
type slice struct {  
    Length        int
    Capacity      int
    ZerothElement *byte
}
```

切片包含长度、容量和指向数组第零个元素的指针。当切片传递给函数时，即使它通过值传递，指针变量也将引用相同的底层数组。因此，当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见。



### 多维切片

类似于数组，切片可以有多个维度。

### 内存优化

切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。在内存管理方面，这是需要注意的。让我们假设我们有一个非常大的数组，我们只想处理它的一小部分。然后，我们由这个数组创建一个切片，并开始处理切片。这里需要重点注意的是，在切片引用时数组仍然存在内存中。

一种解决方法是使用 [copy](https://golang.org/pkg/builtin/#copy) 函数 `func copy(dst，src[]T)int` 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。

```go
func countries() []string {
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}
```

## 3.3 可变参数函数

可变参数函数是一种参数个数可变的函数。

### 语法

如果函数最后一个参数被记作 `...T` ，这时函数可以接受任意个 `T` 类型参数作为最后一个参数。

请注意只有函数的最后一个参数才允许是可变的。

### 例子

你是否曾经想过 append 函数是如何将任意个参数值加入到切片中的。这样 append 函数可以接受不同数量的参数。

```go
func append(slice []Type, elems ...Type) []Type
```

上面是 append 函数的定义。在定义中 elems 是可变参数。这样 append 函数可以接受可变化的参数。



**可变参数函数的工作原理是把可变参数转换为一个新的切片。**

### 给可变参数函数传入切片

```go
func find(num int, nums ...int)
```

由可变参数函数的定义可知，`nums ...int` 意味它可以接受 `int` 类型的可变参数。

**有一个可以直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。如果这样做，切片将直接传入函数，不再创建新的切片**

```go
find(89, nums)` 替换为 `find(89, nums...)
```

## 3.4 Maps

map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。

通过向 `make` 函数传入键和值的类型，可以创建 map。`make(map[type of key]type of value)` 是创建 map 的语法。

```go
personSalary := make(map[string]int)
```

map 的零值是 `nil`。如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 `make` 函数初始化。

### 给 map 添加元素

给 map 添加新元素的语法和数组相同。

```go
personSalary := make(map[string]int)
personSalary["steve"] = 12000
personSalary["jamie"] = 15000
personSalary["mike"] = 9000
```



### 获取 map 中的元素

获取 map 元素的语法是 `map[key]` 。

如果获取一个不存在的元素，会发生什么呢？map 会返回该元素类型的零值。在 `personSalary` 这个 map 里，如果我们获取一个不存在的元素，会返回 `int` 类型的零值 `0`。

如果我们想知道 map 中到底是不是存在这个 `key`，该怎么做：

```go
value, ok := map[key]
```

遍历 map 中所有的元素需要用 `for range` 循环。

**有一点很重要，当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同。**

### 删除 map 中的元素

删除 `map` 中 `key` 的语法是 [*delete(map, key)*](https://golang.org/pkg/builtin/#delete)。这个函数没有返回值。

### 获取 map 的长度

获取 map 的长度使用 [len](https://golang.org/pkg/builtin/#len) 函数。

### Map 是引用类型

和 [slices](https://golangbot.com/arrays-and-slices/) 类似，map 也是引用类型。当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量。

当 map 作为函数参数传递时也会发生同样的情况。函数中对 map 的任何修改，对于外部的调用都是可见的。

### Map 的相等性

map 之间不能使用 `==` 操作符判断，`==` 只能用来检查 map 是否为 `nil`。

判断两个 map 是否相等的方法是遍历比较两个 map 中的每个元素。



## 3.5 字符串

Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。

Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码。

### 单独获取字符串的每一个字节

由于字符串是一个字节切片，所以我们可以获取字符串的每一个字节。

**len(s) 返回字符串中字节的数量**，然后我们用了一个 for 循环以 16 进制的形式打印这些字节。`%x` 格式限定符用于指定 16 进制编码。

上面的程序输出 `48 65 6c 6c 6f 20 57 6f 72 6c 64`。这些打印出来的字符是 "Hello World" 以 [Unicode UTF-8 编码](https://mothereff.in/utf-8#Hello%20World)的结果。为了更好的理解 go 中的字符串，需要对 Unicode 和 UTF-8 有基础的理解。我推荐阅读一下 <https://naveenr.net/unicode-character-set-and-utf-8-utf-16-utf-32-encoding/> 来理解一下什么是 Unicode 和 UTF-8。

### rune

rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个**代码点**。代码点无论占用多少个字节，都可以用一个 rune 来表示。

```go
func printChars(s string) {
    runes := []rune(s)
    for i:= 0; i < len(runes); i++ {
        fmt.Printf("%c ",runes[i])
    }
}
```

### 字符串的 for range 循环

 Go 给我们提供了一种更简单的方法来做到这一点：使用 **for range** 循环。

```go
func printCharsAndBytes(s string) {
    for index, rune := range s {
        fmt.Printf("%c starts at byte %d\n", rune, index)
    }
}
```

使用 `for range` 循环遍历了字符串。循环返回的是是当前 rune 的字节位置。

### 用字节切片构造字符串

```go
func main() {  
    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)
}
```

```go
func main() {  
    byteSlice := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
    str := string(byteSlice)
    fmt.Println(str)
}
```

### 用 rune 切片构造字符串

```go
func main() {  
    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str := string(runeSlice)
    fmt.Println(str)
}
```

### 字符串的长度

[utf8 package](https://golang.org/pkg/unicode/utf8/#RuneCountInString) 包中的 `func RuneCountInString(s string) (n int)` 方法用来获取字符串的长度。这个方法传入一个字符串参数然后返回字符串中的 rune 的数量。

### 字符串是不可变的

Go 中的字符串是不可变的。一旦一个字符串被创建，那么它将无法被修改。

为了修改字符串，可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串。

```go
func mutate(s []rune) string {  
    s[0] = 'a' 
    return string(s)
}
func main() {  
    h := "hello"
    fmt.Println(mutate([]rune(h)))
}
```

## 3.6 指针

指针是一种存储变量内存地址（Memory Address）的变量。

![指针示意图](https://raw.githubusercontent.com/studygolang/gctt-images/master/golang-series/pointer-explained.png)

如上图所示，变量 `b` 的值为 `156`，而 `b` 的内存地址为 `0x1040a124`。变量 `a` 存储了 `b` 的地址。我们就称 `a` 指向了 `b`。

### 指针的声明

指针变量的类型为 ***T**，该指针指向一个 **T** 类型的变量。

**&** 操作符用于获取变量的地址

### 指针的零值（Zero Value）

指针的零值是 `nil`。

### 指针的解引用

指针的解引用可以获取指针所指向的变量的值。将 `a` 解引用的语法是 `*a`。

### 向函数传递指针参数

```go
func change(val *int) {  
    *val = 55
}
func main() {  
    a := 58
    fmt.Println("value of a before function call is",a)
    b := &a
    change(b)
    fmt.Println("value of a after function call is", a)
}
```

### 不要向函数传递数组的指针，而应该使用切片

假如我们想要在函数内修改一个数组，并希望调用函数的地方也能得到修改后的数组，一种解决方案是把一个指向数组的指针传递给这个函数。

```go
func modify(arr *[3]int) {  
    (*arr)[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)
}
```

**a[x] 是 (\*a)[x] 的简写形式，因此上面代码中的 (*arr)[0] 可以替换为 arr[0]**。

```go
func modify(arr *[3]int) {  
    arr[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)
}
```

**这种方式向函数传递一个数组指针参数，并在函数内修改数组。尽管它是有效的，但却不是 Go 语言惯用的实现方式。我们最好使用切片来处理。**

```go
func modify(sls []int) {  
    sls[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(a[:])
    fmt.Println(a)
}
```

### Go 不支持指针运算

Go 并不支持其他语言（例如 C）中的指针运算。



# 4 面向“对象”

## 4.1 结构体

结构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。

例如，一个职员有 `firstName`、`lastName` 和 `age` 三个属性，而把这些属性组合在一个结构体 `employee` 中就很合理。

### 结构体的声明

```go
type Employee struct {
    firstName string
    lastName  string
    age       int
}
```

上面的结构体 `Employee` 称为 **命名的结构体（Named Structure）**。我们创建了名为 `Employee` 的新类型，而它可以用于创建 `Employee` 类型的结构体变量。

声明结构体时也可以不用声明一个新类型，这样的结构体类型称为 **匿名结构体（Anonymous Structure）**。

```go
var employee struct {
    firstName, lastName string
    age int
}
```

上述代码片段创建一个**匿名结构体** `employee`。

### 创建命名的结构体

```go
type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {

    //creating structure using field names
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}
```

### 创建匿名结构体

```go
func main() {
    emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }

    fmt.Println("Employee 3", emp3)
}
```



### 结构体的零值（Zero Value）

因此 `firstName` 和 `lastName` 赋值为 string 的零值（`""`）。而 `age` 和 `salary` 赋值为 int 的零值（0）。

当然还可以为某些字段指定初始值，而忽略其他字段。这样，忽略的字段名会赋值为零值。

```go
type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp5 := Employee{
        firstName: "John",
        lastName:  "Paul",
    }
    fmt.Println("Employee 5", emp5)
}
```

### 访问结构体的字段

点号操作符 `.` 用于访问结构体的字段。

还可以创建零值的 `struct`，以后再给各个字段赋值。

### 结构体的指针

还可以创建指向结构体的指针。

```go
type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp8 := &Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", (*emp8).firstName)
    fmt.Println("Age:", (*emp8).age)
}
```

**Go 语言允许我们在访问 firstName 字段时，可以使用 emp8.firstName 来代替显式的解引用 (\*emp8).firstName**。

```go
type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp8 := &Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp8.firstName)
    fmt.Println("Age:", emp8.age)
}
```

### 匿名字段

当我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）。

```go
type Person struct {  
    string
    int
}

func main() {  
    p := Person{"Naveen", 50}
    fmt.Println(p)
}
```

**虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型**。比如在上面的 `Person` 结构体里，虽说字段是匿名的，但 Go 默认这些字段名是它们各自的类型。所以 `Person` 结构体有两个名为 `string` 和 `int` 的字段。

### 嵌套结构体（Nested Structs）

```go
type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age int
    address Address
}

func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.address = Address {
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:",p.age)
    fmt.Println("City:",p.address.city)
    fmt.Println("State:",p.address.state)
}
```

### 提升字段（Promoted Fields）

如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问。

`Person` 结构体有一个匿名字段 `Address`，而 `Address` 是一个结构体。现在结构体 `Address` 有 `city` 和 `state` 两个字段，访问这两个字段就像在 `Person` 里直接声明的一样，因此我们称之为提升字段。

```go
type Address struct {
    city, state string
}
type Person struct {
    name string
    age  int
    Address
}

func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.Address = Address{
        city:  "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city) //city is promoted field
    fmt.Println("State:", p.state) //state is promoted field
}
```

### 导出结构体和字段

如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。同样，如果结构体里的字段首字母大写，它也能被其他包访问到。



### 结构体相等性（Structs Equality）

**结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的**。（使用 “==” 比较）

**如果结构体包含不可比较的字段，则结构体变量也不可比较。**

结构体类型 `image` 包含一个 `map` 类型的字段。由于 `map` 类型是不可比较的，因此 `image1` 和 `image2` 也不可比较。

## 4.2 方法

方法其实就是一个函数，在 `func` 这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。

```go
func (t Type) methodName(parameter list) {
}
```

上面的代码片段创建了一个接收器类型为 `Type` 的方法 `methodName`。

### 为什么我们已经有函数了还需要方法呢？

- [Go 不是纯粹的面向对象编程语言](https://golang.org/doc/faq#Is_Go_an_object-oriented_language)，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
- 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。假设我们有一个 `Square` 和 `Circle` 结构体。可以在 `Square` 和 `Circle` 上分别定义一个 `Area` 方法。

```go
func (r Rectangle) Area() int {
    return r.length * r.width
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}
```

### 指针接收器与值接收器

到目前为止，我们只看到了使用值接收器的方法。还可以创建使用指针接收器的方法。值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的

```go
type Employee struct {
    name string
    age  int
}

/*
使用值接收器的方法。
*/
func (e Employee) changeName(newName string) {
    e.name = newName
}

/*
使用指针接收器的方法。
*/
func (e *Employee) changeAge(newAge int) {
    e.age = newAge
}

func main() {
    e := Employee{
        name: "Mark Andrew",
        age:  50,
    }
    fmt.Printf("Employee name before change: %s", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    (&e).changeAge(51)
    fmt.Printf("\nEmployee age after change: %d", e.age)
}
```

使用 `e.changeAge(51)` 来代替 `(&e).changeAge(51)`，它输出相同的结果。

### 那么什么时候使用指针接收器，什么时候使用值接收器？

一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。

指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

在其他的所有情况，值接收器都可以被使用。



### 匿名字段的方法

属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样。

```go
type address struct {
    city  string
    state string
}

func (a address) fullAddress() {
    fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
    firstName string
    lastName  string
    address
}

func main() {
    p := person{
        firstName: "Elon",
        lastName:  "Musk",
        address: address {
            city:  "Los Angeles",
            state: "California",
        },
    }

    p.fullAddress() //访问 address 结构体的 fullAddress 方法
}
```

### 在方法中使用值接收器 与 在函数中使用值参数

当一个函数有一个值参数，它只能接受一个值参数。

当一个方法有一个值接收器，它可以接受值接收器和指针接收器。

```go
type rectangle struct {
    length int
    width  int
}

func area(r rectangle) {
    fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
    fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {
    r := rectangle{
        length: 10,
        width:  5,
    }
    area(r)
    r.area()

    p := &r
    /*
       compilation error, cannot use p (type *rectangle) as type rectangle
       in argument to area
    */
    //area(p)

    p.area()//通过指针调用值接收器
}
```

函数 `func area(r rectangle)` 接受一个值参数，方法 `func (r rectangle) area()` 接受一个值接收器。

我们通过值参数 `area(r)` 来调用 area 这个函数，这是合法的。同样，我们使用值接收器来调用 area 方法 `r.area()`，这也是合法的。

我们创建了一个指向 `r` 的指针 `p`。如果我们试图把这个指针传递到只能接受一个值参数的函数 area，编译器将会报错。所以我把代码的第 33 行注释了。

代码 `p.area()` 使用指针接收器 `p` 调用了只接受一个值接收器的方法 `area`。这是完全有效的。原因是当 `area` 有一个值接收器时，为了方便Go语言把 `p.area()` 解释为 `(*p).area()`。

### 在方法中使用指针接收器 与 在函数中使用指针参数

和值参数相类似，函数使用指针参数只接受指针，而使用指针接收器的方法可以使用值接收器和指针接收器。

```go
type rectangle struct {
    length int
    width  int
}

func perimeter(r *rectangle) {
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {
    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r //pointer to r
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

    r.perimeter()//使用值来调用指针接收器
}
```



### 在非结构体上的方法

到目前为止，我们只在结构体类型上定义方法。也可以在非结构体类型上定义方法，但是有一个问题。**为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。到目前为止，我们定义的所有结构体和结构体上的方法都是在同一个main 包中，因此它们是可以运行的。**

让该程序工作的方法是为内置类型 int 创建一个类型别名，然后创建一个以该类型别名为接收器的方法。

```go
package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {
    return a + b
}

func main() {
    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}
```



# 5 面向接口

在面向对象的领域里，接口一般这样定义：**接口定义一个对象的行为**。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。

在 Go 语言中，接口就是方法签名（Method Signature）的集合。当一个类型定义了接口中的所有方法，我们称它实现了该接口。这与面向对象编程（OOP）的说法很类似。**接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法**。

例如，`WashingMachine` 是一个含有 `Cleaning()` 和 `Drying()` 两个方法的接口。任何定义了 `Cleaning()` 和 `Drying()` 的类型，都称它实现了 `WashingMachine` 接口。



### 接口的声明和实现

```go
//interface definition
type VowelsFinder interface {  
    FindVowels() []rune
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

func main() {  
    name := MyString("Sam Anderson")
    var v VowelsFinder
    v = name // possible since MyString implements VowelsFinder
    fmt.Printf("Vowels are %c", v.FindVowels())

}
```

**我们给接受者类型（Receiver Type） MyString 添加了方法 FindVowels() []rune。现在，我们称 MyString 实现了VowelsFinder 接口。这就和其他语言（如 Java）很不同，其他一些语言要求一个类使用 implement 关键字，来显式地声明该类实现了接口。而在 Go 中，并不需要这样。如果一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口**。

### 接口的实际用途

```go
type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId  int
    basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing  
the salaries of the individual employees  
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

}
```

这样做最大的优点是：`totalExpense` 可以扩展新的员工类型，而不需要修改任何代码。假如公司增加了一种新的员工类型 `Freelancer`，它有着不同的薪资结构。`Freelancer`只需传递到 `totalExpense` 的切片参数中，无需 `totalExpense` 方法本身进行修改。只要 `Freelancer` 也实现了 `SalaryCalculator` 接口，`totalExpense` 就能够实现其功能。

### 接口的内部表示

我们可以把接口看作内部的一个元组 `(type, value)`。 `type` 是接口底层的具体类型（Concrete Type），而 `value` 是具体类型的值。

```go
type Test interface {  
    Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {  
    fmt.Println(m)
}

func describe(t Test) {  
    fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {  
    var t Test
    f := MyFloat(89.7)
    t = f
    describe(t)
    t.Tester()
}
```

### 空接口

没有包含方法的接口称为空接口。空接口表示为 `interface{}`。由于空接口没有方法，因此所有类型都实现了空接口。

`describe(i interface{})` 函数接收空接口作为参数，因此，可以给这个函数传递任何类型。



### 类型断言

类型断言用于提取接口的底层值（Underlying Value）。

在语法 `i.(T)` 中，接口 `i` 的具体类型是 `T`，该语法用于获得接口的底层值。

```go
func assert(i interface{}) {  
    s := i.(int) //get the underlying int value from i
    fmt.Println(s)
}
func main() {  
    var s interface{} = 56
    assert(s)
}
```











# 6 函数式编程



# 格式说明符

 `%T` 用于打印类型，

 `%d` 用于打印数字。

`%x` 格式限定符用于指定 16 进制编码

`%c` 格式限定符用于打印字符串的字符

