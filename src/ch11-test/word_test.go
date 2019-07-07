package word

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// 如果我们真的需要停止测试, 或许是因为初始化失败或可能是早先的错误导致了后续错误等原因, 我们可以使用 t.Fatal 或 t.Fatalf 停止测试.
// 它们必须在和测试函数同一个 goroutine 内调用.

// 测试函数
func TestIsPalindrome(t *testing.T) {

	// 表格驱动的测试
	var tests = []struct{
		input string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},  // non-palindrome
		{"desserts", false},	// semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

//func TestNonPalindrome(t *testing.T) {
//	if IsPalindrome("palindrome") {
//		t.Error(`IsPalindrome(palindrome) = true`)
//	}
//}
//
//func TestFrenchPalindrome(t *testing.T) {
//	if !IsPalindrome("été") {
//		t.Error(`IsPalindrome("été") = false`)
//	}
//}
//
//func TestCanalPalindrome(t *testing.T) {
//	input := "A man, a plan, a canal: Panama"
//	if !IsPalindrome(input) {
//		t.Errorf(`IsPalindrome(%q) = false`, input)
//	}
//}

func TestRandomPalindrome(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalinrome(%q) = false", p)
		}
	}
}

/*
func TestRandomNonPalindrome(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalinrome(%q) = false", p)
		}
	}
}
*/

func TestSplit(t *testing.T) {
	var tests = []struct{
		splitStr string
		sep string
		words int
	}{
		{"a:b:c", ":", 3},
		{"a:b", ":", 2},
		{"abc", ":", 1},
	}

	for _, test := range tests {
		wordsReal := strings.Split(test.splitStr, test.sep)
		if got := len(wordsReal); got != test.words {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.splitStr, test.sep, got, test.words)
		}
	}
	//s, sep := "a:b:c", ":"
	//words := strings.Split(s, sep)
	//if got, want := len(words), 3; got != want {
	//	t.Errorf("Split(%q, %q) returned %d words, want %d",
	//		s, sep, got, want)
	//}
}

// 基准测试函数
// go test -bench=.
// BenchmarkIsPalindrome-4
// 10000000       123 ns/op      128 B/op    1 allocs/op
//
// 报告显示的每次调用IsPalindrome函数花费0.256微妙, 是执行5000000次的平均时间
//     因为基准测试驱动器并不知道没给基准测试函数运行所花的时间, 它会尝试在真正运行基准测试前先尝试用较小的N运行测试
// 来估算基准测试函数所需要的时间, 然后推断一个较大的时间保证稳定的测量结果
func BenchmarkIsPalindrome(b *testing.B) {
	fmt.Println(b.N)  // 100 -> 10000 -> 1000000 -> 10000000

	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// 通过size控制输入大小， 而非直接修改b.N来控制
//func benchmark(b *testing.B, size int) { /* ... */ }
//func Benchmark10(b *testing.B)         { benchmark(b, 10) }
//func Benchmark100(b *testing.B)        { benchmark(b, 100) }
//func Benchmark1000(b *testing.B)       { benchmark(b, 1000) }


// 示例函数
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("IsPalindrome"))
	// Output:
	// true
	// false
}
