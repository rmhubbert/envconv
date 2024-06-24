package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

type ToBoolTest struct {
	env           string
	value         string
	expected      bool
	errorExpected bool
}

type ToBoolWithDefaultTest struct {
	env          string
	value        string
	expected     bool
	defaultValue bool
}

func TestToBool(t *testing.T) {
	testData := []ToBoolTest{
		{"TEST_BOOL_1", "1", true, false},
		{"TEST_BOOL_true", "true", true, false},
		{"TEST_BOOL_TRUE", "TRUE", true, false},
		{"TEST_BOOL_tRuE", "tRuE", true, false},
		{"TEST_BOOL_0", "0", false, false},
		{"TEST_BOOL_false", "false", false, false},
		{"TEST_BOOL_FALSE", "FALSE", false, false},
		{"TEST_BOOL_FaLsE", "FaLsE", false, false},
		{"TEST_BOOL_NOTABOOL", "notabool", false, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToBool(td.env)
			if td.errorExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToBool("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, false, v, "they should be equal")
	})
}

func TestToBoolWithDefault(t *testing.T) {
	testData := []ToBoolWithDefaultTest{
		{"TEST_BOOL_WITH_DEFAULT_1", "1", true, false},
		{"TEST_BOOL_WITH_DEFAULT_true", "true", true, false},
		{"TEST_BOOL_WITH_DEFAULT_TRUE", "TRUE", true, false},
		{"TEST_BOOL_WITH_DEFAULT_tRuE", "tRuE", true, false},
		{"TEST_BOOL_WITH_DEFAULT_0", "0", false, true},
		{"TEST_BOOL_WITH_DEFAULT_false", "false", false, true},
		{"TEST_BOOL_WITH_DEFAULT_FALSE", "FALSE", false, true},
		{"TEST_BOOL_WITH_DEFAULT_FaLsE", "FaLsE", false, true},
		{"TEST_BOOL_WITH_DEFAULT_NOTABOOL", "notabool", false, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToBoolWithDefault(td.env, td.defaultValue)
			if v != td.expected {
				assert.Equal(t, td.defaultValue, v, "they should be equal")
			} else {
				assert.Equal(t, td.expected, v, "they should be equal")
			}
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToBoolWithDefault("TEST_NON_EXISTANT", false)
		assert.Equal(t, false, v, "they should be equal")
	})
}
