package models

type User struct {
	Id       int    `orm:"column(id);auto" json:"id"`
	Name     string `orm:"column(name);size(64);" description:"name of the user" json:"name"`
	Email    string `orm:"column(email);size(64);" description:"email of the user" json:"email"`
	Password string `orm:"column(password);size(64);" description:"password of the user" json:"password"`
}
