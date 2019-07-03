package _go

import (
    "fmt"
    "log"
    "math"
    "math/rand"
    "runtime"
    "strings"
    "time"
)

// Uint8FormInt change the type:int to type Uint8
func Uint8FormInt(n int) (uint8, error) {
    if 0 <= n && n <= math.MaxUint8 {
        return uint8(n), nil
    }
    return 0, fmt.Errorf("%d is out of uint8 range", n)
}

func IntFromFloat64(x float64) int {
    if math.MinInt32 <= x && x <= math.MaxInt32 {
        whole, fraction := math.Modf(x)
        if fraction >= 0.5 {
            whole++
        }
        return int(whole)
    }
    panic(fmt.Sprint("%g is out of int32 range", x))
}

func random10() {
    for i := 0; i < 10; i++ {
        a := rand.Int()
        fmt.Printf("%d / ", a)
    }
    for i := 0; i < 5; i++ {
        r := rand.Intn(8)
        fmt.Printf("%d / ", r)
    }
    fmt.Println()
    timens := int64(time.Now().Nanosecond())
    rand.Seed(timens)
    for i := 0; i < 10; i++ {
        fmt.Printf("%2.2f / ", 100*rand.Float32())
    }
}

func testchar() {
    var ch int = '\u0041'
    var ch2 int = '\u03B2'
    var ch3 int = '\U00101234'

    fmt.Printf("%d - %d - %d\n", ch, ch2, ch3)  // integer
    fmt.Printf("%c - %c - %c\n", ch, ch2, ch3)  // character
    fmt.Printf("%X - %X - %X\n", ch, ch2, ch3)  // UTF-8 bytes
    fmt.Printf("%U - %U - %U\n", ch, ch2, ch3)  // UTF-8 code point

}

func teststrings(str string) {
    // str = "This is an example of a string"
    // 1. strings.HasPrefix : 判断字符串 str 是否以 prefix 开头
    fmt.Printf("T/F Does the string \"%s\" have prefix %s? \n", str, "Th")
    fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
    // strings.HasSuffix : 判断字符串 str 是否以 suffix 结尾
    fmt.Printf("T/F Does the string \"%s\" have suffix %s? \n", str, "string")
    fmt.Printf("%t\n", strings.HasSuffix(str, "string"))

    // 2. strings.Contains : 判断字符串 str 是否包含 substr
    fmt.Printf("T/F Does the string \"%s\" is containes the \"%s\"? \n", str, "string")
    fmt.Printf("%t\n", strings.Contains(str, "string"))

    // 3. 子字符串或字符在父字符串中出现的位置
    fmt.Printf("The position of \"an\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "an"))

    fmt.Printf("The position of the first instance of \"a\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "a"))
    fmt.Printf("The position of the last instance of \"a\" is: ")
    fmt.Printf("%d\n", strings.LastIndex(str, "a"))

    fmt.Printf("The position of \"lalala\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "lalala"))

    // 4. 字符串替换
    // strings.Replace(str, old, new, n) string
    fmt.Printf("%s\n", strings.Replace(str, str, "lalala", -1))

    // 5. 统计字符串出现次数
    fmt.Printf("%d\n", strings.Count(str, "a"))

    // 6. 拼接重复字符串
    fmt.Printf("%s\n", strings.Repeat(str, 2))

    // 7. 修改字符串大小写
    fmt.Printf("%s\n", strings.ToLower(str))
    fmt.Printf("%s\n", strings.ToUpper(str))

    // 8. 修剪字符串
    fmt.Printf("%s\n", strings.TrimSpace(str))
    fmt.Printf("%s\n", strings.Trim(str, "string"))
    fmt.Printf("%s\n", strings.TrimLeft(str, "Th"))
    fmt.Printf("%s\n", strings.TrimRight(str, "ng"))

    // 9. 分割字符串
    fmt.Printf("%s\n", strings.Fields(str))
    fmt.Printf("%s\n", strings.Split(str, "a")[0])

    // 10. 拼接slice到字符串
    // 11. 从字符串中读取内容
    // 12. 字符串与其它类型的转换
    // ...
}

var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."

func init() {
    if runtime.GOOS == "windows" {
        prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
    } else { //Unix-like
        prompt = fmt.Sprintf(prompt, "Ctrl+D")
    }
    fmt.Printf("%s\n", prompt)
}

func testGoto(num int) {
    i := 0
    START:
        fmt.Printf("The count is at %d\n", i)
        i++
        if i < num {
            goto START
        }
    //for i := 0; i < num; i++ {
    //    fmt.Printf("The count is at %d\n", i)
    //}
}

func forCharacter(num int) {

    // 1 使用 2层for循环
    for i := 1; i <= num; i++ {
       for j := 1; j <= i; j++ {
         fmt.Printf("G")
       }
       fmt.Println()
    }
    // 2 仅用 1层for循环，及字符串拼接
    for i := 1; i <= num; i++ {
        fmt.Printf("%s\n", strings.Repeat("G", i))
    }
}

// rectangleStars 打印 '*' 矩形
func rectangleStars(rows int, cols int) {

    // 使用defer语句记录函数的参数与返回值
    defer func() {
        log.Printf("rectangleStars(%d, %d)", rows, cols)
    }()

    for i := 0; i < cols; i++ {
        fmt.Printf("%s\n", strings.Repeat("*", rows))
    }
}

// fizzbuzz
func fizzBuzz(num int) {
    const (
        FIZZ = 3
        BUZZ = 5
        FIZZBUZZ = 15
    )
    for i := 0; i <= num; i++ {
        switch {
        case i%FIZZBUZZ == 0:
            fmt.Println("FizzBuzz")
        case i%FIZZ == 0:
            fmt.Println("Fizz")
        case i%BUZZ == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}

func min(s ...int) int {
    defer un(trace("min"))
    if len(s) == 0 {
        return 0
    }
    min := s[0]
    for _, v := range s {
        if v < min {
            min = v
        }
    }
    fmt.Println(min)
    return min
}

func trace(s string) (str string) {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}


func main() {
    // var d func() bool  // 声明一个返回值为bool类型的函数变量，这种形式一般用于回调函数，即将函数以变量的形式保存下来，在需要的时候重新
    // 调用这个函数
    var aVar = 10
    fmt.Println(aVar == 5)
    fmt.Println(aVar != 5)
    fmt.Println(aVar == 10)
    fmt.Println(aVar != 10)


    i, err := Uint8FormInt(2147483648)
    fmt.Println(i, err)
    i2, err2 := Uint8FormInt(214)
    fmt.Println(i2, err2 )

    aVar = IntFromFloat64(3.12345678912345678)
    fmt.Println(aVar)
    fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界: %s\n", runtime.GOOS)
    random10()  // 随机数
    testchar()  // 字符
    var str string = "This is an example of a string"
    teststrings(str)  // 字符串操作
    testGoto(5)  // goto用法
    forCharacter(5)  // for 循环
    rectangleStars(20, 10)
    fizzBuzz(40)  // switch 用法

    for i := 0; i < 5; i++ {
        var v int
        fmt.Printf("%d ", v)  // 0 0 0 0 0
        v = 5
    }

    x := min(1, 3, 2, 0)
    fmt.Printf("The minimum is : %d\n", x)
    slice := []int{7, 9, 3, 5, 1}
    x = min(slice...)
    fmt.Printf("The minimum in the slice is : %d\n", x)
}
