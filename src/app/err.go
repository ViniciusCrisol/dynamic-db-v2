package app

import "strconv"

const (
	INTERNAL_ERR_STATUS = 500
	INTERNAL_ERR_MSG    = "500 - Internal server error."
	//
	VALIDATION_ERR_STATUS = 400
	//
	ROUTE_NOT_FOUND_ERR_STATUS = 404
	ROUTE_NOT_FOUND_ERR_MSG    = "404 - Route not found. Change the URL and try it again."
)

// Errs provides all the formatted application error messages. It could handle errors in a
// friendly way for the final user. Every message has an HTTP status prefix, use the func
// GetHTTPErr to get this status.
var Errs = map[string]string{
	// Default messages
	"INTERNAL":        INTERNAL_ERR_MSG,
	"ROUTE_NOT_FOUND": ROUTE_NOT_FOUND_ERR_MSG,
	// Application messages
	"INVALID-SCHEMA":      "400 - Invalid schema ID. Change the ID and try it again.",
	"INVALID-CLUSTER":     "400 - Invalid cluster name. Change the name and try it again.",
	"CLUSTER-NAME-IN-USE": "400 - Cluster name already in use. Change the name and try it again.",
}

// GetHTTPErr returns the Errs message HTTP status code. If the received error doesn't
// exist in the error list, a default message will be returned.
func GetHTTPErr(err error) (string, int) {
	errMessage := err.Error()
	appErrMsg, ok := Errs[errMessage]
	if !ok {
		return INTERNAL_ERR_MSG, INTERNAL_ERR_STATUS
	}
	status := appErrMsg[:3]
	parsed, _ := strconv.Atoi(status)
	return appErrMsg, parsed
}
