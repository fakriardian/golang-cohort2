package middlewares

import (
	"final-challenge/dtos"
	"final-challenge/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var uri dtos.ProductIDUri

		if err := ctx.ShouldBindUri(&uri); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
				Error:   err.Error(),
				Status:  http.StatusNotFound,
				Message: "error",
			})
			return
		}

		adminID := ctx.Request.Header.Get("AdminID")

		var product models.Product
		err := db.Select("user_id").Where("id = ?", uri.ID).First(&product).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
				Error:   err.Error(),
				Status:  http.StatusNotFound,
				Message: "Data Not Found",
			})
			return
		}

		if product.UserID.String() != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.Response{
				Error:   "Unauthorized",
				Status:  http.StatusUnauthorized,
				Message: "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
