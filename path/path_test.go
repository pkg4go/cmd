package path

import "github.com/pkg4go/pkgs/assert"
import "testing"
import "strings"

func TestResolve(t *testing.T) {
	assert.Equal(strings.HasSuffix(Resolve("", "c"), "path/c"), true)
	assert.Equal(Resolve("/a/b", "./c"), "/a/b/c")
	assert.Equal(Resolve("/a/b", "../c"), "/a/c")
	assert.Equal(Resolve("/a/b", "c"), "/a/b/c")
	assert.Equal(Resolve("/a/b", "/c"), "/c")
	assert.Equal(Resolve("", "/c"), "/c")
}
