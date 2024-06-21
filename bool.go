package envconv

import (
	"strconv"
	"strings"
)

// ToBool returns the value of the requested environment variable
// converted to a boolean. An error will be returned if the
// environment variable is not found or the conversion to
// boolean fails.
func ToBool(varName string) (bool, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return false, err
	}

	return convertBool(value)
}

// ToBoolWithDefault returns the value of the requested environment
// variable converted to a boolean. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to boolean fails.
func ToBoolWithDefault(varName string, defaultValue bool) bool {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue bool
	convertedValue, err = convertBool(value)
	if err != nil {
		return defaultValue
	}
	return convertedValue
}

// convertBool converts the passed string to a bool. It will
// also return an error, if applicable.
func convertBool(value string) (bool, error) {
	convertedValue, err := strconv.ParseBool(strings.ToLower(value))
	if err != nil {
		return false, err
	}
	return convertedValue, nil
}
