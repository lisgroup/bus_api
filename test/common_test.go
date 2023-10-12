package test

import (
	"bus_api/core/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	code1 := helper.GenerateCode()
	code2 := helper.GenerateCode()
	code3 := helper.GenerateCode()
	assert.Equal(t, 4, len(code1))
	assert.Equal(t, 4, len(code2))
	assert.Equal(t, 4, len(code3))
	if code1 == code2 || code2 == code3 {
		t.Fatalf("code1==code2 || code2 == code3, code: %s", code1)
	}
	t.Logf("code1: %s, code2: %s, code3: %s", code1, code2, code3)
}
