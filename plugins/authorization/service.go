package authorization

import (
	"poc-plugin/internal/configuration/database"
)

type Service struct {
	Repository Repository
}



func (s Service) Create(user database.User) (database.User, error) {
	return s.Repository.Create(user)
}

func (s Service) FindByEmail(username string) (database.User,error) {
	return s.Repository.FindByEmail(username)
}


func (s Service) FindById(id int) (database.User,error) {
	return s.Repository.findById(id)
}
