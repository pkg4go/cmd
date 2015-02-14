package command

import "github.com/pkg4go/pkgs/assert"
import "strings"
import "testing"

func TestRun(t *testing.T) {
	out, err := Run("ls", "-a")
	assert.Equal(err, nil)
	assert.Equal(strings.Contains(out, "command.go"), true)
	assert.Equal(strings.Contains(out, "command_test.go"), true)

	out, err = Run("xxoo")
	assert.NotEqual(err, nil)
	assert.NotEqual(out, "")
}
