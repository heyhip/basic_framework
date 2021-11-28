package common

import "strconv"

/******************* 错误码信息 *******************/

/****** 中文简体 ******/
var zhMsgInt = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	PARAM_ERROR: "参数错误(" + strconv.Itoa(PARAM_ERROR) + ")",

	TOKEN_EXPIRED: "请登录(" + strconv.Itoa(TOKEN_EXPIRED) + ")",
}

// 根据下标code获取信息
func GetZhMsgByCode(code int) string {
	msg, ok := zhMsgInt[code]
	if ok {
		return msg
	}
	return zhMsgInt[ERROR]
}
