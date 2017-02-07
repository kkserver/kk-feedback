package feedback

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FeedbackCreateTaskResult struct {
	app.Result
	Feedback *Feedback `json:"feedback,omitempty"`
}

type FeedbackCreateTask struct {
	app.Task
	Uid       int64  `json:"uid"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Ip        string `json:"ip"`
	UserAgent string `json:"userAgent"`
	Remark    string `json:"remark"`
	Result    FeedbackCreateTaskResult
}

func (task *FeedbackCreateTask) GetResult() interface{} {
	return &task.Result
}

func (task *FeedbackCreateTask) GetInhertType() string {
	return "feedback"
}

func (task *FeedbackCreateTask) GetClientName() string {
	return "Feedback.Create"
}
