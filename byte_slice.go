package envconv

// ToByteSlice returns the value of the requested environment variable
// converted to a byte slice. An error will be returned if the
// environment variable is not found.
func ToByteSlice(varName string) ([]byte, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return []byte{}, err
	}

	return []byte(value), nil
}

// ToByteSliceWithDefault returns the value of the requested environment
// variable converted to a byte slice. The default value passed as
// the second parameter will be returned if the environment
// variable is not found.
func ToByteSliceWithDefault(varName string, defaultValue []byte) []byte {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	return []byte(value)
}
