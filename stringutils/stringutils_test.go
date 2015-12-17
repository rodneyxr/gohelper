package stringutils

import (
	"testing"
)

var regexTestMap = map[string]string{
	`^a.*e$`: "apple",
	`50`:     "That costs $50!",
}

func TestMatches(t *testing.T) {
	for k, v := range regexTestMap {
		if matched, err := Matches(k, v); err != nil {
			t.Errorf("fail: '%s':'%s' should not be an error.", k, v)
		} else if !matched {
			t.Errorf("fail: Matches(\"%s\", \"%s\") returned false", k, v)
		}
	}
	if _, err := Matches(`\`, "error test"); err == nil {
		t.Error("fail: unescaped backslash should result in an error")
	}
}

var reverseTestMap = map[string]string{
	`Hello`:               `olleH`,
	`1234567890`:          `0987654321`,
	`~!@#$%^&*()_+|\`:     `\|+_)(*&^%$#@!~`,
	`ápplesareágoodfruit`: `tiurfdoogáeraselppá`,
	`asdf`:                `fdsa`,
}

func TestReverse(t *testing.T) {
	// iterate through the forwardStrings array
	for s, expected := range reverseTestMap {
		// reverse the string
		rev := Reverse(s)
		// check if the reversed string matches its mirror array
		if rev != expected {
			t.Errorf("fail: forward:'%s' != reverse:'%s'", s, rev)
		}
	}
}

var reverseCCTestMap = map[string]string{
	`The quick bròwn 狐 jumped over the lazy 犬`: `犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT`,
	`as⃝df̅`:              `f̅ds⃝a`,
	`ápplesareágoodfruit`: `tiurfdoogáeraselppá`,
}

func TestReverseCC(t *testing.T) {
	// iterate through the forwardStrings array
	for s, expected := range reverseCCTestMap {
		// reverse the string
		rev := ReverseCC(s)
		// check if the reversed string matches its mirror array
		if rev != expected {
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
