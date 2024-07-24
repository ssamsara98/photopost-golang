package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssamsara98/photopost-golang/src/constants"
)

func BindBody[T any](c *gin.Context) *T {
	var body T
	if err := c.ShouldBind(&body); err != nil {
		ErrorJSON(c, http.StatusBadRequest, err)
		return nil
	}
	return &body
}

func BindUri[T any](c *gin.Context) *T {
	var uri T
	if err := c.ShouldBindUri(&uri); err != nil {
		ErrorJSON(c, http.StatusBadRequest, err)
		return nil
	}
	return &uri
}

func GetUser[T any](c *gin.Context) (*T, bool) {
	user, boolean := c.MustGet(constants.User).(*T)
	return user, boolean
}
