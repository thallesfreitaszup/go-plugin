package todo

type Service struct {
	Repository Repository
}

func (s Service) Create(todo Todo) (int64, error){
	return s.Repository.Create(todo)
}


func (s Service) Update(todo Todo) {
	s.Repository.Update(todo)
}

func (s Service) Delete(todo Todo) {
 s.Repository.Delete(todo)
}

func (s Service) Get(todo Todo) {
 s.Repository.Get(todo)
}