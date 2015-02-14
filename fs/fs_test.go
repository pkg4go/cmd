package fs

import "github.com/pkg4go/pkgs/assert"
import "testing"

func TestCopy(t *testing.T) {
	err := Copy("fs.go", "fs.go.tmp")
	// TODO: stat, remove *.tmp
	assert.Equal(err, nil)
}
