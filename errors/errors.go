package errors

import (
	. "github.com/andrezhz/go-tools/context"
)

var (
	errorMsg         = make(map[int]string, 0)
	errcode2standMap = make(map[int]Error, 0)
)

type Error int

func (e Error) Error() string {
	return errorMsg[int(e)]
}

// ERROR
func ERROR(code int, msg string) Error {
	errorMsg[code] = msg
	return Error(code)
}

// SetErrCode2Stand
func SetErrCode2Stand(code int, e Error) {
	errcode2standMap[code] = e
}

// Result
func (e Error) Result() *OutPackage {
	return &OutPackage{
		Status: int32(e),
		Msg:    e.Error(),
	}
}

// ResultWithMsg
func (e Error) ResultWithMsg(msg string) *OutPackage {
	return &OutPackage{
		Status: int32(e),
		Msg:    msg,
	}
}

// ToErrMsg
func (e Error) ToErrMsg(errMsg *OutPackage) {
	errMsg.Status = int32(e)
	errMsg.Msg = e.Error()
}

// ErrCodeToStand
// conver err code to stand business code
func ErrCodeToStand(errMsg *OutPackage) error {
	err, ok := errcode2standMap[int(errMsg.Status)]
	if !ok {
		return ERR_SYSTEM_ERROR
	}
	err.ToErrMsg(errMsg)
	return err
}
