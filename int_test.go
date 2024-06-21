package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

// Tests for int type
type ToIntTest struct {
	env         string
	value       string
	expected    int
	errExpected bool
}

type ToIntWithDefaultTest struct {
	env          string
	value        string
	expected     int
	defaultValue int
}

func TestToInt(t *testing.T) {
	testData := []ToIntTest{
		{"TEST_INT_0", "0", 0, false},
		{"TEST_INT_9223372036854775807", "9223372036854775807", 9223372036854775807, false},
		{"TEST_INT_9223372036854775808", "9223372036854775808", 0, true},
		{"TEST_INT_-9223372036854775808", "-9223372036854775808", -9223372036854775808, false},
		{"TEST_INT_-9223372036854775809", "-9223372036854775809", 0, true},
		{"TEST_INT_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToInt(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToInt("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, 0, v, "they should be equal")
	})
}

func TestToIntWithDefault(t *testing.T) {
	testData := []ToIntWithDefaultTest{
		{"TEST_INT_WITH_DEFAULT_0", "0", 0, 0},
		{"TEST_INT_WITH_DEFAULT_9223372036854775807", "9223372036854775807", 9223372036854775807, 9223372036854775807},
		{"TEST_INT_WITH_DEFAULT_9223372036854775808", "9223372036854775808", 0, 0},
		{"TEST_INT_WITH_DEFAULT_-9223372036854775808", "-9223372036854775808", -9223372036854775808, -9223372036854775808},
		{"TEST_INT_WITH_DEFAULT_-9223372036854775809", "-9223372036854775809", 0, 0},
		{"TEST_INT_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 0},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToIntWithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToIntWithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, int(105), v, "they should be equal")
	})
}

// Tests for int8 type
type ToInt8Test struct {
	env         string
	value       string
	expected    int8
	errExpected bool
}

type ToInt8WithDefaultTest struct {
	env          string
	value        string
	expected     int8
	defaultValue int8
}

func TestToInt8(t *testing.T) {
	testData := []ToInt8Test{
		{"TEST_INT8_0", "0", 0, false},
		{"TEST_INT8_127", "127", 127, false},
		{"TEST_INT8_128", "128", 0, true},
		{"TEST_INT8_-128", "-128", -128, false},
		{"TEST_INT8_-129", "-129", 0, true},
		{"TEST_INT8_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToInt8(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToInt8("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, int8(0), v, "they should be equal")
	})
}

func TestToInt8WithDefault(t *testing.T) {
	testData := []ToInt8WithDefaultTest{
		{"TEST_INT8_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_INT8_WITH_DEFAULT_127", "127", 127, 127},
		{"TEST_INT8_WITH_DEFAULT_128", "128", 0, 0},
		{"TEST_INT8_WITH_DEFAULT_-128", "-128", -128, 105},
		{"TEST_INT8_WITH_DEFAULT_-129", "-129", 0, 0},
		{"TEST_INT8_WITH_DEFAULT_NOTANUMBER", "notanumber", 105, 105},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToInt8WithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToInt8WithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, int8(105), v, "they should be equal")
	})
}

// Tests for int16 type
type ToInt16Test struct {
	env         string
	value       string
	expected    int16
	errExpected bool
}

type ToInt16WithDefaultTest struct {
	env          string
	value        string
	expected     int16
	defaultValue int16
}

func TestToInt16(t *testing.T) {
	testData := []ToInt16Test{
		{"TEST_INT16_0", "0", 0, false},
		{"TEST_INT16_32767", "32767", 32767, false},
		{"TEST_INT16_32768", "32768", 0, true},
		{"TEST_INT16_-32768", "-32768", -32768, false},
		{"TEST_INT16_-32769", "-32769", 0, true},
		{"TEST_INT16_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToInt16(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToInt16("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, int16(0), v, "they should be equal")
	})
}

func TestToInt16WithDefault(t *testing.T) {
	testData := []ToInt16WithDefaultTest{
		{"TEST_INT16_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_INT16_WITH_DEFAULT_32767", "32767", 32767, 32767},
		{"TEST_INT16_WITH_DEFAULT_32768", "32768", 0, 0},
		{"TEST_INT16_WITH_DEFAULT_-32768", "-32767", -32767, -32767},
		{"TEST_INT16_WITH_DEFAULT_-32769", "-32769", 0, 0},
		{"TEST_INT16_WITH_DEFAULT_NOTANUMBER", "notanumber", 105, 105},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToInt16WithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToInt16WithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, int16(105), v, "they should be equal")
	})
}

// Tests for int32 type
type ToInt32Test struct {
	env         string
	value       string
	expected    int32
	errExpected bool
}

type ToInt32WithDefaultTest struct {
	env          string
	value        string
	expected     int32
	defaultValue int32
}

func TestToInt32(t *testing.T) {
	testData := []ToInt32Test{
		{"TEST_INT32_0", "0", 0, false},
		{"TEST_INT32_2147483647", "2147483647", 2147483647, false},
		{"TEST_INT32_2147483648", "2147483648", 0, true},
		{"TEST_INT32_-2147483648", "-2147483648", -2147483648, false},
		{"TEST_INT32_-2147483649", "-2147483649", 0, true},
		{"TEST_INT32_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToInt32(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToInt32("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, int32(0), v, "they should be equal")
	})
}

func TestToInt32WithDefault(t *testing.T) {
	testData := []ToInt32WithDefaultTest{
		{"TEST_INT32_WITH_DEFAULT_0", "0", 0, 0},
		{"TEST_INT32_WITH_DEFAULT_2147483647", "2147483647", 2147483647, 2147483647},
		{"TEST_INT32_WITH_DEFAULT_2147483648", "2147483648", 0, 0},
		{"TEST_INT32_WITH_DEFAULT_-2147483648", "-2147483648", -2147483648, -2147483648},
		{"TEST_INT32_WITH_DEFAULT_-2147483649", "-2147483649", 0, 0},
		{"TEST_INT32_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 0},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToInt32WithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToInt32WithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, int32(105), v, "they should be equal")
	})
}

// Tests for int64 type
type ToInt64Test struct {
	env         string
	value       string
	expected    int64
	errExpected bool
}

type ToInt64WithDefaultTest struct {
	env          string
	value        string
	expected     int64
	defaultValue int64
}

func TestToInt64(t *testing.T) {
	testData := []ToInt64Test{
		{"TEST_INT64_0", "0", 0, false},
		{"TEST_INT64_9223372036854775807", "9223372036854775807", 9223372036854775807, false},
		{"TEST_INT64_9223372036854775808", "9223372036854775808", 0, true},
		{"TEST_INT64_-9223372036854775808", "-9223372036854775808", -9223372036854775808, false},
		{"TEST_INT64_-9223372036854775809", "-9223372036854775809", 0, true},
		{"TEST_INT64_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToInt64(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToInt64("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, int64(0), v, "they should be equal")
	})
}

func TestToInt64WithDefault(t *testing.T) {
	testData := []ToInt64WithDefaultTest{
		{"TEST_INT64_WITH_DEFAULT_0", "0", 0, 0},
		{"TEST_INT64_WITH_DEFAULT_9223372036854775807", "9223372036854775807", 9223372036854775807, 9223372036854775807},
		{"TEST_INT64_WITH_DEFAULT_9223372036854775808", "9223372036854775808", 0, 0},
		{"TEST_INT64_WITH_DEFAULT_-9223372036854775808", "-9223372036854775808", -9223372036854775808, -9223372036854775808},
		{"TEST_INT64_WITH_DEFAULT_-9223372036854775809", "-9223372036854775809", 0, 0},
		{"TEST_INT64_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 0},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToInt64WithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToInt64WithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, int64(105), v, "they should be equal")
	})
}

// Tests for uint type
type ToUIntTest struct {
	env         string
	value       string
	expected    uint
	errExpected bool
}

type ToUIntWithDefaultTest struct {
	env          string
	value        string
	expected     uint
	defaultValue uint
}

func TestToUInt(t *testing.T) {
	testData := []ToUIntTest{
		{"TEST_UINT_0", "0", 0, false},
		{"TEST_UINT_-1", "-1", 0, true},
		{"TEST_UINT_18446744073709551615", "18446744073709551615", 18446744073709551615, false},
		{"TEST_UINT_18446744073709551616", "18446744073709551616", 0, true},
		{"TEST_UINT_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToUInt(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToUInt("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, uint(0), v, "they should be equal")
	})
}

func TestToUIntWithDefault(t *testing.T) {
	testData := []ToUIntWithDefaultTest{
		{"TEST_UINT_WITH_DEFAULT_0", "0", 0, 0},
		{"TEST_UINT_WITH_DEFAULT_18446744073709551615", "18446744073709551615", 18446744073709551615, 18446744073709551615},
		{"TEST_UINT_WITH_DEFAULT_18446744073709551616", "18446744073709551616", 0, 0},
		{"TEST_UINT_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 0},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToUIntWithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToUIntWithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, uint(105), v, "they should be equal")
	})
}

// Tests for uint8 type
type ToUInt8Test struct {
	env         string
	value       string
	expected    uint8
	errExpected bool
}

type ToUInt8WithDefaultTest struct {
	env          string
	value        string
	expected     uint8
	defaultValue uint8
}

func TestToUInt8(t *testing.T) {
	testData := []ToUInt8Test{
		{"TEST_UINT_0", "0", 0, false},
		{"TEST_UINT_-1", "-1", 0, true},
		{"TEST_UINT_255", "255", 255, false},
		{"TEST_UINT_256", "256", 0, true},
		{"TEST_UINT_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToUInt8(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToUInt8("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, uint8(0), v, "they should be equal")
	})
}

func TestToUInt8WithDefault(t *testing.T) {
	testData := []ToUInt8WithDefaultTest{
		{"TEST_UINT_WITH_DEFAULT_0", "0", 0, 0},
		{"TEST_UINT_WITH_DEFAULT_255", "255", 255, 255},
		{"TEST_UINT_WITH_DEFAULT_256", "256", 0, 0},
		{"TEST_UINT_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 0},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToUInt8WithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToUInt8WithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, uint8(105), v, "they should be equal")
	})
}

// Tests for uint16 type
type ToUInt16Test struct {
	env         string
	value       string
	expected    uint16
	errExpected bool
}

type ToUInt16WithDefaultTest struct {
	env          string
	value        string
	expected     uint16
	defaultValue uint16
}

func TestToUInt16(t *testing.T) {
	testData := []ToUInt16Test{
		{"TEST_UINT16_0", "0", 0, false},
		{"TEST_UINT16_-1", "-1", 0, true},
		{"TEST_UINT16_65535", "65535", 65535, false},
		{"TEST_UINT16_65536", "65536", 0, true},
		{"TEST_UINT16_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToUInt16(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToUInt16("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, uint16(0), v, "they should be equal")
	})
}

func TestToUInt16WithDefault(t *testing.T) {
	testData := []ToUInt16WithDefaultTest{
		{"TEST_UINT16_WITH_DEFAULT_0", "0", 0, 0},
		{"TEST_UINT16_WITH_DEFAULT_65535", "65535", 65535, 65535},
		{"TEST_UINT16_WITH_DEFAULT_65536", "65536", 0, 0},
		{"TEST_UINT16_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 0},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToUInt16WithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToUInt16WithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, uint16(105), v, "they should be equal")
	})
}

// Tests for uint32 type
type ToUInt32Test struct {
	env         string
	value       string
	expected    uint32
	errExpected bool
}

type ToUInt32WithDefaultTest struct {
	env          string
	value        string
	expected     uint32
	defaultValue uint32
}

func TestToUInt32(t *testing.T) {
	testData := []ToUInt32Test{
		{"TEST_UINT32_0", "0", 0, false},
		{"TEST_UINT32_-1", "-1", 0, true},
		{"TEST_UINT32_4294967295", "4294967295", 4294967295, false},
		{"TEST_UINT32_4294967296", "4294967296", 0, true},
		{"TEST_UINT32_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToUInt32(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToUInt32("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, uint32(0), v, "they should be equal")
	})
}

func TestToUInt32WithDefault(t *testing.T) {
	testData := []ToUInt32WithDefaultTest{
		{"TEST_UINT32_WITH_DEFAULT_0", "0", 0, 0},
		{"TEST_UINT32_WITH_DEFAULT_4294967295", "4294967295", 4294967295, 4294967295},
		{"TEST_UINT32_WITH_DEFAULT_4294967296", "4294967296", 0, 0},
		{"TEST_UINT32_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 0},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToUInt32WithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToUInt32WithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, uint32(105), v, "they should be equal")
	})
}

// Tests for uint64 type
type ToUInt64Test struct {
	env         string
	value       string
	expected    uint64
	errExpected bool
}

type ToUInt64WithDefaultTest struct {
	env          string
	value        string
	expected     uint64
	defaultValue uint64
}

func TestToUInt64(t *testing.T) {
	testData := []ToUInt64Test{
		{"TEST_UINT64_0", "0", 0, false},
		{"TEST_UINT64_-1", "-1", 0, true},
		{"TEST_UINT64_4294967295", "18446744073709551615", 18446744073709551615, false},
		{"TEST_UINT64_4294967296", "18446744073709551616", 0, true},
		{"TEST_UINT64_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToUInt64(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToUInt64("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, uint64(0), v, "they should be equal")
	})
}

func TestToUInt64WithDefault(t *testing.T) {
	testData := []ToUInt64WithDefaultTest{
		{"TEST_UINT64_WITH_DEFAULT_0", "0", 0, 0},
		{"TEST_UINT64_WITH_DEFAULT_18446744073709551615", "18446744073709551615", 18446744073709551615, 18446744073709551615},
		{"TEST_UINT64_WITH_DEFAULT_18446744073709551616", "18446744073709551616", 0, 0},
		{"TEST_UINT64_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 0},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToUInt64WithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToUInt64WithDefault("TEST_NON_EXISTANT", 105)
		assert.Equal(t, uint64(105), v, "they should be equal")
	})
}
