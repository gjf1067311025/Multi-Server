package models

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/form"
	"github.com/jinzhu/gorm"
	"time"
)

type StarlingBasicInfo struct {
	ID                 uint      `json:"id" gorm:"column:id;primary_key"`
	Key                string    `json:"key" gorm:"column:key"`
	ChineseText        string    `json:"chinese_text" gorm:"column:chinese_text"`
	EnglishText        string    `json:"english_text" gorm:"column:english_text"`
	IsDel              int8      `json:"is_del" gorm:"column:is_del"`
	CreateTime         time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime         time.Time `json:"update_time" gorm:"column:update_time"`
	Remark             string    `json:"remark" gorm:"column:remark"`
	IsMachineTranslate bool      `json:"is_machine_translate" gorm:"column:is_machine_translate"`
}

func (StarlingBasicInfo) TableName() string {
	return "starling_basic_info"
}

func GetStarlingBasicByCondition(condition form.GetBasicInfoListByCondition) ([]*StarlingBasicInfo, int, error) {
	var ret []*StarlingBasicInfo
	var total int
	if condition.Key != "" {
		querySql := db.Table("starling_basic_info").Where("starling_basic_info.is_del = 0").
			Where(fmt.Sprintf(" starling_basic_info.key like '%%%s%%'", condition.Key))
		querySql.Count(&total)
		err := querySql.Offset((condition.PageNum - 1) * condition.PageSize).Limit(condition.PageSize).Order("create_time desc").Find(&ret).Error
		if err != nil {
			return ret, 0, err
		}
		return ret, total, nil
	}
	if condition.ChineseText != "" {
		querySql := db.Table("starling_basic_info").Where("starling_basic_info.is_del = 0 ").
			Where(fmt.Sprintf(" starling_basic_info.chinese_text like '%%%s%%'", condition.ChineseText))
		err := querySql.Count(&total).Offset((condition.PageNum - 1) * condition.PageSize).Limit(condition.PageSize).Order("create_time desc").Find(&ret).Error
		if err != nil {
			return ret, 0, err
		}
		return ret, total, nil
	}
	if condition.EnglishText != "" {
		querySql := db.Table("starling_basic_info").Where("starling_basic_info.is_del = 0").Where(fmt.Sprintf(" starling_basic_info.english_text like '%%%s%%'", condition.EnglishText))

		err := querySql.Count(&total).Offset((condition.PageNum - 1) * condition.PageSize).Limit(condition.PageSize).Order("create_time desc").Find(&ret).Error
		if err != nil {
			return ret, 0, err
		}
		return ret, total, nil
	}
	querySql := db.Table("starling_basic_info").Where("starling_basic_info.is_del = 0")
	err := querySql.Count(&total).Offset((condition.PageNum - 1) * condition.PageSize).Limit(condition.PageSize).Order("create_time desc").Find(&ret).Error
	if err != nil {
		return ret, 0, err
	}
	return ret, total, nil
}

func GetStarlingBasicBySearchKey(condition form.GetBasicInfoListBySearchKey) ([]*StarlingBasicInfo, int, error) {
	var ret []*StarlingBasicInfo
	var total int
	if condition.SearchKey != "" {
		querySql := db.Table("starling_basic_info").Where("starling_basic_info.is_del = 0").
			Where(fmt.Sprintf(" starling_basic_info.key like '%%%s%%' or starling_basic_info.chinese_text like '%%%s%%' or starling_basic_info.english_text like '%%%s%%'", condition.SearchKey, condition.SearchKey, condition.SearchKey)).Group("id")
		querySql.Count(&total)
		err := querySql.Offset((condition.PageNum - 1) * condition.PageSize).Limit(condition.PageSize).Order("create_time desc").Find(&ret).Error
		if err != nil {
			return ret, 0, err
		}
		return ret, total, nil
	}
	querySql := db.Table("starling_basic_info").Where("starling_basic_info.is_del = 0")
	err := querySql.Count(&total).Offset((condition.PageNum - 1) * condition.PageSize).Limit(condition.PageSize).Order("create_time desc").Find(&ret).Error
	if err != nil {
		return ret, 0, err
	}
	return ret, total, nil
}

func GetStarBasicByKeyAndLang(key string, lang string) (error, string) {
	var ret StarlingBasicInfo
	err := db.Table("starling_basic_info").Where("starling_basic_info.is_del = 0 and starling_basic_info.key = ?", key).First(&ret).Error
	if err != nil {
		return err, ""
	}
	if lang == "en" {
		return nil, ret.EnglishText
	}
	return nil, ret.ChineseText
}

func GetStarBasicInfoByID(id int) (error, *StarlingBasicInfo) {
	var ret StarlingBasicInfo
	err := db.Table("starling_basic_info").Where("starling_basic_info.id = ? and starling_basic_info.is_del = 0", id).First(&ret).Error
	return err, &ret
}

func CreateStarling(StarlingBasicInfo *StarlingBasicInfo, tx *gorm.DB) error {
	StarlingBasicInfo.CreateTime = time.Now()
	StarlingBasicInfo.UpdateTime = time.Now()
	if tx != nil {
		err := tx.Table("starling_basic_info").Create(StarlingBasicInfo).Error
		return err
	}
	err := db.Table("starling_basic_info").Create(StarlingBasicInfo).Error
	return err
}

func CreateBatch(StarlingBasicInfo []StarlingBasicInfo) error {
	// 事务一旦开始，你就应该使用 tx 处理数据
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if len(StarlingBasicInfo) > 0 {
		for _, basicInfo := range StarlingBasicInfo {
			if err := CreateStarling(&basicInfo, tx); err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

func UpdateStarling(StarlingBasicInfo *form.NewBasicInfo) error {
	update := map[string]interface{}{
		"update_time":          time.Now(),
		"key":                  StarlingBasicInfo.Key,
		"chinese_text":         StarlingBasicInfo.ChineseText,
		"english_text":         StarlingBasicInfo.EnglishText,
		"remark":               StarlingBasicInfo.Remark,
		"is_machine_translate": StarlingBasicInfo.IsMachineTranslate,
	}
	err := db.Table("starling_basic_info").Where("starling_basic_info.is_del = 0 and starling_basic_info.id = ?", StarlingBasicInfo.ID).Update(update).Error
	return err
}

func DeleteStarling(id int, tx *gorm.DB) error {
	update := map[string]interface{}{
		"is_del": 1,
	}
	if tx != nil {
		err := tx.Table("starling_basic_info").Where("starling_basic_info.id = ?", id).Update(update).Error
		return err
	}
	err := db.Table("starling_basic_info").Where("starling_basic_info.id = ?", id).Update(update).Error
	return err
}

func DeleteBatch(IDs []int) error {
	// 事务一旦开始，你就应该使用 tx 处理数据
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if len(IDs) > 0 {
		for _, id := range IDs {
			if err := DeleteStarling(id, tx); err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

func ExistStarByKey(key string) (bool, error) {
	count := 0
	err := db.Table("starling_basic_info").Where("starling_basic_info.key = ? and starling_basic_info.is_del = 0", key).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func GetStarByKey(key string) ([]*StarlingBasicInfo, error) {
	var ret []*StarlingBasicInfo
	err := db.Table("starling_basic_info").Where("starling_basic_info.key = ? and starling_basic_info.is_del = 0", key).Find(&ret).Error
	return ret, err
}
