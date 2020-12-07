package service

import (
	"expert-back/pkg/conf"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"expert-back/vo"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"time"
)

type DownloadType int

const (
	Recommend	DownloadType = 0
)


// 存放映射关系
var downloadMap map[DownloadType]conf.DownloadFileInfo

// 文件相关
type FileService struct {

}

// 注意初始化顺序
func InitDownloadMap() {
	downloadMap = map[DownloadType]conf.DownloadFileInfo{
		0: conf.SystemConfig.File.Download.Recommend.DownloadFileInfo,
	}
}

// 下载文件
func (service *FileService) DownloadFile(c *gin.Context, downloadFileVO *vo.DownloadFileVO) response.Response {
	downloadType := DownloadType(downloadFileVO.Type)
	if _, ok := downloadMap[downloadType]; !ok {
		return response.BuildResponse(map[int]interface{}{
			response.CODE: e.ERROR_FILE_DOWNLOAD_INVALID_TYPE,
		})
	}
	file, err := os.Open(downloadMap[downloadType].Path)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.CODE: e.ERROR_DOWNLOAD,
		})
	}
	defer file.Close()
	// 设置头
	c.Header("Content-Disposition", "attachment; filename=" + url.QueryEscape(downloadMap[downloadType].Name))
	c.Header("Content-Type", "application/octet-stream")
	http.ServeContent(c.Writer, c.Request, downloadMap[downloadType].Name, time.Now(), file)
	return response.BuildResponse(map[int]interface{}{})
}

// 上传文件
/*
func (service *FileService) UploadFile(c *gin.Context) response.Response {

}
*/