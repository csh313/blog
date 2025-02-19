package common

import (
	"log"
	"server/global"
	"server/models"
)

// 分页查询
func PageList[T any](page models.PageInfo, model T) (list []T, count int64, err error) {
	offset := (page.Page - 1) * page.Limit
	if offset < 0 {
		offset = 0
	}
	if page.Sort == "" {
		page.Sort = "created_at desc" //默认排序创建顺序desc从晚到早.asc从早到晚
	}
	//如果limit为0，则查询所有
	if page.Limit == 0 {
		page.Limit = -1
	}
	global.DB.Where(&model).Find(&list)
	//log.Println("-------", model)
	//log.Println("-------", list)
	//global.DB.Find(&model)
	//log.Println("===================", model)
	//count = global.DB.Select("id").Find(&model).RowsAffected
	global.DB.Model(&model).Count(&count)

	log.Println(count)
	err = global.DB.Debug().Model(model).Limit(page.Limit).Offset(offset).Order(page.Sort).Find(&list).Error
	//log.Println("===================", list)
	return list, count, err
}
