package envconv

import (
	"strconv"
	"strings"
)

// convertBool converts the passed string to a bool. It will
// also return an error, if applicable.
func convertBool(value string) (bool, error) {
	convertedValue, err := strconv.ParseBool(strings.ToLower(value))
	if err != nil {
		return false, err
	}
	return convertedValue, nil
}

// ToBool returns the value of the requested environment variable
// converted to a boolean. An error will be returned if the
// environment variable is not found or the conversion to
// boolean fails.
func ToBool(varName string) (bool, error) {
	value, err := loadFromEnvironment(varName, true)
	if err != nil {
		return false, err
	}

	return convertBool(value)
}

// ToBoolSlice returns the value of the requested environment variable
// converted to a slice of bools. An error will be returned if the
// environment variable is not found or the conversion to
// slice of bools fails.
func ToBoolSlice(varName string, separator string) ([]bool, error) {
	value, err := loadFromEnvironment(varName, true)
	if err != nil {
		return []bool{}, err
	}

	boolStrings := strings.Split(value, separator)
	var bools = []bool{}
	for _, b := range boolStrings {
		convertedBool, err := convertBool(b)
		if err != nil {
			return []bool{}, err
		}
		bools = append(bools, convertedBool)
	}
	return bools, nil
}

// ToBoolWithDefault returns the value of the requested environment
// variable converted to a boolean. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to boolean fails.
func ToBoolWithDefault(varName string, defaultValue bool) bool {
	value, err := ToBool(varName)
	if err != nil {
		return defaultValue
	}
	return value
}

// ToBoolSliceWithDefault returns the value of the requested environment
// variable converted to a slice of time.Bools. The default value
// passed as the second parameter will be returned if the environment
// variable is not found or the conversion to time.Bool fails.
func ToBoolSliceWithDefault(varName string, separator string, defaultValue []bool) []bool {
	value, err := ToBoolSlice(varName, separator)
	if err != nil {
		return defaultValue
	}
	return value
}
