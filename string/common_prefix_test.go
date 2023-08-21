package string

import (
	"testing"
)

func TestCommonPrefix(t *testing.T) {
	strs := []string{"cb", "a"}
	prefix := longestCommonPrefix(strs)
	t.Log(prefix)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
