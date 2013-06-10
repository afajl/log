package log

import (
	"bytes"
	"github.com/bmizerany/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	var b bytes.Buffer
	setOutput(&b)

	OutLevel = DEBUG
	Trace("should not print")
	Debug(5, 3, " cats")

	assert.Equal(t, b.String()[20:], "debug: 5 3 cats\n")

	b.Reset()

	OutLevel = WARN
	Info("info")
	Warnf("warn %v", 5)
	assert.Equal(t, b.String()[20:], "warn: warn 5\n")
}
