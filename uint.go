package envconv

import "strconv"

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
