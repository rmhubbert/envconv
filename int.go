package envconv

import "strconv"

// convertInt converts the passed string to an int. It will
// also return an error, if applicable.
func convertInt(value string, bitSize int) (int, error) {
	convertedValue, err := strconv.ParseInt(value, 10, bitSize)
	if err != nil {
		return int(0), err
	}
	return int(convertedValue), nil
}

// convertUInt converts the passed string to an unsigned int. It will
// also return an error, if applicable.
func convertUInt(value string, bitSize int) (uint, error) {
	convertedValue, err := strconv.ParseUint(value, 10, bitSize)
	if err != nil {
		return uint(0), err
	}
	return uint(convertedValue), nil
}

// ToInt returns the value of the requested environment variable
// converted to an int. An error will be returned if the
// environment variable is not found or the conversion to
// int fails.
func ToInt(varName string) (int, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}

	return convertInt(value, 64)
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
	convertedValue, err = convertInt(value, 64)
	if err != nil {
		return defaultValue
	}
	return convertedValue
}

// ToInt8 returns the value of the requested environment variable
// converted to an int8. An error will be returned if the
// environment variable is not found or the conversion to
// int8 fails.
func ToInt8(varName string) (int8, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}

	var convertedValue int
	convertedValue, err = convertInt(value, 8)
	if err != nil {
		return 0, err
	}
	return int8(convertedValue), nil
}

// ToInt8WithDefault returns the value of the requested environment
// variable converted to an int8. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int8 fails.
func ToInt8WithDefault(varName string, defaultValue int8) int8 {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue int
	convertedValue, err = convertInt(value, 8)
	if err != nil {
		return defaultValue
	}
	return int8(convertedValue)
}

// ToInt16 returns the value of the requested environment variable
// converted to an int16. An error will be returned if the
// environment variable is not found or the conversion to
// int16 fails.
func ToInt16(varName string) (int16, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}

	var convertedValue int
	convertedValue, err = convertInt(value, 16)
	if err != nil {
		return 0, err
	}
	return int16(convertedValue), nil
}

// ToInt16WithDefault returns the value of the requested environment
// variable converted to an int16. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int16 fails.
func ToInt16WithDefault(varName string, defaultValue int16) int16 {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue int
	convertedValue, err = convertInt(value, 16)
	if err != nil {
		return defaultValue
	}
	return int16(convertedValue)
}

// ToInt32 returns the value of the requested environment variable
// converted to an int32. An error will be returned if the
// environment variable is not found or the conversion to
// int32 fails.
func ToInt32(varName string) (int32, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}

	var convertedValue int
	convertedValue, err = convertInt(value, 32)
	if err != nil {
		return 0, err
	}
	return int32(convertedValue), nil
}

// ToInt32WithDefault returns the value of the requested environment
// variable converted to an int32. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int32 fails.
func ToInt32WithDefault(varName string, defaultValue int32) int32 {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue int
	convertedValue, err = convertInt(value, 32)
	if err != nil {
		return defaultValue
	}
	return int32(convertedValue)
}

// ToInt64 returns the value of the requested environment variable
// converted to an int64. An error will be returned if the
// environment variable is not found or the conversion to
// int64 fails.
func ToInt64(varName string) (int64, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}

	var convertedValue int
	convertedValue, err = convertInt(value, 64)
	if err != nil {
		return 0, err
	}
	return int64(convertedValue), nil
}

// ToInt64WithDefault returns the value of the requested environment
// variable converted to an int64. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int64 fails.
func ToInt64WithDefault(varName string, defaultValue int64) int64 {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue int
	convertedValue, err = convertInt(value, 64)
	if err != nil {
		return defaultValue
	}
	return int64(convertedValue)
}

// ToUInt returns the value of the requested environment variable
// converted to an int. An error will be returned if the
// environment variable is not found or the conversion to
// int fails.
func ToUInt(varName string) (uint, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}

	return convertUInt(value, 64)
}

// ToUIntWithDefault returns the value of the requested environment
// variable converted to an int. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int fails.
func ToUIntWithDefault(varName string, defaultValue uint) uint {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 64)
	if err != nil {
		return defaultValue
	}
	return convertedValue
}

// ToUInt8 returns the value of the requested environment variable
// converted to an uint8. An error will be returned if the
// environment variable is not found or the conversion to
// uint8 fails.
func ToUInt8(varName string) (uint8, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return uint8(0), err
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 8)
	if err != nil {
		return uint8(0), err
	}
	return uint8(convertedValue), nil
}

// ToUInt8WithDefault returns the value of the requested environment
// variable converted to an uint8. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint8 fails.
func ToUInt8WithDefault(varName string, defaultValue uint8) uint8 {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 8)
	if err != nil {
		return defaultValue
	}
	return uint8(convertedValue)
}

// ToUInt16 returns the value of the requested environment variable
// converted to an uint16. An error will be returned if the
// environment variable is not found or the conversion to
// uint16 fails.
func ToUInt16(varName string) (uint16, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return uint16(0), err
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 16)
	if err != nil {
		return uint16(0), err
	}
	return uint16(convertedValue), nil
}

// ToUInt16WithDefault returns the value of the requested environment
// variable converted to an uint16. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint16 fails.
func ToUInt16WithDefault(varName string, defaultValue uint16) uint16 {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 16)
	if err != nil {
		return defaultValue
	}
	return uint16(convertedValue)
}

// ToUInt32 returns the value of the requested environment variable
// converted to an uint32. An error will be returned if the
// environment variable is not found or the conversion to
// uint32 fails.
func ToUInt32(varName string) (uint32, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return uint32(0), err
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 32)
	if err != nil {
		return uint32(0), err
	}
	return uint32(convertedValue), nil
}

// ToUInt32WithDefault returns the value of the requested environment
// variable converted to an uint32. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint32 fails.
func ToUInt32WithDefault(varName string, defaultValue uint32) uint32 {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 32)
	if err != nil {
		return defaultValue
	}
	return uint32(convertedValue)
}

// ToUInt64 returns the value of the requested environment variable
// converted to an uint64. An error will be returned if the
// environment variable is not found or the conversion to
// uint64 fails.
func ToUInt64(varName string) (uint64, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return uint64(0), err
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 64)
	if err != nil {
		return uint64(0), err
	}
	return uint64(convertedValue), nil
}

// ToUInt64WithDefault returns the value of the requested environment
// variable converted to an uint64. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint64 fails.
func ToUInt64WithDefault(varName string, defaultValue uint64) uint64 {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue uint
	convertedValue, err = convertUInt(value, 64)
	if err != nil {
		return defaultValue
	}
	return uint64(convertedValue)
}
