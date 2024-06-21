package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

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

func TestToIntWithDefault(t *testing.T) {
	testData := []ToIntWithDefaultTest{
		{"TEST_INT_WITH_DEFAULT_105", "105", 105, 10},
		{"TEST_INT_WITH_DEFAULT_-105", "-105", -105, 11},
		{"TEST_INT_WITH_DEFAULT_0", "0", 0, 12},
		{"TEST_INT_WITH_DEFAULT_NOTANUMBER", "notanumber", 13, 13},
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
		assert.Equal(t, 105, v, "they should be equal")
	})
}
