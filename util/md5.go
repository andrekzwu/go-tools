package util

import (
    "crypto/md5"
    "encoding/hex"
)

/*
md5 encoding
params: string
return: encoding string
*/
func Md5(data string) string {
    hash := md5.New()                    // init md5 object
    hash.Write([]byte(data))             // need encoding string
    cipherStr := hash.Sum(nil)           // calculate
    return hex.EncodeToString(cipherStr) // return encoding string
}
