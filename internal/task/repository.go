package task

import "github.com/beego/beego/v2/client/orm"

type Repository interface {
	Create(Task Task) (Task, error)
	Update(Task Task) (Task, error)
	Delete(Task Task) error
	Find() []Task
}

type RepositoryImpl struct {
	Manager orm.Ormer
}

func (r RepositoryImpl) Create(task Task) (Task, error) {
	id, err := r.Manager.Insert(&task)
	if err != nil {
		return Task{}, err
	}
	return r.FindById(int(id))
}

func (r RepositoryImpl) Update(task Task) (Task, error) {
	_, err := r.Manager.Update(&task)
	if err != nil {
		return task, err
	}
	return r.FindById(task.Id)
}

func (r RepositoryImpl) Delete(task Task) error {
	_, err := r.Manager.Delete(&task)
	return err
}

func (r RepositoryImpl) FindById(id int) (Task, error) {
	var task Task
	err := r.Manager.Raw("SELECT id,name,created_at,finished_at,status FROM task WHERE id = ?", id).QueryRow(&task)
	return task, err
}

func (r RepositoryImpl) Find() []Task {
	var tasks []Task
	r.Manager.QueryTable(&Task{}).All(&tasks)
	return tasks
}
