package Service

import (
	"URLShortener/Models"
	"errors"

	"gorm.io/gorm"
)

func GetAndIncreasesVisitCount(db *gorm.DB, shortCode string) (*Models.URL, error) {
	var url Models.URL

	// 查找对应的短链接
	if err := db.Where("short_code=?", shortCode).First(&url).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Short URL not found")
		}
	}

	// 增加访问计数
	if err := db.Model(&url).
		Where("id=?", url.ID).
		UpdateColumn("visit_count", gorm.Expr("visit_count+1")).Error; err != nil {
		return nil, err
	}

	return &url, nil
}
