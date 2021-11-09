package todo

import "github.com/beego/beego/v2/client/orm"

type Repository interface {
	Create(todo Todo) (int64, error)
	Update(todo Todo)
	Delete(todo Todo)
	Get(todo Todo)
}

type RepositoryImpl struct {
	Orm orm.Ormer
}

func (r RepositoryImpl) Create(todo Todo) (int64,error) {
	id, err :=  r.Orm.Insert(&todo)
	return id,err
}

func (r RepositoryImpl) Update(todo Todo) {
	r.Orm.Update(&todo)
}

func (r RepositoryImpl) Delete(todo Todo) {
	r.Orm.Delete(&todo)
}

func (r RepositoryImpl) Get(todo Todo) {
	r.Orm.Read(&todo)
}