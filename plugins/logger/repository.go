package log

import (
  "fmt"
  "github.com/beego/beego/v2/client/orm"
  "poc-plugin/internal/configuration/database"
)

type Repository struct {
  Orm orm.Ormer
}

func (r Repository) Create(taskLog database.TaskLog) (int64, error) {
  return r.Orm.Insert(&taskLog)
}

func (r Repository) FindByRequestId(id string) (database.UserLog,error) {
  var userLogs database.UserLog
  query := fmt.Sprintf("SELECT id, action, request_id, user_id, timestamp FROM user_log where request_id = '%s' ", id)
  err := r.Orm.Raw(query).QueryRow(&userLogs)
  if err != nil {
    return userLogs, err
  }
  return userLogs, nil
}

func (r Repository) CreateUserLog(log database.UserLog) (int64, error) {
  return r.Orm.Insert(&log)
}

func (r Repository) getTasks() []database.TaskLog {
  var taskLogs []database.TaskLog
  r.Orm.QueryTable(database.TaskLog{}).RelatedSel().All(&taskLogs)
  return taskLogs
}
