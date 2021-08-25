package util

import (
	"fmt"
	"strconv"
)

// StringSlice2ISlice
func StringSlice2ISlice(items []string) []interface{} {
	objs := make([]interface{}, 0, len(items))
	for _, item := range items {
		objs = append(objs, item)
	}
	return objs
}

// Uint32Slice2ISlice
func Uint32Slice2ISlice(items []uint32) []interface{} {
	objs := make([]interface{}, 0, len(items))
	for _, item := range items {
		objs = append(objs, item)
	}
	return objs
}

// Uint64Slice2ISlice
func Uint64Slice2ISlice(items []uint64) []interface{} {
	objs := make([]interface{}, 0, len(items))
	for _, item := range items {
		objs = append(objs, item)
	}
	return objs
}

// StringSlice2U32Slice
func StringSlice2U32Slice(items []string) ([]uint32, error) {
	array := make([]uint32, 0, len(items))
	for _, item := range items {
		u, err := strconv.ParseUint(item, 10, 32)
		if err != nil {
			return nil, err
		}
		array = append(array, uint32(u))
	}
	return array, nil
}

// Uint32Slice2SSlice
func Uint32Slice2SSlice(array []uint32) []string {
	if array == nil {
		return nil
	}
	strArray := make([]string, 0, len(array))
	for _, item := range array {
		strArray = append(strArray, fmt.Sprintf("%v", item))
	}
	return strArray
}

// InterfaceSlice2SSlice
func InterfaceSlice2SSlice(array []interface{}) []string {
	if array == nil {
		return nil
	}
	strArray := make([]string, 0, len(array))
	for _, item := range array {
		strArray = append(strArray, fmt.Sprintf("%v", item))
	}
	return strArray
}

type Uint32Slice []uint32

func (s Uint32Slice) Len() int {
	return len(s)
}

func (s Uint32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Uint32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}
