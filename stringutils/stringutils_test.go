package stringutils

import (
	"testing"
)

var forwardStrings = [...]string{`Hello`, `1234567890`, `~!@#$%^&*()_+|\`, `ápplesareágoodfruit`, `asdf`}
var reverseStrings = [...]string{`olleH`, `0987654321`, `\|+_)(*&^%$#@!~`, `tiurfdoogáeraselppá`, `fdsa`}

func TestReverse(t *testing.T) {
	// iterate through the forwardStrings array
	for i, s := range forwardStrings {
		// reverse the string
		rev := Reverse(s)
		// check if the reversed string matches its mirror array
		if rev != reverseStrings[i] {
			t.Errorf("fail: forward:'%s' != reverse:'%s'", s, rev)
		}
	}
}

var forwardCCStrings = [...]string{`The quick bròwn 狐 jumped over the lazy 犬`, `ápplesareágoodfruit`, `as⃝df̅`}
var reverseCCStrings = [...]string{`犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT`, `tiurfdoogáeraselppá`, `f̅ds⃝a`}

func TestReverseCC(t *testing.T) {
	// iterate through the forwardStrings array
	for i, s := range forwardCCStrings {
		// reverse the string
		rev := ReverseCC(s)
		// check if the reversed string matches its mirror array
		if rev != reverseCCStrings[i] {
			t.Errorf("fail: forward:'%s' != reverse:'%s'", s, rev)
		}
	}
}

const benchmarkString string = "The quick brown fox jumped over the lazy dog."

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(benchmarkString)
	}
}

func BenchmarkReverseCC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseCC(benchmarkString)
	}
}
