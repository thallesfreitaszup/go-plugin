package database

import (
	"github.com/beego/beego/v2/client/orm"
	"poc-plugin/internal/task"
	"time"
)

func CreateDBManager() orm.Ormer {
	orm.Debug = true
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// set default database
	orm.RegisterDataBase("default", "postgres", "postgres://teste:teste@localhost/teste?sslmode=disable")

	// register model
	orm.RegisterModel(new(task.Task))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(WebhookDB))
	orm.RegisterModel(new(TaskLog))
	orm.RegisterModel(new(UserLog))

	// create table
	orm.RunSyncdb("default", false, true)
	return orm.NewOrm()
}

type User struct {
	Id int		`orm:"auto,column(id)"`
	Name   string  `orm:"column(name)"`
	Email string	 `orm:"column(email)"`
	Password string `orm:"column(password)"`
}

type WebhookDB struct {
	Id   int `orm:"auto,column(id)"`
	URL string `orm:"column(url)"`
	Events string `orm:"type(jsonb);column(events)"`
}

type TaskLog struct {
	Id int `orm:"auto;column(id)"`
	Action string `orm:"column(action)"`
	RequestId string `orm:"column(request_id)"`
	User *User `orm:"null;rel(fk);on_delete(set_null)"`
	Task *task.Task `orm:"null;rel(fk);on_delete(set_null)"`
	Timestamp  time.Time `orm:"column(timestamp)"`
}

type UserLog struct {
	Id int `orm:"auto;column(id)"`
	Action string `orm:"column(action)"`
	RequestId string `orm:"column(request_id)"`
	User *User `orm:"null;rel(fk);on_delete(set_null)"`
	Timestamp  time.Time `orm:"column(timestamp)"`
}
