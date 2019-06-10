package implement_strstr

import "testing"

func TestStrStr(t *testing.T) {
	if strStr("hello", "llo") != 2 {
		t.Fatalf("hello/llo excepted 2, got %d", strStr("hello", "llo"))
	}
}
