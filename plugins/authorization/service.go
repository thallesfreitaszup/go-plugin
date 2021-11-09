package authorization

type Service struct {
	Repository Repository
}

type User struct {
Id int		`orm:"auto,column(id)"`
Name   string  `orm:"column(name)"`
Email string	 `orm:"column(email)"`
Password string `orm:"column(password)"`
}


func (s Service) Create(user User) (User, error) {
	return s.Repository.Create(user)
}

func (s Service) FindByEmail(username string) (User,error) {
	return s.Repository.FindByEmail(username)
}


func (s Service) FindById(id int) (User,error) {
	return s.Repository.findById(id)
}
