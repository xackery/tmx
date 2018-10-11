package model

import (
	"fmt"
	"strconv"
)

// ToInt64 converts a string to int64
func ToInt64(source string) (value int64) {
	value, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		fmt.Println("invalid integer value:", source)
	}
	return
}

// ToBool converts a string to bool
func ToBool(source string) (value bool) {
	val, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		fmt.Println("invalid integer value:", source)
	}
	value = val > 0
	return
}

// ToFloat32 converts a string to float32
func ToFloat32(source string) (value float32) {
	val, err := strconv.ParseFloat(source, 32)
	if err != nil {
		fmt.Println("invalid float value:", source)
	}
	value = float32(val)
	return
}
