package feedback

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FeedbackTaskResult struct {
	app.Result
	Feedback *Feedback `json:"feedback,omitempty"`
}

type FeedbackTask struct {
	app.Task
	Id     int64 `json:"id"`
	Result FeedbackTaskResult
}

func (task *FeedbackTask) GetResult() interface{} {
	return &task.Result
}

func (task *FeedbackTask) GetInhertType() string {
	return "feedback"
}

func (task *FeedbackTask) GetClientName() string {
	return "Feedback.Get"
}
