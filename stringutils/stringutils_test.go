package stringutils

import (
	"testing"
)

var forwardStrings = [...]string{`Hello`, `The quick bròwn 狐 jumped over the lazy 犬`, `1234567890`, `~!@#$%^&*()_+|\`, `ápplesareágoodfruit`}
var backwardStrings = [...]string{`olleH`, `犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT`, `0987654321`, `\|+_)(*&^%$#@!~`, `tiurfdoogáeraselppá`}

// FIXME: Fails for combining characters (normalize input)
func TestReverse(t *testing.T) {
	// iterate through the forwardStrings array
	for i, s := range forwardStrings {
		// reverse the string
		rev := Reverse(s)
		// check if the reversed string matches its mirror array
		if rev == backwardStrings[i] {
			t.Logf("success: '%s'", s)
		} else {
			t.Errorf("fail: forward:'%s' != rev:'%s'", s, rev)
		}
	}
}

const benchmarkString string = "The quick brown fox jumped over the lazy dog."

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(benchmarkString)
	}
}
