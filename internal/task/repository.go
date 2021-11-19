package task

import "github.com/beego/beego/v2/client/orm"

type Repository interface {
	Create(Task Task) (Task, error)
	Update(Task Task)
	Delete(Task Task)
	Find(id int) (Task, error)
}

type RepositoryImpl struct {
	Manager orm.Ormer
}

func (r RepositoryImpl) Create(task Task) (Task, error) {
	id, err := r.Manager.Insert(&task)
	if err != nil {
		return Task{}, err
	}
	return r.Find(int(id))
}

func (r RepositoryImpl) Update(task Task) {
	r.Manager.Update(&task)
}

func (r RepositoryImpl) Delete(task Task) {
	r.Manager.Delete(&task)
}

func (r RepositoryImpl) Find(id int) (Task, error) {
	var task Task
	err := r.Manager.Raw("SELECT id,name,created_at,finished_at,status FROM task WHERE id = ?", id).QueryRow(&task)
	return task, err
}
