
package configuration

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	"poc-plugin/internal/todo"
	"poc-plugin/plugins/authorization"
	"poc-plugin/plugins/logger"
	"poc-plugin/plugins/webhook"
)


func init() {
	// set default database
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// set default database
	orm.RegisterDataBase("default", "postgres", "postgres://teste:teste@localhost/teste?sslmode=disable")

	// register model
	orm.RegisterModel(new(todo.Todo))
	orm.RegisterModel(new(webhook.WebhookDB))
	orm.RegisterModel(new(authorization.User))
	orm.RegisterModel(new(log.TodoLog))
	orm.RegisterModel(new(log.UserLog))
	// create table
	orm.RunSyncdb("default", false, true)
}

func GetDBConnection( ) orm.Ormer {
	return orm.NewOrm()
}



