package task

type Service struct {
	Repository Repository
}

func (s Service) Create(Task Task) (Task, error) {
	return s.Repository.Create(Task)
}

func (s Service) Update(Task Task) {
	s.Repository.Update(Task)
}

func (s Service) Delete(Task Task) {
	s.Repository.Delete(Task)
}

func (s Service) Get(Task Task) {

}
