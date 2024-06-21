package envconv

import "strconv"

// ToInt returns the value of the requested environment variable
// converted to an int. An error will be returned if the
// environment variable is not found or the conversion to
// int fails.
func ToInt(varName string) (int, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}

	return convertInt(value)
}

// ToIntWithDefault returns the value of the requested environment
// variable converted to an int. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int fails.
func ToIntWithDefault(varName string, defaultValue int) int {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue int
	convertedValue, err = convertInt(value)
	if err != nil {
		return defaultValue
	}
	return convertedValue
}

// convertInt converts the passed string to an int. It will
// also return an error, if applicable.
func convertInt(value string) (int, error) {
	convertedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return convertedValue, nil
}
