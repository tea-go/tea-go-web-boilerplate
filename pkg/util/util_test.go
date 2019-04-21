package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStr(t *testing.T) {
	rs1 := RandStr(1, "")
	assert.Equal(t, 3, len(rs1))

	rs1 = RandStr(3, "")
	assert.Equal(t, 3, len(rs1))

	rs1 = RandStr(4, "")
	assert.Equal(t, 4, len(rs1))

	rs1 = RandStr(4, "abc")
	assert.Equal(t, false, strings.Contains(rs1, "$"))

	rs1 = RandStr(4, "strong")
	assert.Equal(t, 4, len(rs1))
}
