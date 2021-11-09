package authorization

import (
  "github.com/beego/beego/v2/client/orm"
)

type Repository struct {
  Orm orm.Ormer
}

func (r Repository) Create(user User) (User, error) {
  _, err := r.Orm.Insert(&user)
  if err != nil {
    return User{}, err
  }
  return r.FindByEmail(user.Email)
}

func (r Repository) Find() ([]User, error) {
  var userList []User
  _, err := r.Orm.QueryTable("user").All(&userList)

  if err != nil {
    return []User{}, err
  }
  return  userList, err
}

func (r Repository) FindByEmail(email string) (User, error){
  user := User{}
  err := r.Orm.Raw("SELECT id,name,email, password FROM \"user\" WHERE email = ?", email).QueryRow(&user)
  return user, err
}

func (r Repository) findById(id int) (User, error){
  user := User{}
  err := r.Orm.Raw("SELECT id,name,email, password FROM \"user\" WHERE id = ?", id).QueryRow(&user)
  return user, err
}