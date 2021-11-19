package log

import "poc-plugin/internal/configuration/database"

type Service struct {
	Repository Repository
}

func (s Service) CreateTaskLog(taskLog database.TaskLog) (int64, error) {
	return s.Repository.Create(taskLog)
}

func (s Service) CreateUserLog(userLog database.UserLog) (int64, error) {
	return s.Repository.CreateUserLog(userLog)
}
func (s Service) findByRequestId(id string) (database.UserLog, error) {
	return s.Repository.FindByRequestId(id)
}

func (s Service) GetTasks() []database.TaskLog {
	return s.Repository.getTasks()
}
