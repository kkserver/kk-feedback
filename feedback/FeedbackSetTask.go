package feedback

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FeedbackSetTaskResult struct {
	app.Result
	Feedback *Feedback `json:"feedback,omitempty"`
}

type FeedbackSetTask struct {
	app.Task
	Id      int64       `json:"id"`
	Uid     interface{} `json:"uid"`
	Phone   interface{} `json:"phone"`
	Email   interface{} `json:"email"`
	Author  interface{} `json:"author"`
	Content interface{} `json:"content"`
	Remark  interface{} `json:"remark"`
	Result  FeedbackSetTaskResult
}

func (task *FeedbackSetTask) GetResult() interface{} {
	return &task.Result
}

func (task *FeedbackSetTask) GetInhertType() string {
	return "feedback"
}

func (task *FeedbackSetTask) GetClientName() string {
	return "Feedback.Set"
}
