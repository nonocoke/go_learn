
go test 是一个按照一定的约定和组织的测试代码的驱动程序
在包目录内，以_test.go为后缀名的源文件并不是go build构建包的一部分, 而是go test 测试的一部分。

在*_test.go 文件中, 有三种类型的函数: 测试函数(Test*)、基准测试函数(Benchmark*)、例子函数(Example*)

go test -v
"-v" 参数用于打印每个测试函数的名字和运行时间

go test -v -run="French|Cancl"
"-run" 参数是一个正则表达式, 只有测试函数名被它正确匹配的测试函数才会被 go test 执行

go tool cover 测试覆盖率工具的help
go test -run=Palindrome -coverprofile=c.out src/ch11-test

-coverprofile参数通过插入生成钩子代码来统计覆盖率数据. 即, 在运行每个测试前, 它会修改要测试代码的副本,在每个块都设置一个bool标志变量。
当被修改后的被测试代码运行退出时, 将统计日志数据写入c.out文件, 并打印一部分执行的语句的一个总结(摘要 go test -cover)
-covermode=count参数, 将在每个代码块插入计数器而不是bool变量。在统计结果中记录每个块的执行次数, 可用于衡量哪些是被频繁执行的热点代码。

go tool cover -html=c.out
收集数据后, 运行了测试覆盖率工具, 打印了测试日志, 生成一个HTML报告


剖析程序
自动化的剖析技术是基于程序执行期间的一些抽样数据, 然后推断后面的执行状态；最终产生一个运行时间的统计数据文件
go test 工具的几种分析方式
   1. go test -cpuprofile=cpu.out
    e.g., go test -run=NONE -bench=. -benchmem -cpuprofile=cpu.log
    CPU分析文件标识了函数执行时所需要的CPU时间. 当前运行的系统线程在每个几毫秒都会遇到操作系统的中断事件, 每次中断时都会记录一个分析文件
    然后恢复正常的运行。
   2. go test -blockprofile=block.out
    堆分析记录了程序的内存使用情况. 每个内存分配操作都会出发内部平均内存分配例程, 每个521KB的内存申请都会出发一个事件。
   3. go test -memprofile=mem.out
    阻塞分析记录了goroutine最大的阻塞操作, 例如系统调用, 管道发送和接收, 还有获取锁等。分析库会记录每个goroutine被阻塞时的相关操作。

e.g.,
go test -run=NONE -bench=. -benchmem -cpuprofile=cpu.log -blockprofile=block.log -memprofile=mem.log

go tool pprof -text -nodecount=10 cpu.log  (pprof参数 分析获取到的抽样数据)

