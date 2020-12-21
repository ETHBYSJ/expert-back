package service

import (
	"expert-back/model"
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
func removeFileListById(prefix string, path string) {
	dir, err := os.Open(path)
	if err != nil {
		return
	}
	defer dir.Close()
	list, err := dir.Readdir(-1)
	for _, f := range list {
		fileName := f.Name()
		if strings.HasPrefix(fileName, prefix) {
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
	c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(name))
	c.Header("Content-Type", "application/octet-stream")
	http.ServeContent(c.Writer, c.Request, name, time.Now(), file)
	return response.BuildResponse(map[int]interface{}{})
}

// 上传申请表
func (service *FileService) UploadApplyFile(c *gin.Context, userID primitive.ObjectID) response.Response {
	file, err := c.FormFile("file")
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorUpload,
		})
	}
	fileName := userID.Hex() + "_" + file.Filename
	fullPath := filepath.Join(conf.SystemConfig.File.Upload.Apply.Path, fileName)
	// 删除旧文件
	removeFileListById(userID.Hex(), conf.SystemConfig.File.Upload.Apply.Path)
	err = SaveUploadedFile(file, fullPath)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorUpload,
		})
	}
	// 保存记录
	record := &model.Record{
		Type: model.Apply,
		UserID: userID,
		SubmitID: "",
		File: file.Filename,
	}
	err = model.SaveOrUpdateRecordBaseInfo(record)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorApplyRecordSet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: fullPath,
	})
}

// 上传推荐表
func (service *FileService) UploadRecommendFile(c *gin.Context, submitID string, userID primitive.ObjectID) response.Response {
	file, err := c.FormFile("file")
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorUpload,
		})
	}
	fileName := submitID + "_" + file.Filename
	fullPath := filepath.Join(conf.SystemConfig.File.Upload.Recommend.Path, fileName)
	// 删除旧文件
	removeFileListById(submitID, conf.SystemConfig.File.Upload.Recommend.Path)
	err = SaveUploadedFile(file, fullPath)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorUpload,
		})
	}
	// 保存记录
	record := &model.Record{
		Type: model.Recommend,
		UserID: userID,
		SubmitID: submitID,
		File: file.Filename,
	}
	err = model.SaveOrUpdateRecordBaseInfo(record)
	if err != nil {
		return response.BuildResponse(map[int]interface{}{
			response.Code: e.ErrorRecommendRecordSet,
		})
	}
	return response.BuildResponse(map[int]interface{}{
		response.Data: fullPath,
	})
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
	removeFileListById(userID.Hex(), path)
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
