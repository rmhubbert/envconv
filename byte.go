package envconv

import "strconv"

// ToByteSlice returns the value of the requested environment variable
// converted to a byte. An error will be returned if the
// environment variable is not found or the conversion to
// byte fails.
func ToByte(varName string) (byte, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return byte(0), err
	}

	var convertedValue int64
	convertedValue, err = strconv.ParseInt(value, 10, 8)
	if err != nil {
		return byte(0), err
	}

	return byte(convertedValue), nil
}

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

// ToByteWithDefault returns the value of the requested environment
// variable converted to a byte. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to byte fails.
func ToByteWithDefault(varName string, defaultValue byte) byte {
	value, err := ToByte(varName)
	if err != nil {
		return defaultValue
	}
	return value
}

// ToByteSliceWithDefault returns the value of the requested environment
// variable converted to a byte slice. The default value passed as
// the second parameter will be returned if the environment
// variable is not found.
func ToByteSliceWithDefault(varName string, defaultValue []byte) []byte {
	value, err := ToByteSlice(varName)
	if err != nil {
		return defaultValue
	}
	return value
}
