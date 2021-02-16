package e

import "errors"

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	NOT_LOGINED:    "please login first",
	INVALID_PARAMS: "please check your request",
	UNAUTHORIZED:   "please check your authorize",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

func GetError(code int) error {
	return errors.New(GetMsg(code))
}
