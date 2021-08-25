package errors

var (
	SUCCESS               = ERROR(0, "success")
	ERR_SYSTEM_ERROR      = ERROR(-1, "系统繁忙")
	ERR_DECRYP_ERROR      = ERROR(1004, "解密错误")
	ERR_PARAM_ERROR       = ERROR(1005, "请求参数非法")
	ERR_PARAM_PARSE_ERROR = ERROR(1007, "json或xml解析错误")
	ERR_RATE_ERROR        = ERROR(1014, "请求过于频繁")
	ERR_EMPTY_LIST_ERROR  = ERROR(6003, "空白的列表")
)
