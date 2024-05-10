package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrNotFound = errors.New(http.StatusText(http.StatusNotFound))                       // 404
var ErrInternalServerError = errors.New(http.StatusText(http.StatusInternalServerError)) // 500

func ErrorHandler(ctx *gin.Context, err error) {
	if err == ErrNotFound {
		// 404
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if err == ErrInternalServerError {
		// 500
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
