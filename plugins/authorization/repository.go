package authorization

import (
  "github.com/beego/beego/v2/client/orm"
  "poc-plugin/internal/configuration/database"
)

type Repository struct {
  Orm orm.Ormer
}

func (r Repository) Create(user database.User) (database.User, error) {
  _, err := r.Orm.Insert(&user)
  if err != nil {
    return database.User{}, err
  }
  return r.FindByEmail(user.Email)
}

func (r Repository) Find() ([]database.User, error) {
  var userList []database.User
  _, err := r.Orm.QueryTable("user").All(&userList)

  if err != nil {
    return []database.User{}, err
  }
  return  userList, err
}

func (r Repository) FindByEmail(email string) (database.User, error){
  user := database.User{}
  err := r.Orm.Raw("SELECT id,name,email, password FROM \"user\" WHERE email = ?", email).QueryRow(&user)
  return user, err
}

func (r Repository) findById(id int) (database.User, error){
  user := database.User{}
  err := r.Orm.Raw("SELECT id,name,email, password FROM \"user\" WHERE id = ?", id).QueryRow(&user)
  return user, err
}