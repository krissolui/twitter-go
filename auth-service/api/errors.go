package main

const (
	EInvalidPath ErrorCode = iota
	EBadRequest  ErrorCode = iota
	EInvalidParams
	EInvalidCredentials
	EInternalError
)

func (c *ErrorCode) toString() string {
	code := int(*c)
	errorCodes := []string{
		"EInvalidPath",
		"EBadRequest",
		"EInvalidParams",
		"EInvalidCredentials",
		"EInternalError",
	}
	if code >= len(errorCodes) {
		return ""
	}

	return errorCodes[code]
}
