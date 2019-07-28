
## Introduction

关于 Go 测试，我们应该知道测试方式（或者说测试手段）、测试对象及测试原因。

## How 测试方式

### 测试实现

举个例子。针对字符串分割函数（如下），实现单元测试。
```go
package goTest

import "strings"

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}
```
在当前目录下且一样的包名 goTest ，写一个简单的 go 测试函数，如下：
```go
package goTest

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
```
测试函数必须以 Test 开头， 且必须携带一个 \*testing.T 参数。 t \*testing.T 提供改测试函数的打印、跳过、失败功能。

### 测试执行
<!--more-->

当前目录下，执行 go test ，输出如下：
```bash
> go test
PASS
ok      goTest  0.005s
```
如果项目中存在多个 package ，若要执行所有包的测试可以在项目根目录下使用 go test ./... ，输出如下（例子：github.com/mattn/go-sqlite3）：
```bash
> go test ./...
ok  	github.com/mattn/go-sqlite3	14.693s
?   	github.com/mattn/go-sqlite3/upgrade	[no test files]
```

### 代码测试覆盖率

还是以字符串分割函数为例， 获取当前代码测试覆盖率方式如下：
```bash
> go test -coverprofile=c.out
PASS
coverage: 100.0% of statements
ok  	goTest	0.005s
```
数据显示覆盖率为 100% 。若要以 HTML 方式显示可以使用命令 **go tool cover -html=c.out** 。

【tip】 一行命令 cover 获取当前目录下的代码测试覆盖度。 在 ~/.bashrc 中添加如下命令：
```bash
cover () {
    local t=$(mktemp -t cover)
    go test $COVERFLAGS -coverprofile=$t $@ \
      && go tool cover -func=$t \
      && unlink $t
}
```
执行后获取的测试覆盖度结果如下：
```bash
> cover
PASS
coverage: 100.0% of statements
ok  	goTest	0.008s
goTest/wwg_split.go:7:	Split		 100.0%
total:			(statements)     100.0%
```

问题：**测试覆盖率 100% ，结束了？**

多个测试用例的情况下，使用表组测试用例装填。更改 TestSplit 如下：
```go
func TestSplit(t *testing.T) {
	tests := []struct{
		input string
		sep string
		want []string
	}{
	    {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
	    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
        {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},  // trailing sep
        {input: "abc", sep: "/", want: []string{"abc"}},
	}

	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
```
增加测试用例 trailing sep 后，执行测试，结果如下：
```bash
> go test
--- FAIL: TestSplit (0.00s)
    wwg_split_test.go:23: expected: [a b c], got: [a b c ]
FAIL
exit status 1
FAIL	goTest	0.005s
```
根据该结果很难一下子在表组测试用例中查出是哪条。可以将 **表组测试用例实现改为 map 形式** ，具体如下：
```go
func TestSplit(t *testing.T) {
	tests := map[string]struct{
		input string
		sep string
		want []string
	}{
		"simple": {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		"no sep": {input: "abc", sep: "/", want: []string{"abc"}},
	}

	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s expected: %v, got: %v", name, tc.want, got)
		}
	}
}
```
执行测试结果如下：
```bash
> go test
--- FAIL: TestSplit (0.00s)
    wwg_split_test.go:23: trailing sep expected: [a b c], got: [a b c ]
FAIL
exit status 1
FAIL	goTest	0.005s
```

Sub tests 使用，及 '%#v' format 使用，更改 TestSplit 如下：
```go
func TestSplit(t *testing.T) {
	tests := map[string]struct{
		input string
		sep string
		want []string
	}{
		"simple": {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		"no sep": {input: "abc", sep: "/", want: []string{"abc"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s expected: %#v, got: %#v", name, tc.want, got)
			}
		})
	}
}
```
测试结果如下：
```bash
> go test
--- FAIL: TestSplit (0.00s)
    --- FAIL: TestSplit/trailing_sep (0.00s)
        wwg_split_test.go:24: trailing sep expected: []string{"a", "b", "c"}, got: []string{"a", "b", "c", ""}
FAIL
exit status 1
FAIL	goTest	0.005s
```
更好的打印格式，可以访问：
* https://github.com/k0kubun/pp
* https://github.com/davecgh/go-spew
* https://github.com/google/go-cmp

使用 google/go-cmp 优化打印， 更改 TestSplit 如下：
```go
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
            diff := cmp.Diff(tc.want, got)
            if diff != "" {
                t.Fatalf(diff)
            }
		})
	}
```
执行测试结果如下：
```bash
> go test
--- FAIL: TestSplit (0.00s)
    --- FAIL: TestSplit/trailing_sep (0.00s)
        wwg_split_test.go:29:   []string{
              	"a",
              	"b",
              	"c",
            + 	"",
              }
FAIL
exit status 1
FAIL	goTest	0.005s
```

修复bug后 Split 代码如下：
```go
// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	**if len(s) > 0 {**
		result = append(result, s)
	}
	return result**
}
```
执行测试，结果如下：
```bash
> go test
PASS
ok  	goTest	0.006s

> cover
PASS
coverage: 100.0% of statements
ok  	goTest	0.006s
goTest/wwg_split.go:7:	Split		100.0%
total:                  (statements)	100.0%
```

## What 测试对象

Q_1：Go 应该测试所有因子吗？<br>
A_1：显然不是。

Q_2：何时编写测试? 1.编码完成后？ 2.编码前？ 3.其他人遍写测试，像QA、TE？ 4.项目设计人员编写测试？<br>
A_2：编码的同时编写测试代码（TDD）Article [TheThreeRulesOfTdd](http://butunclebob.com/ArticleS.UncleBob.TheThreeRulesOfTdd)

Q_3：C 单元测试对象是 function ，Java 单元测试对象是 Class ，类内部的方法， Go 的单元测试对象是？<br>
A_3：package 。测试行为，而非实施。 "The public API of a package declare this is **what**(行为) I do, not this is **how**(实施) I do it."

## Why 测试原因

即使你不做代码测试，别人也会做。自己发现 issues 总比别人发现来得好，不是吗？<br>
1. 大部分的测试(自动化)应该是开发人员自己做。
2. 手工测试不应该是你测试的主体部分，因为手工测试的复杂度为O(n)
3. 测试可以确保您始终可以运送主分支
4. 测试确定软件行为（做什么、不做什么）
5. 测试让你有信心修改他人的代码

## 总结

* You should write tests.
* You should write tests at the same time as you write your code.
Each Go package is a self contained unit.
* Your tests should assert the observable behaviour of your package, not its implementation.
* You should design your packages around their behaviour, not their implementation.



【注】部分资料源于GopherChina 2019 - 'How to write testable code'