package main

const (
	EInvalidParams ErrorCode = iota
	EInvalidCredentials
	EInternalError
)

func (c *ErrorCode) toString() string {
	code := int(*c)
	errorCodes := []string{
		"EInvalidParams",
		"EInvalidCredentials",
		"EInternalError",
	}
	if code >= len(errorCodes) {
		return ""
	}

	return errorCodes[code]
}
