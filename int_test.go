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
		{"TEST_INT_105", "105", 105, false},
		{"TEST_INT_-105", "-105", -105, false},
		{"TEST_INT_0", "0", 0, false},
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
