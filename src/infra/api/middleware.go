package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciuscrisol/dynamic-db-v2/app"
)

// SendJSON returns to the HTTP client the data in a JSON.
func SendJSON(status int, data any, ctx *gin.Context) {
	resp := DefaultResponse{
		Data:   data,
		Status: status,
	}
	ctx.JSON(status, resp)
}

// HandleErr searches for error messages in the Errs list using the GetHTTPErr func. If an
// matches, its message and status will be returned to the HTTP client. Otherwise, a
// standard error will be returned.
func HandleErr(err error, context *gin.Context) {
	msg, status := app.GetHTTPErr(err)
	resp := DefaultResponse{
		Message: msg,
		Status:  status,
	}
	context.JSON(status, resp)
}
