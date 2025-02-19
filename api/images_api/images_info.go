package images_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"path"
	"server/global"
	"server/models"
	"server/models/res"
	"server/plugins/AliOss"
	"server/utils"
	"strings"
)

var (
	// 规定图片上传白名单
	WhiteImageList = []string{
		"pjp",
		"svgz",
		"jpg",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"jfif",
		"webp",
		"png",
		"bmp",
		"pjpeg",
		"avif",
	}
)

// 图片更新请求
type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择有效文件id"`
	Name string `json:"name" binding:"required" msg:"请输入修改后的文件名称"`
}

// 规定文件上传的响应格式，便于客户端解析上传结果信息
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 提示信息
}

func (ImagesApi) ImageUploadView(c *gin.Context) {
	//单文件上传
	//file, err := c.FormFile("image")
	//if err != nil {
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
	//log.Println(file.Filename)
	//多文件上传
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	files, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在文件", c)
		return
	}

	var resList []FileUploadResponse

	for _, file := range files {
		size := float64(file.Size) / float64(1024*1024)

		//判断文件是否为图片类型
		fileName := file.Filename
		suffix := path.Ext(fileName)
		suffix = strings.ToLower(suffix[1:])
		fmt.Println(fileName, suffix)
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}

		if size > global.Config.Upload.Size {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定大小(%vMB)", global.Config.Upload.Size),
			})
			continue
		}

		//将图片变为MD5
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		Data, err := io.ReadAll(fileObj)
		imageHash := utils.MD5(Data)
		var bannerModel models.BannerModel
		if err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error; err == nil {
			//数据库中已存在
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: true,
				Msg:       "数据库中已存在，图片上传成功",
			})
			continue
		}

		dst := path.Join(global.Config.Upload.Path, file.Filename)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		if global.Config.AliyunOSS.Enable {
			fmt.Println("-----------", global.Config.AliyunOSS.BucketName)
			dst, _ = AliOss.UploadFile(global.Config.AliyunOSS.BucketName, file.Filename, dst)
		}
		bannerModel = models.BannerModel{
			Path: dst,
			Hash: imageHash,
			Name: file.Filename,
		}
		global.DB.Create(&bannerModel)
		resList = append(resList, FileUploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "图片上传成功",
		})

	}

	res.OkWithData(resList, c)
	//context.JSON(200, gin.H{
	//	"code": 200,
	//})
}

func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var image models.BannerModel
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	if err = global.DB.Find(&image, cr.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.FailWithMessage("文件未找到", c)
			return
		} else {
			res.FailWithMessage(err.Error(), c)
			return
		}
	}

	if err = global.DB.Model(&image).Update("name", cr.Name).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
}
