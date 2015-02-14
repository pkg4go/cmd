package assert

import "testing"

func TestEqual(t *testing.T) {
	Equal(123, 123)
	Equal("123", "123")
}

func TestNotEqual(t *testing.T) {
	NotEqual(123, 100)
	NotEqual(123, "123")
}
