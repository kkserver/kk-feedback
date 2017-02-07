package feedback

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type FeedbackRemoveTaskResult struct {
	app.Result
	Feedback *Feedback `json:"feedback,omitempty"`
}

type FeedbackRemoveTask struct {
	app.Task
	Id     int64 `json:"id"`
	Result FeedbackRemoveTaskResult
}

func (task *FeedbackRemoveTask) GetResult() interface{} {
	return &task.Result
}

func (task *FeedbackRemoveTask) GetInhertType() string {
	return "feedback"
}

func (task *FeedbackRemoveTask) GetClientName() string {
	return "Feedback.Remove"
}
