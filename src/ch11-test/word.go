// package word provides utilities for word game
package word

import (
	"math/rand"
	"unicode"
)

/*
每个测试函数必须导入testing包, 测试函数有如下签名:
func TestName (t *testing.T) {
	// ...
}
*/


// Ispalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
// Wrong version
/*
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
 */

func IsPalindrome_0(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r){
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func IsPalindrome(s string) bool {
	// var letters []rune
	// 快的程序往往是有很少的内存分配
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r){
			letters = append(letters, unicode.ToLower(r))
		}
	}

	n := len(letters)/2
	for i := 0; i < n; i++ {
	//for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)  // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))  // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)  // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))  // random rune up to '\u0999'
		p := rune(rng.Intn(0x500))  // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = p
	}
	return string(runes)
}