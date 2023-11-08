package helpers

import (
	"encoding/json"
	"errors"
	"final-challenge/dtos"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func errTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "format is invalid"
	case "gte":
		return "minimum length is 6 characters"
	}
	return fe.Error()
}

func formatedErrBind(ctx *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errorMessage := ""
		// uncomment this if want to array object
		// errorMessage := make([]struct {
		// 	Param, Message string
		// }, len(ve))
		for _, fe := range ve {
			errorMessage += fmt.Sprintf("%s %s;", fe.Field(), errTag(fe))
			// uncomment this if want to array object
			// errorMessage[i] = struct {
			// 	Param   string
			// 	Message string
			// }{fe.Field(), errTag(fe)}
		}
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   errorMessage,
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	} else if syntaxErr, ok := err.(*json.SyntaxError); ok {
		ctx.JSON(http.StatusBadRequest, dtos.Response{
			Error:   syntaxErr.Error(),
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

}

func BindRequest(ctx *gin.Context, req interface{}) error {
	appJSON := "application/json"
	contentType := ctx.Request.Header.Get("Content-Type")

	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(req); err != nil {
			formatedErrBind(ctx, err)
			return err
		}
	} else {
		if err := ctx.ShouldBind(req); err != nil {
			formatedErrBind(ctx, err)
			return err
		}
	}

	return nil
}
