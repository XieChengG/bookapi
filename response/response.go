package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// error handler
func Failed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err})
}

// success handler
func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}
