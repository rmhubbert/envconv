package envconv_test

import (
	"os"
	"testing"
	"time"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

type ToDurationTest struct {
	env         string
	value       string
	expected    time.Duration
	errExpected bool
}

type ToDurationWithDefaultTest struct {
	env          string
	value        string
	expected     time.Duration
	defaultValue time.Duration
}

func TestToDuration(t *testing.T) {
	testData := []ToDurationTest{
		{"TEST_DURATION_1s", "1s", 1000000000, false},
		{"TEST_DURATION_1m", "1m", 60000000000, false},
		{"TEST_DURATION_1h", "1h", 3600000000000, false},
		{"TEST_DURATION_105", "105", 0, true},
		{"TEST_DURATION_NOTADURATION", "notaduration", time.Duration(0), true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToDuration(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToDuration("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, time.Duration(0), v, "they should be equal")
	})
}

func TestToDurationWithDefault(t *testing.T) {
	def, _ := time.ParseDuration("1m")
	testData := []ToDurationWithDefaultTest{
		{"TEST_DURATION_WITH_DEFAULT_1s", "1s", 1000000000, def},
		{"TEST_DURATION_WITH_DEFAULT_1m", "1m", 60000000000, def},
		{"TEST_DURATION_WITH_DEFAULT_1h", "1h", 3600000000000, def},
		{"TEST_DURATION_WITH_DEFAULT_105", "105", def, def},
		{"TEST_DURATION_WITH_DEFAULT_NOTADURATION", "notaduration", def, def},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToDurationWithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToDurationWithDefault("TEST_NON_EXISTANT", def)
		assert.Equal(t, def, v, "they should be equal")
	})
}
