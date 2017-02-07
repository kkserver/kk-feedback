package feedback

import (
	"bytes"
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/dynamic"
	"time"
)

type FeedbackService struct {
	app.Service

	Get    *FeedbackTask
	Set    *FeedbackSetTask
	Remove *FeedbackRemoveTask
	Create *FeedbackCreateTask
	Query  *FeedbackQueryTask
}

func (S *FeedbackService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *FeedbackService) HandleFeedbackCreateTask(a IFeedbackApp, task *FeedbackCreateTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Feedback{}

	v.Uid = task.Uid
	v.Phone = task.Phone
	v.Email = task.Email
	v.Author = task.Author
	v.Content = task.Content
	v.Ip = task.Ip
	v.UserAgent = task.UserAgent
	v.Remark = task.Remark
	v.Ctime = time.Now().Unix()

	_, err = kk.DBInsert(db, a.GetFeedbackTable(), a.GetPrefix(), &v)

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Feedback = &v

	return nil
}

func (S *FeedbackService) HandleFeedbackSetTask(a IFeedbackApp, task *FeedbackSetTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Feedback{}

	rows, err := kk.DBQuery(db, a.GetFeedbackTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)
		if err != nil {
			task.Result.Errno = ERROR_FEEDBACK
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {
		task.Result.Errno = ERROR_FEEDBACK_NOT_FOUND
		task.Result.Errmsg = "Not Found feedback"
		return nil
	}

	keys := map[string]bool{}

	if task.Uid != nil {
		v.Uid = dynamic.IntValue(task.Uid, v.Uid)
		keys["uid"] = true
	}

	if task.Phone != nil {
		v.Phone = dynamic.StringValue(task.Phone, v.Phone)
		keys["phone"] = true
	}

	if task.Email != nil {
		v.Email = dynamic.StringValue(task.Email, v.Email)
		keys["email"] = true
	}

	if task.Author != nil {
		v.Author = dynamic.StringValue(task.Author, v.Author)
		keys["author"] = true
	}

	if task.Content != nil {
		v.Content = dynamic.StringValue(task.Content, v.Content)
		keys["content"] = true
	}

	if task.Remark != nil {
		v.Remark = dynamic.StringValue(task.Remark, v.Remark)
		keys["remark"] = true
	}

	_, err = kk.DBUpdateWithKeys(db, a.GetFeedbackTable(), a.GetPrefix(), &v, keys)

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Feedback = &v

	return nil
}

func (S *FeedbackService) HandleFeedbackTask(a IFeedbackApp, task *FeedbackTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Feedback{}

	rows, err := kk.DBQuery(db, a.GetFeedbackTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)
		if err != nil {
			task.Result.Errno = ERROR_FEEDBACK
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {
		task.Result.Errno = ERROR_FEEDBACK_NOT_FOUND
		task.Result.Errmsg = "Not Found feedback"
		return nil
	}

	task.Result.Feedback = &v

	return nil
}

func (S *FeedbackService) HandleFeedbackRemoveTask(a IFeedbackApp, task *FeedbackRemoveTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Feedback{}

	rows, err := kk.DBQuery(db, a.GetFeedbackTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)
		if err != nil {
			task.Result.Errno = ERROR_FEEDBACK
			task.Result.Errmsg = err.Error()
			return nil
		}

		_, err = kk.DBDelete(db, a.GetFeedbackTable(), a.GetPrefix(), " WHERE id=?", task.Id)

		if err != nil {
			task.Result.Errno = ERROR_FEEDBACK
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {
		task.Result.Errno = ERROR_FEEDBACK_NOT_FOUND
		task.Result.Errmsg = "Not Found feedback"
		return nil
	}

	task.Result.Feedback = &v

	return nil
}

func (S *FeedbackService) HandleFeedbackQueryTask(a IFeedbackApp, task *FeedbackQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	var feedbacks = []Feedback{}

	var args = []interface{}{}

	var sql = bytes.NewBuffer(nil)

	sql.WriteString(" WHERE 1")

	if task.Id != 0 {
		sql.WriteString(" AND id=?")
		args = append(args, task.Id)
	}

	if task.Keyword != "" {
		q := "%" + task.Keyword + "%"
		sql.WriteString(" AND (content LIKE ? OR author LIKE ? OR phone LIKE ? OR email LIKE ? OR remark LIKE ?)")
		args = append(args, q, q, q, q, q)
	}

	if task.OrderBy == "asc" {
		sql.WriteString(" ORDER BY id ASC")
	} else {
		sql.WriteString(" ORDER BY id DESC")
	}

	var pageIndex = task.PageIndex
	var pageSize = task.PageSize

	if pageIndex < 1 {
		pageIndex = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	if task.Counter {
		var counter = FeedbackQueryCounter{}
		counter.PageIndex = pageIndex
		counter.PageSize = pageSize
		counter.RowCount, err = kk.DBQueryCount(db, a.GetFeedbackTable(), a.GetPrefix(), sql.String(), args...)
		if err != nil {
			task.Result.Errno = ERROR_FEEDBACK
			task.Result.Errmsg = err.Error()
			return nil
		}
		if counter.RowCount%pageSize == 0 {
			counter.PageCount = counter.RowCount / pageSize
		} else {
			counter.PageCount = counter.RowCount/pageSize + 1
		}
		task.Result.Counter = &counter
	}

	sql.WriteString(fmt.Sprintf(" LIMIT %d,%d", (pageIndex-1)*pageSize, pageSize))

	var v = Feedback{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, a.GetFeedbackTable(), a.GetPrefix(), sql.String(), args...)

	if err != nil {
		task.Result.Errno = ERROR_FEEDBACK
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	for rows.Next() {

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_FEEDBACK
			task.Result.Errmsg = err.Error()
			return nil
		}

		feedbacks = append(feedbacks, v)
	}

	task.Result.Feedbacks = feedbacks

	return nil
}
