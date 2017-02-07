package feedback

import (
	"database/sql"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

type Feedback struct {
	Id        int64  `json:"id"`
	Uid       int64  `json:"uid"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Ip        string `json:"ip"`
	UserAgent string `json:"userAgent"`
	Remark    string `json:"remark"`
	Ctime     int64  `json:"ctime"`
}

type IFeedbackApp interface {
	app.IApp
	GetDB() (*sql.DB, error)
	GetPrefix() string
	GetFeedbackTable() *kk.DBTable
}

type FeedbackApp struct {
	app.App
	DB *app.DBConfig

	Remote *remote.Service

	Feedback      *FeedbackService
	FeedbackTable kk.DBTable
}

func (C *FeedbackApp) GetDB() (*sql.DB, error) {
	return C.DB.Get(C)
}

func (C *FeedbackApp) GetPrefix() string {
	return C.DB.Prefix
}

func (C *FeedbackApp) GetFeedbackTable() *kk.DBTable {
	return &C.FeedbackTable
}
