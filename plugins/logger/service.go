package log

type Service struct {
	Repository Repository
}


func (s Service) CreateTodoLog(todoLog TodoLog) (int64, error) {
	return s.Repository.Create(todoLog)
}

func (s Service) CreateUserLog(userLog UserLog) (int64, error) {
	return s.Repository.CreateUserLog(userLog)
}
func (s Service) findByRequestId(id string) ([]TodoLog, error){
	return s.Repository.FindByRequestId(id)
}