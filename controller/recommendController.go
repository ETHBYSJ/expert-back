// 专家推荐
package controller

import (
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/service"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecommendController struct {
	recommendService service.RecommendService
}

// 专家推荐
func (controller *RecommendController) RecommendCommit(c *gin.Context) {
	var recommendVO vo.RecommendVO
	if err := c.ShouldBindJSON(&recommendVO); err == nil {
		res := controller.recommendService.RecommendCommit(c, &recommendVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.CODE: e.HTTP_BADREQUEST,
		}))
	}
}

// 专家推荐文件下载
func (controller *RecommendController) RecommendDownload(c *gin.Context) {
	res := controller.recommendService.RecommendDownload(c)
	if res.Code != e.SUCCESS {
		c.JSON(http.StatusOK, res)
	}
}

// 专家推荐文件上传
func (controller *RecommendController) RecommendUpload(c *gin.Context) {
	var recommendUploadVO vo.RecommendUploadVO
	if err := c.ShouldBindQuery(&recommendUploadVO); err == nil {
		res := controller.recommendService.RecommendUpload(c, &recommendUploadVO)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BuildResponse(map[int]interface{}{
			response.CODE: e.HTTP_BADREQUEST,
		}))
	}
}

/*
func (controller *RecommendController) TestParse(c *gin.Context) {
	doc, err := document.Open("./static/download/recommend.docx")
	if err != nil {
		util.Log().Panic("read doc error %v", err)
	}
	tables := doc.Tables()
	util.Log().Info("table num %v", len(tables))
	rows := tables[1].Rows()
	for i, row := range rows {
		cells := row.Cells()
		for _, cell := range cells {
			util.Log().Info("row %v", i)
			paragraphs := cell.Paragraphs()
			for j, para := range paragraphs {
				// util.Log().Info("%v", para.Runs())
				for _, run := range para.Runs() {
					util.Log().Info("    para %v %v", j, run.Text())
				}
			}
		}
	}
}
*/