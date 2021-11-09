package log

import (
  "github.com/beego/beego/v2/client/orm"
)

type Repository struct {
  Orm orm.Ormer
}

func (r Repository) Create(todoLog TodoLog) (int64, error) {
  return r.Orm.Insert(&todoLog)
}

func (r Repository) FindByRequestId(id string) ([]TodoLog,error) {
  var todoLog []TodoLog
  _, err := r.Orm.QueryTable(TodoLog{}).Filter("request_id", id).All(&todoLog)
  if err != nil {
    return todoLog, err
  }
  return todoLog, nil
}

func (r Repository) CreateUserLog(log UserLog) (int64, error) {
  return r.Orm.Insert(&log)
}
