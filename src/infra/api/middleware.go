package api

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	englishTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/viniciuscrisol/dynamic-db-v2/app"
)

const VALIDATION_ERR_PREFIX = "validation-error\n"

// SendJSON returns to the HTTP client the data in a JSON.
func SendJSON(status int, data any, ctx *gin.Context) {
	resp := DefaultResponse{
		Data:   data,
		Status: status,
	}
	ctx.JSON(status, resp)
}

// SendRouteNotFound returns to the HTTP client a route not found message in a JSON.
func SendRouteNotFound(ctx *gin.Context) {
	status := app.ROUTE_NOT_FOUND_ERR_STATUS
	resp := DefaultResponse{
		Status:  status,
		Message: app.ROUTE_NOT_FOUND_ERR_MSG,
	}
	ctx.JSON(status, resp)
}

// HandleErr returns to the HTTP client BindRequestBody errors. BindRequestBody could
// return a binding or a validation error. This function identifies the error type by the
// message prefix and deals with it.
func HandleErr(err error, ctx *gin.Context) {
	msg := err.Error()
	if strings.HasPrefix(msg, VALIDATION_ERR_PREFIX) {
		handleValidationErr(msg, ctx)
		return
	}
	handleAppErr(err, ctx)
}

// handleValidationErr  formats a validation error group and return it to the client.
func handleValidationErr(msg string, ctx *gin.Context) {
	formattedMsg := strings.TrimPrefix(msg, VALIDATION_ERR_PREFIX)
	resp := DefaultResponse{
		Message: formattedMsg,
		Status:  app.VALIDATION_ERR_STATUS,
	}
	ctx.JSON(app.VALIDATION_ERR_STATUS, resp)
}

// handleAppErr searches for error messages in the Errs list using the GetHTTPErr func. If
// an matches, its message and status will be returned to the HTTP client. Otherwise, a
// standard error will be returned.
func handleAppErr(err error, ctx *gin.Context) {
	msg, status := app.GetHTTPErr(err)
	resp := DefaultResponse{
		Message: msg,
		Status:  status,
	}
	ctx.JSON(status, resp)
}

// BindRequestBody binds the request body into the binder variable. It also validates the
// entries according to the "validate" tag of the struct fields. To change the validation
// message language, switch the translator language.
func BindRequestBody(binder any, ctx *gin.Context) error {
	err := ctx.ShouldBindJSON(binder)
	if err != nil {
		return err
	}
	validate := validator.New()
	err = validate.Struct(binder)
	if err == nil {
		return nil
	}
	english := en.New()
	translator := universalTranslator.New(english, english)
	englishTranslator, _ := translator.GetTranslator("en")
	englishTranslations.RegisterDefaultTranslations(validate, englishTranslator)

	return formatValidationMessages(err, englishTranslator)
}

// formatValidationMessages formats the validator output into a simple message. It is also
// used to translate messages using the BindRequestBody translator.
func formatValidationMessages(validationErrs error, englishTranslator universalTranslator.Translator) error {
	msg := VALIDATION_ERR_PREFIX
	errs := validationErrs.(validator.ValidationErrors)
	for _, validationErr := range errs {
		err := validationErr.Translate(englishTranslator)
		msg += err + "\n"
	}
	return errors.New(msg)
}
