// Code generated by gen_code. DO NOT EDIT.
package errmsg

const (
	TypeMsgSuccess = 10000
	TypeMsgFailed  = 10001
)

func GetSuccess() Resp {
	return Resp{Code: TypeCodeSuccess, Msg: TypeMsgSuccess}
}

func GetFailed() Resp {
	return Resp{Code: TypeCodeFailed, Msg: TypeMsgFailed}
}