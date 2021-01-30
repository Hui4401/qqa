package model

import "gorm.io/gorm"

// Answer 回答模型
type Answer struct {
	gorm.Model
	UserID     uint   `gorm:"not null;"`           // 回答所属用户Id
	QuestionID uint   `gorm:"not null;"`           // 回答所属问题Id
	Content    string `gorm:"type:text;not null;"` // 内容
	LikeCount  uint   `gorm:"not null;"`           // 点赞数
}

// GetAnswer 用ID获取回答
func GetAnswer(id uint) (*Answer, error) {
	var answer Answer
	result := DB.First(&answer, id)
	return &answer, result.Error
}
//GetAnswers 用ID获取回答列表
func GetAnswers(ids []uint)([]Answer,error){
	var ans []Answer
	for _,id:= range ids{
		a,_:=GetAnswer(id)
		ans=append(ans, *a)
	}
	return ans,nil
}
//根据questionID获取回答总数
func GetAnswerCount(id uint) int64 {
	var cnt int64
	DB.Model(&Answer{}).Where("question_id = ?", id).Count(&cnt)
	return cnt
}

// 删除回答
func DeleteAnswer(id uint) error {
	result := DB.Delete(&Answer{}, id).Error
	return result
}

// 获取回答列表，按创建时间降序排列
func GetAnswersByTime(questionID uint, limit int, offset int) ([]Answer, error) {
	var answers []Answer
	result := DB.Where("question_id = ?", questionID).Order("created_at desc").Limit(limit).Offset(offset).Find(&answers)
	return answers, result.Error
}

// 获取回答列表，按热度排序，暂按时间升序
func GetAnswersByScore(questionID uint, limit int, offset int) ([]Answer, error) {
	var answers []Answer
	result := DB.Where("question_id = ?", questionID).Order("created_at").Limit(limit).Offset(offset).Find(&answers)
	return answers, result.Error
}
//获取某回答的点赞总数
func  GetAnswerlikedCount(aid uint)(uint,error){
	cnt,err:=GetLikeCountInCache(aid)
	if err!=nil{
		return 0,err
	}
	ans,err:=GetAnswer(aid)
	cnt+=ans.LikeCount
	return cnt,err
}
//获取某用户对某问题的点赞状态
func GetUserLikeStatus(uid uint,aid uint)(uint,error){
	return GetUserLike(uid,aid)
}
