package e

var MsgFlags = map[int]string {
	SUCCESS: 				"ok",
	ERROR: 					"error",
	INVALID_PARAMS: 		"请求参数错误",
	ERROR_EXIST_TAG:		"已经存在该标签名称",
	ERROR_NOT_EXIST_TAG:	"该标签名称不存在",
	ERROR_NOT_EXIST_ARTICLE:"该文章不存在",
}

func GetMsg(code int) string{
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}