// Copyright 2017, maizuo.andre. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package util

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// EXPORT FUNCTION
//
// Struct to string
func Struct2String(param interface{}, pretty ...bool) string {
	body, _ := json.Marshal(param)
	if pretty != nil && len(pretty) > 0 && pretty[0] {
		var buff bytes.Buffer
		json.Indent(&buff, body, "", "\t")
		return buff.String()
	}
	return string(body)
}

// String2Struct
func String2Struct(body []byte, param interface{}) error {
	if err := json.Unmarshal(body, param); err != nil {
		return err
	}
	return nil
}

// String2uint32,conver string to uint32
func String2Uint32(str string) uint32 {
	value, _ := strconv.Atoi(str)
	return uint32(value)
}

// String2Int32
func String2Int32(str string) int32 {
	value, _ := strconv.Atoi(str)
	return int32(value)
}

// String2Int64
func String2Int64(str string) int64 {
	value, _ := strconv.ParseInt(str, 10, 64)
	return value
}

// Uint322String
func Uint322String(value uint32) string {
	return strconv.Itoa(int(value))
}

// GetListPages
func GetListPages(total, pageSize uint32) uint32 {
	pages := total / pageSize
	if total%pageSize > 0 {
		pages += 1
	}
	return pages
}
