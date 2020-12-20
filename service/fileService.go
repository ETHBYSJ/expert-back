package service

import (
	"expert-back/pkg/conf"
	"expert-back/pkg/e"
	"expert-back/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)


// 文件相关
type FileService struct {

}

// 删除某id对应的文件
func removeFileListById(id primitive.ObjectID, path string) {
	dir, err := os.Open(path)
	if err != nil {
		return
	}
	defer dir.Close()
	list, err := dir.Readdir(-1)
	for _, f := range list {
		fileName := f.Name()
		if strings.HasPrefix(fileName, id.Hex()) {
			os.Remove(filepath.Join(path, fileName))
		}
	}
}

// 保存上传的文件
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}

// 下载文件
func (service *FileService) DownloadFile(c *gin.Context, path string, name string) response.Response {
	file, err := os.Open(path)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorDownload,
		})
	}
	defer file.Close()
	// 设置头
	c.Header("Content-Disposition", "attachment; filename=" + url.QueryEscape(name))
	c.Header("Content-Type", "application/octet-stream")
	http.ServeContent(c.Writer, c.Request, name, time.Now(), file)
	return response.BuildResponse(map[int]interface{}{})
}

// 上传文件，通用函数
func (service *FileService) uploadFile(c *gin.Context, userID primitive.ObjectID, path string) response.Response {
	file, err := c.FormFile("file")
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorUpload,
		})
	}
	fileName := userID.Hex() + "_" + file.Filename
	fullPath := filepath.Join(path, fileName)
	removeFileListById(userID, path)
	err = SaveUploadedFile(file, fullPath)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorUpload,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: fullPath,
	})
}

// 上传文件
func (service *FileService) UploadRecommendFile(c *gin.Context, userID primitive.ObjectID) response.Response {
	return service.uploadFile(c, userID, conf.SystemConfig.File.Upload.Recommend.Path)
}

// 上传照片
func (service *FileService) UploadPhoto(c *gin.Context, userID primitive.ObjectID) response.Response {
	file, err := c.FormFile("photo")
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorUpload,
		})
	}
	path := conf.SystemConfig.File.Upload.Picture.Path
	fileName := userID.Hex() + "_" + file.Filename
	fullPath := filepath.Join(path, fileName)
	removeFileListById(userID, path)
	err = SaveUploadedFile(file, fullPath)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorUpload,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: "./static/" + fileName,
	})
}


