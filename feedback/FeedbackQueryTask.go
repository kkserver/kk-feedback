package feedback

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FeedbackQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
	RowCount  int `json:"rowCount"`
}

type FeedbackQueryTaskResult struct {
	app.Result
	Counter   *FeedbackQueryCounter `json:"counter,omitempty"`
	Feedbacks []Feedback            `json:"feedbacks,omitempty"`
}

type FeedbackQueryTask struct {
	app.Task
	Id        int64  `json:"id"`
	Keyword   string `json:"q"`
	OrderBy   string `json:"orderBy"` // desc, asc
	PageIndex int    `json:"p"`
	PageSize  int    `json:"size"`
	Counter   bool   `json:"counter"`
	Result    FeedbackQueryTaskResult
}

func (task *FeedbackQueryTask) GetResult() interface{} {
	return &task.Result
}

func (task *FeedbackQueryTask) GetInhertType() string {
	return "feedback"
}

func (task *FeedbackQueryTask) GetClientName() string {
	return "Feedback.Query"
}
