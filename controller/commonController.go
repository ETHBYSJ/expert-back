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
	c.SetCookie("iadmin", "fc41979b-736a-4842-95d3-9b9fa61ede3a", 3 * 24 * 3600, "/", "", false, true)
	c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{}))
}
