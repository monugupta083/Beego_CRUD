package persistence

import (
	"myproject/models"

	"github.com/astaxie/beego/orm"
)

func FindAllUsers() []*models.User {
	o := orm.NewOrm()
	var users []*models.User
	o.QueryTable("user").All(&users)
	return users
}

func FindUserById(id int) (*models.User, error) {
	o := orm.NewOrm()
	user := models.User{Id: id}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		return nil, err
	} else if err == orm.ErrMissPK {
		return nil, err
	} else {
		return &user, nil
	}
}

func CreateUser(user *models.User) error {
	o := orm.NewOrm()
	_, err := o.Insert(user)
	return err
}

func UpdateUser(user *models.User) error {
	o := orm.NewOrm()
	_, err := o.Update(user)
	return err
}

func DeleteUser(id int) error {
	o := orm.NewOrm()
	user := models.User{Id: id}
	_, err := o.Delete(&user)
	return err
}
