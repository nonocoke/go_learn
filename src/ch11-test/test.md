go test 是一个按照一定的约定和组织的测试代码的驱动程序。
在包目录内，以_test.go为后缀名的源文件并不是go build构建包的一部分, 而是go test 测试的一部分。

在*_test.go 文件中, 有三种类型的函数: 测试函数(Test*)、基准测试函数(Benchmark*)、例子函数(Example*)

go test -v
"-v" 参数用于打印每个测试函数的名字和运行时间

go test -v -run="French|Cancl"
"-run" 参数是一个正则表达式, 只有测试函数名被它正确匹配的测试函数才会被 go test 执行