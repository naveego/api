package utils

import (
	"testing"
)

func TestStringSplit2(t *testing.T) {
	testStringSplit2(t, "a:b", ":", "a", "b")
	testStringSplit2(t, "ab", ":", "ab", "")
	testStringSplit2(t, "a:", ":", "a", "")
	testStringSplit2(t, ":b", ":", "", "b")
}

func testStringSplit2(t *testing.T, ab, s, a, b string) {
	a2, b2 := StringSplit2(ab, s)
	if a2 != a {
		t.Errorf("want: '%s', got: '%s'", a, a2)
	}
	if b2 != b {
		t.Errorf("want: '%s', got: '%s'", b, b2)
	}
}
