package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"status": false, "message": err.Error()})
		return // signal that there was an error and the caller should return
	}
}
