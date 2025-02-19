package message_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/global"
	"server/models"
	"server/models/res"
	"server/service/pageService/common"
	"time"
)

type MessageCreateRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"`
	RevUserID  uint   `json:"rev_user_id" binding:"required"`
	Content    string `json:"content"      binding:"required"`
}

type Message struct {
	SendUserID       uint      `json:"send_user_id"` // 发送人id
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"` // 接收人id
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`       // 消息内容
	CreatedAt        time.Time `json:"created_at"`    // 最新的消息时间
	MessageCount     int       `json:"message_count"` // 消息条数
}

type MessageGroup map[uint]*Message

// MessageCreate发布消息
// @Tages消息管理
// @Summary发布消息
// @Description发布消息
// @Param data body MessageCreateRequest true "表示多个参数"
// @Param token header string true "token"
// @Router /api/message [post]
// Produce json
// Success 200 {object} res.Response{}
func (MessageApi) MessageCreate(c *gin.Context) {
	var req MessageCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	//查看数据库中发送者和接收者是否都存在
	var sendUser, revUser models.UserModel
	err := global.DB.Take(&sendUser, req.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Take(&revUser, req.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接收人不存在", c)
		return
	}
	fmt.Println("------")

	fmt.Println(req)
	fmt.Println(sendUser)
	fmt.Println(revUser)
	fmt.Println("-------")

	//将数据存入数据库
	if err = global.DB.Create(&models.MessageModel{
		SendUserID:        req.SendUserID,
		SendUserNicekName: sendUser.NickName,
		SendUserAvatar:    sendUser.Avatar,
		RevUserID:         req.RevUserID,
		RevUserNicekName:  revUser.NickName,
		RevUserAvatar:     revUser.Avatar,
		Content:           req.Content,
		IsRead:            false}).Error; err != nil {
		res.FailWithMessage("消息发送失败", c)
		return
	}
	res.OkWithMessage("消息发布成功", c)
}

func (MessageApi) MessageList(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)

	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages []Message

	global.DB.Order("created_at asc").Where("send_user_id=? or rev_user_id=?", claims["send_user_id"], claims["rev_user_id"]).Find(&messageList)
	for _, model := range messageList {
		message := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNicekName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserNicekName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	res.OkWithData(messages, c)
}

func (MessageApi) MessageListAll(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithError(err, &page, c)
		return
	}
	var messageList []models.MessageModel
	list, count, err := common.PageList(page, messageList)
	if err != nil {
		res.FailWithMessage("查询失败", c)
		return
	}
	res.OkWithList(list, count, c)

}
