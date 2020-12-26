package serializer

import "qa_go/model"

// 回答列表每一项数据
type AnswersData struct {
	ID          uint   `json:"id"`
	QuestionID  uint   `json:"qid"`
	Content     string `json:"content"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

// 回答列表响应信息
type AnswersResponse struct {
	Count   int          `json:"count"`
	Answers []AnswerData `json:"answers"`
}

// 序列化回答列表响应
func BuildAnswersResponse(answers []model.Answer) *AnswersResponse {
	var answersResponse AnswersResponse
	answersResponse.Count = len(answers)
	for _, answer := range answers {
		userProfile, _ := model.GetUserProfile(answer.UserID)
		answersResponse.Answers = append(answersResponse.Answers, AnswerData{
			ID:          answer.ID,
			QuestionID:  answer.QuestionID,
			Content:     answer.Content,
			Avatar:      userProfile.Avatar,
			Nickname:    userProfile.Nickname,
			Description: userProfile.Description,
			CreatedAt:   answer.CreatedAt.Unix(),
		})
	}
	return &answersResponse
}
