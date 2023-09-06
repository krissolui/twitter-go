package main

const (
	EBadRequest = iota
	EInvalidParams
	EInvalidAction
)

func errorCode(err int) ErrorCode {
	return []ErrorCode{"EBadRequest", "EInvalidParams", "EInvalidAction"}[err]
}
