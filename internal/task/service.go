package task

type Service struct {
	Repository Repository
}

func (s Service) Create(Task Task) (Task, error) {
	return s.Repository.Create(Task)
}

func (s Service) Update(Task Task) (Task, error) {
	return s.Repository.Update(Task)
}

func (s Service) Delete(Task Task) error {
	return s.Repository.Delete(Task)
}

func (s Service) Get(Task Task) {

}

func (s Service) Find() []Task {
	return s.Repository.Find()
}
