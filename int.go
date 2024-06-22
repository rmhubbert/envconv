package envconv

import "strconv"

// Type intType is a convenience interface to wrap all of the possible int types
type IntType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// toIntType returns the value of the requested environment variable
// converted to type T. An error will be returned if the
// environment variable is not found or the conversion to
// type T fails.
func toIntType[T IntType, RT int64 | uint64](varName string, bitSize int, conversionFunc func(string, int, int) (RT, error)) (T, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return T(0), err
	}

	convertedValue, err := conversionFunc(value, 10, bitSize)
	if err != nil {
		return T(0), err
	}

	return T(convertedValue), err
}

// TointTypeWithDefault returns the value of the requested environment
// variable converted to type T. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to typeT fails.
func toIntTypeWithDefault[T IntType, RT int64 | uint64](varName string, defaultValue T, bitSize int, conversionFunc func(string, int, int) (RT, error)) T {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return defaultValue
	}

	convertedValue, err := conversionFunc(value, 10, bitSize)
	if err != nil {
		return defaultValue
	}

	return T(convertedValue)
}

// ToInt returns the value of the requested environment variable
// converted to an int. An error will be returned if the
// environment variable is not found or the conversion to
// int fails.
func ToInt(varName string) (int, error) {
	return toIntType[int](varName, 64, strconv.ParseInt)
}

// ToIntWithDefault returns the value of the requested environment
// variable converted to an int. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int fails.
func ToIntWithDefault(varName string, defaultValue int) int {
	return toIntTypeWithDefault[int](varName, defaultValue, 64, strconv.ParseInt)
}

// ToInt8 returns the value of the requested environment variable
// converted to an int8. An error will be returned if the
// environment variable is not found or the conversion to
// int8 fails.
func ToInt8(varName string) (int8, error) {
	return toIntType[int8](varName, 8, strconv.ParseInt)
}

// ToInt8WithDefault returns the value of the requested environment
// variable converted to an int8. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int8 fails.
func ToInt8WithDefault(varName string, defaultValue int8) int8 {
	return toIntTypeWithDefault[int8](varName, defaultValue, 8, strconv.ParseInt)
}

// ToInt16 returns the value of the requested environment variable
// converted to an int16. An error will be returned if the
// environment variable is not found or the conversion to
// int16 fails.
func ToInt16(varName string) (int16, error) {
	return toIntType[int16](varName, 16, strconv.ParseInt)
}

// ToInt16WithDefault returns the value of the requested environment
// variable converted to an int16. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int16 fails.
func ToInt16WithDefault(varName string, defaultValue int16) int16 {
	return toIntTypeWithDefault[int16](varName, defaultValue, 16, strconv.ParseInt)
}

// ToInt32 returns the value of the requested environment variable
// converted to an int32. An error will be returned if the
// environment variable is not found or the conversion to
// int32 fails.
func ToInt32(varName string) (int32, error) {
	return toIntType[int32](varName, 32, strconv.ParseInt)
}

// ToInt32WithDefault returns the value of the requested environment
// variable converted to an int32. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int32 fails.
func ToInt32WithDefault(varName string, defaultValue int32) int32 {
	return toIntTypeWithDefault[int32](varName, defaultValue, 32, strconv.ParseInt)
}

// ToInt64 returns the value of the requested environment variable
// converted to an int64. An error will be returned if the
// environment variable is not found or the conversion to
// int64 fails.
func ToInt64(varName string) (int64, error) {
	return toIntType[int64](varName, 64, strconv.ParseInt)
}

// ToInt64WithDefault returns the value of the requested environment
// variable converted to an int64. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int64 fails.
func ToInt64WithDefault(varName string, defaultValue int64) int64 {
	return toIntTypeWithDefault[int64](varName, defaultValue, 64, strconv.ParseInt)
}

// ToUInt returns the value of the requested environment variable
// converted to an int. An error will be returned if the
// environment variable is not found or the conversion to
// int fails.
func ToUInt(varName string) (uint, error) {
	return toIntType[uint](varName, 64, strconv.ParseUint)
}

// ToUIntWithDefault returns the value of the requested environment
// variable converted to an int. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int fails.
func ToUIntWithDefault(varName string, defaultValue uint) uint {
	return toIntTypeWithDefault[uint](varName, defaultValue, 64, strconv.ParseUint)
}

// ToUInt8 returns the value of the requested environment variable
// converted to an uint8. An error will be returned if the
// environment variable is not found or the conversion to
// uint8 fails.
func ToUInt8(varName string) (uint8, error) {
	return toIntType[uint8](varName, 8, strconv.ParseUint)
}

// ToUInt8WithDefault returns the value of the requested environment
// variable converted to an uint8. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint8 fails.
func ToUInt8WithDefault(varName string, defaultValue uint8) uint8 {
	return toIntTypeWithDefault[uint8](varName, defaultValue, 8, strconv.ParseUint)
}

// ToUInt16 returns the value of the requested environment variable
// converted to an uint16. An error will be returned if the
// environment variable is not found or the conversion to
// uint16 fails.
func ToUInt16(varName string) (uint16, error) {
	return toIntType[uint16](varName, 16, strconv.ParseUint)
}

// ToUInt16WithDefault returns the value of the requested environment
// variable converted to an uint16. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint16 fails.
func ToUInt16WithDefault(varName string, defaultValue uint16) uint16 {
	return toIntTypeWithDefault[uint16](varName, defaultValue, 16, strconv.ParseUint)
}

// ToUInt32 returns the value of the requested environment variable
// converted to an uint32. An error will be returned if the
// environment variable is not found or the conversion to
// uint32 fails.
func ToUInt32(varName string) (uint32, error) {
	return toIntType[uint32](varName, 32, strconv.ParseUint)
}

// ToUInt32WithDefault returns the value of the requested environment
// variable converted to an uint32. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint32 fails.
func ToUInt32WithDefault(varName string, defaultValue uint32) uint32 {
	return toIntTypeWithDefault[uint32](varName, defaultValue, 32, strconv.ParseUint)
}

// ToUInt64 returns the value of the requested environment variable
// converted to an uint64. An error will be returned if the
// environment variable is not found or the conversion to
// uint64 fails.
func ToUInt64(varName string) (uint64, error) {
	return toIntType[uint64](varName, 64, strconv.ParseUint)
}

// ToUInt64WithDefault returns the value of the requested environment
// variable converted to an uint64. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint64 fails.
func ToUInt64WithDefault(varName string, defaultValue uint64) uint64 {
	return toIntTypeWithDefault[uint64](varName, defaultValue, 64, strconv.ParseUint)
}
