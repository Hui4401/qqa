package model

import "gorm.io/gorm"

// Question 问题模型
type Question struct {
	gorm.Model
	UserId  uint    `gorm:"not null;"`           // 问题所属用户Id
	Title   string `gorm:"not null;"`           // 标题
	Content string `gorm:"type:text;not null;"` // 内容
}

// 用ID获取问题
func GetQuestion(id uint) (Question, error) {
	var question Question
	result := DB.First(&question, id)
	return question, result.Error
}