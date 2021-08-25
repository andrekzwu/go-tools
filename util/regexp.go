package util

import (
    "regexp"
    "strings"
)

var (

    // match key work
    StrKeyWord string = `\b(select|insert|delete|from|count\(|drop table|update|truncate|asc\(|mid\(|char\(|xp_cmdshell|exec master|netlocalgroup administrators|net user|or|and)\b`
    // rKeyWord
    rKeyWord = regexp.MustCompile(StrKeyWord)
)

// IsSqlInjectionAttack
func IsSqlInjectionAttack(paramsStr string) bool {
    lowParamsStr := strings.ToLower(paramsStr)
    // match str key word
    if rKeyWord.MatchString(lowParamsStr) {
        return true
    }
    return false
}
