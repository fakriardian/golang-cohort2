package middlewares

import (
	"final-challenge/dtos"
	"final-challenge/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.Request.Header.Get("Authorization")
		bearer := strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.Response{
				Error:   "Unauthorized",
				Status:  http.StatusUnauthorized,
				Message: "error",
			})
			return
		}

		stringToken := strings.Split(headerToken, " ")[1]
		verifyToken, err := helpers.VerifyToken(stringToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.Response{
				Error:   cases.Title(language.AmericanEnglish).String(err.Error()),
				Status:  http.StatusUnauthorized,
				Message: "error",
			})
			return
		}

		ctx.Request.Header.Set("AdminID", verifyToken)
		ctx.Next()
	}
}
