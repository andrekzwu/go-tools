package log

import (
    "encoding/json"
    "net"
    "reflect"
    "runtime"
    "strings"
)

const (
    CALLDEPTH = 3
)

// getLocalIP
func getLocalIP() string {
    addrs, _ := net.InterfaceAddrs()
    for _, address := range addrs {
        // check the ip is loop address?
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }

        }
    }
    return "127.0.0.1"
}

// getCaller
func formatFuncName(f string) string {
    i := strings.LastIndex(f, "/")
    j := strings.Index(f[i+1:], ".")
    if j < 1 {
        return "???"
    }
    fun := f[i+j+2:]
    i = strings.LastIndex(fun, ".")
    return fun[i+1:]
}

// formatFile
func formatFile(f string) string {
    i := strings.LastIndex(f, "/")
    return f[i+1:]
}

// getCaller
func getCaller(calldepth int) (string, int, string) {
    pc, file, line, ok := runtime.Caller(calldepth)
    if !ok {
        return "???", 0, "???"
    }
    return formatFile(file), line, formatFuncName(runtime.FuncForPC(pc).Name())
}

// struct to string
func structToString(param interface{}) string {
    t := reflect.TypeOf(param)
    switch t.Kind() {
    case reflect.String:
        return param.(string)
    }
    body, _ := json.Marshal(param)
    return string(body)
}
