package controller

import (
	"expert-back/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonController struct {

}

// 设置cookie
func (controller *CommonController) SetCookie(c *gin.Context) {
	c.SetCookie("iadmin", "1b12ec79-7b1b-400d-bc81-78871dae7fea", 3 * 24 * 3600, "/", "", false, true)
	c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{}))
}
