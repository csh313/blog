package models

// 广告表
type AdvertModel struct {
	MODEL  `structs:"-"`
	Title  string `json:"title" binding:"required" msg:"请输入标题"  structs:"title" gorm:"size:32"` //广告标题
	Href   string `json:"href" binding:"required,url" msg:"请输入跳转链接"  structs:"href"`            //跳转链接
	Images string `json:"images" binding:"required,url" msg:"请输入图片地址"  structs:"images"`        //广告图片
	IsShow bool   `json:"is_show" structs:"is_show"`                                            //是否显示
}
