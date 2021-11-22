package task

type Task struct {
	Id         int    `orm:"auto,column(id)"`
	Name       string `orm:"column(name)"`
	CreatedAt  string `orm:"column(created_at)"`
	FinishedAt string `orm:"column(finished_at)"`
	Status     string `orm:"column(status)"`
}
