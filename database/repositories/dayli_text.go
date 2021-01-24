package repositories

import (
	"github.com/StanislavDimitrenco/telegram-bot-reminder/database/models"
	"gorm.io/gorm"
)

type DailyText struct {
	db *gorm.DB
}

func NewDailyText(db *gorm.DB) *DailyText {
	return &DailyText{db: db}
}

func (repo DailyText) Create(dt *models.DailyText) *models.DailyText {
	repo.db.Create(&dt)
	return dt
}

func (repo DailyText) FindByDate(date string) *models.DailyText {
	var u models.DailyText
	repo.db.First(&u, "date = ?", date)
	return &u
}
