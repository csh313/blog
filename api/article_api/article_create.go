package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/models/res"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

// 文章上传请求参数
type ArticleRequest struct {
	Title    string      `json:"title" binding:"required" msg:"请输入文章标题"`   // 文章标题
	Abstract string      `json:"abstract"`                                 // 文章简介
	Content  string      `json:"content" binding:"required" msg:"请输入文章内容"` // 文章内容
	Category string      `json:"category"`                                 // 文章分类
	Source   string      `json:"source"`                                   // 文章来源
	Link     string      `json:"link"`                                     // 原文链接
	BannerID uint        `json:"banner_id"`                                // 文章封面id
	Tags     ctype.Array `json:"tags"`                                     // 文章标签

}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//获取当前用户信息
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	userID := claims["user_id"].(uint)
	userNickName := claims["nick_name"].(string)

	//处理content
	// 将markdown转为html
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// 是不是有script标签
	// 从html中获取文本内容
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//fmt.Println(doc.Text())
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		// 有script标签,移除
		doc.Find("script").Remove()
		// 将html转为md
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}
	//文章简介不存在，截取文章内容作为简介
	if cr.Abstract == "" {
		//汉字的截取不一样
		abs := []rune(doc.Text())
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs)
		}
	}

	//不传banner_id，后台就随机选择一张
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			res.FailWithMessage("没有banner数据", c)
			return
		}
		rand.Seed(int64(uint64(time.Now().UnixNano())))
		cr.BannerID = bannerIDList[rand.Intn(len(bannerIDList))]
	}

	// 查banner_id下的banner_url
	var bannerUrl string
	err := global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		res.FailWithMessage("banner不存在", c)
		return
	}

	// 查用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")

	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	//数据库中是否存在文章
	if affected := global.DB.Where("Title is ? and UserID is ?", cr.Title, userID).Find(&article).RowsAffected; affected > 0 {
		res.FailWithMessage("文章已存在，请修改文章或编辑标题", c)
		return
	}

	if err = global.DB.Create(&article).Error; err != nil {
		res.FailWithError(err, article, c)
		return
	}
	res.OkWithMessage("发布文章成功", c)

	// 查询文章是否已经存在
	// 将title多存储一份，改名为keyword
	// "keyword": {
	// "type": "keyword"
	// }
	// 便于搜索
	//if article.ISExistData() {
	//	res.FailWithMessage("文章已经存在", c)
	//	return
	//}

	// 创建一条文章数据
	//err = article.Create()
	//if err != nil {
	//	global.Log.Error(err)
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
	res.OkWithMessage("文章发布成功", c)
}
