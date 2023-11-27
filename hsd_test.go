package hsd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringDistance(t *testing.T) {
	// given（前提条件）
	tests := []struct {
		name     string
		left     string
		right    string
		expected int
	}{
		{"文字列の長さが異なる場合", "ab", "a", -1},
		{"すべての文字が同じ場合", "abcd", "abcd", 0},
		{"1文字だけ異なる場合", "abcd", "abce", 1},
		{"2文字異なる場合", "abcd", "abef", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when（操作）
			result := StringDistance(tt.left, tt.right)

			// then（期待する結果）
			assert.Equal(t, tt.expected, result)
		})
	}
}
