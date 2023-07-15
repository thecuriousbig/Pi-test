package errmsg

import "pi/pkg/meta"

var (

	// 1000 - 1999: system error
	InternalServer   = meta.MetaErrorInternalServer.AppendMessage(1000, "The server encountered an internal error or misconfiguration and was unable to complete your request.")
	Forbidden        = meta.MetaErrorForbidden.AppendMessage(1001, "You do not have permission to access this resource.")
	MetaDataNotFound = meta.Error.AppendMessage(1002, "Metadata not found.")

	// 2000 - 2999: user error
	UserNotFound     = meta.Error.AppendMessage(2000, "User not found.")
	UserExisted      = meta.Error.AppendMessage(2001, "User already existed.")
	UserUpdateFailed = meta.Error.AppendMessage(2002, "User update failed.")
	UserCreateFailed = meta.Error.AppendMessage(2003, "User create failed.")
	UserDeleteFailed = meta.Error.AppendMessage(2004, "User delete failed.")
	UserCacheUpdate  = meta.Error.AppendMessage(2005, "User cache update failed.")
	UserCacheDelete  = meta.Error.AppendMessage(2006, "User cache delete failed.")
)

func ErrorInvalidRequest(msg string) *meta.MetaError {
	return meta.MetaErrorBadRequest.AppendMessage(1002, msg)
}

func ParseError(c int, desc string) *meta.MetaError {
	return meta.Error.AppendMessage(c, desc)
}
