package util

import (
	"math/rand"
	"slices"
)

func InArray[T comparable](element T, slice []T) bool {
	return slices.Contains(slice, element)
}

func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveEmpty[T string](sliceList []T) []T {
	list := []T{}
	for _, item := range sliceList {
		if item != "" {
			list = append(list, item)
		}
	}
	return list
}

// func (i *Util)  to ternary
func Ternary[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func RandomInSlice[T any](list []T) T {
	return list[rand.Intn(len(list))]
}
