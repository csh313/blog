package user_api

import (
	"fmt"
	"server/global"
	"server/models"
	"server/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRemoveView 批量删除用户
// @Tags 用户管理
// @Summary 批量删除用户
// @Description 批量删除用户
// @Param token header string true "用户token"
// @Param data body models.RemoveRequest    true  "用户id列表"
// @Router /api/users [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (*UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 判断用户是否存在
	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 删除用户前先删除用户和UserCollectModel以及ArticleModel的关联
	// 删除用户的收藏文章和发布的文章
	// 然后才能删除用户
	// 使用事务
	// 确保整个操作（清除关联和删除菜单）都在事务内处理，
	// 并且在操作过程中出现任何问题都会导致事务回滚，从而保持数据的完整性。

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO: 删除用户，评论表，消息表，用户收藏的文章，用户发布的文章

		// 使用事务对象来删除 menuList
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}

		return nil
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)
}
