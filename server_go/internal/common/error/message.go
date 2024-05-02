package error

var MsgFlag = map[int]string{
	SUCCESS:          "OK",
	INTERNAL_SERVER:  "Failed",
	UNAUTHORIZED:     "Authorization is failed. Please check the secure token",
	BAD_REQUEST:      "Request is invalid",
	VALIDATION_ERROR: "Validation: Invalid params",
	NOT_FOUND:        "Data not found",
}

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[INTERNAL_SERVER]
}
