package controller

import (
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/service"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileController struct {

}

// 文件下载
func (controller *FileController) DownloadFile(c *gin.Context) {
	var downloadFileVO vo.DownloadFileVO
	var fileService service.FileService
	if err := c.ShouldBindQuery(&downloadFileVO); err == nil {
		res := fileService.DownloadFile(c, &downloadFileVO)
		if res.Code != e.SUCCESS {
			c.JSON(http.StatusOK, res)
		}
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.CODE: e.HTTP_BADREQUEST,
		}))
	}
}