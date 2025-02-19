package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
	"server/global"
	"server/models/res"
	"server/plugins/AliOss"
)

// 上传文件，如果开启了oss则再上传到oss，返回上传oss的文件路径
func Upload(file *multipart.FileHeader, c *gin.Context) string {
	var err error
	dst := path.Join(global.Config.Upload.Path, file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		res.FailWithMessage("gin上传文件失败", c)
		return ""
	}
	if global.Config.AliyunOSS.Enable {
		fmt.Println("-----------", global.Config.AliyunOSS.BucketName)
		dst, _ = AliOss.UploadFile(global.Config.AliyunOSS.BucketName, file.Filename, dst)
	}
	return dst
}
