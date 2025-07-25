package response

import (
	"net/http"

	"github.com/XieChengG/bookapi/exception"
	"github.com/gin-gonic/gin"
)

// error handler
func Failed(ctx *gin.Context, err error) {
	if e, ok := err.(*exception.ApiException); ok {
		if e.HttpCode == 0 {
			e.HttpCode = 500
		}
		ctx.JSON(e.HttpCode, e)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error()})
	}
}

// success handler
func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}
