package convert

import "github.com/pkg4go/pkgs/assert"
import "testing"

// convert to string

func TestConvertStringToString(t *testing.T) {
	var in string = "123456"
	result, _ := String(in)
	assert.Equal(result, "123456")
}

func TestConvertIntToString(t *testing.T) {
	var in int = 10086
	result, _ := String(in)
	assert.Equal(result, "10086")
}

func TestConvertInt32ToString(t *testing.T) {
	var in int32 = 12345
	result, _ := String(in)
	assert.Equal(result, "12345")
}

func TestConvertInt64ToString(t *testing.T) {
	var in int64 = 1258096
	result, _ := String(in)
	assert.Equal(result, "1258096")
}

func TestConvertFloat32ToString(t *testing.T) {
	var in float32 = 123.321
	result, _ := String(in, 3)
	assert.Equal(result, "123.321")
}

func TestConvertFloat64ToString(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := String(in, 7)
	assert.Equal(result, "1234567.7654321")
}

// convert to int

func TestConvertStringToInt(t *testing.T) {
	var in string = "123"
	result, _ := Int(in)
	assert.Equal(result, int(123))
}

func TestConvertIntToInt(t *testing.T) {
	var in int = 10086
	result, _ := Int(in)
	assert.Equal(result, int(in))
}

func TestConvertInt32ToInt(t *testing.T) {
	var in int32 = 12345
	result, _ := Int(in)
	assert.Equal(result, int(in))
}

func TestConvertInt64ToInt(t *testing.T) {
	var in int64 = 1258096
	result, _ := Int(in)
	assert.Equal(result, int(in))
}

func TestConvertFloat32ToInt(t *testing.T) {
	var in float32 = 123.321
	result, _ := Int(in)
	assert.Equal(result, int(123))
}

func TestConvertFloat64ToInt(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Int(in)
	assert.Equal(result, int(1234567))
}

// convert to int32

func TestConvertStringToInt32(t *testing.T) {
	var in string = "123"
	result, _ := Int32(in)
	assert.Equal(result, int32(123))
}

func TestConvertIntToInt32(t *testing.T) {
	var in int = 10086
	result, _ := Int32(in)
	assert.Equal(result, int32(in))
}

func TestConvertInt32ToInt32(t *testing.T) {
	var in int32 = 12345
	result, _ := Int32(in)
	assert.Equal(result, int32(in))
}

func TestConvertInt64ToInt32(t *testing.T) {
	var in int64 = 1258096
	result, _ := Int32(in)
	assert.Equal(result, int32(in))
}

func TestConvertFloat32ToInt32(t *testing.T) {
	var in float32 = 123.321
	result, _ := Int32(in)
	assert.Equal(result, int32(123))
}

func TestConvertFloat64ToInt32(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Int32(in)
	assert.Equal(result, int32(1234567))
}

// convert to int64

func TestConvertStringToInt64(t *testing.T) {
	var in string = "123"
	result, _ := Int64(in)
	assert.Equal(result, int64(123))
}

func TestConvertIntToInt64(t *testing.T) {
	var in int = 10086
	result, _ := Int64(in)
	assert.Equal(result, int64(in))
}

func TestConvertInt32ToInt64(t *testing.T) {
	var in int32 = 12345
	result, _ := Int64(in)
	assert.Equal(result, int64(in))
}

func TestConvertInt64ToInt64(t *testing.T) {
	var in int64 = 1258096
	result, _ := Int64(in)
	assert.Equal(result, int64(in))
}

func TestConvertFloat32ToInt64(t *testing.T) {
	var in float32 = 123.321
	result, _ := Int64(in)
	assert.Equal(result, int64(123))
}

func TestConvertFloat64ToInt64(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Int64(in)
	assert.Equal(result, int64(1234567))
}

// convert to float32

func TestConvertStringToFloat32(t *testing.T) {
	var in string = "123.456"
	result, _ := Float32(in)
	assert.Equal(result, float32(123.456))
}

func TestConvertIntToFloat32(t *testing.T) {
	var in int = 10086
	result, _ := Float32(in)
	assert.Equal(result, float32(in))
}

func TestConvertInt32ToFloat32(t *testing.T) {
	var in int32 = 12345
	result, _ := Float32(in)
	assert.Equal(result, float32(in))
}

func TestConvertInt64ToFloat32(t *testing.T) {
	var in int64 = 1258096
	result, _ := Float32(in)
	assert.Equal(result, float32(in))
}

func TestConvertFloat32ToFloat32(t *testing.T) {
	var in float32 = 123.321
	result, _ := Float32(in)
	assert.Equal(result, float32(in))
}

func TestConvertFloat64ToFloat32(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Float32(in)
	assert.Equal(result, float32(in))
}

// convert to float64

func TestConvertStringToFloat64(t *testing.T) {
	var in string = "123.456"
	result, _ := Float64(in)
	assert.Equal(result, float64(123.456))
}

func TestConvertIntToFloat64(t *testing.T) {
	var in int = 10086
	result, _ := Float64(in)
	assert.Equal(result, float64(in))
}

func TestConvertInt32ToFloat64(t *testing.T) {
	var in int32 = 12345
	result, _ := Float64(in)
	assert.Equal(result, float64(in))
}

func TestConvertInt64ToFloat64(t *testing.T) {
	var in int64 = 1258096
	result, _ := Float64(in)
	assert.Equal(result, float64(in))
}

func TestConvertFloat32ToFloat64(t *testing.T) {
	var in float32 = 123.321
	result, _ := Float64(in)
	assert.Equal(result, float64(in))
}

func TestConvertFloat64ToFloat64(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Float64(in)
	assert.Equal(result, float64(in))
}
