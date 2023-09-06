package main

const (
	EBadRequest = iota
	EInvalidParams
	EInvalidAction
	EInvalidCredentials
)

func errorCode(err int) ErrorCode {
	return []ErrorCode{"EBadRequest", "EInvalidParams", "EInvalidAction", "EInvalidCredentials"}[err]
}
