package config

import (
	"myproject/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	// Register the MySQL driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Set the default database
	orm.RegisterDataBase("default", "mysql", "root:MyPassword@123@tcp(127.0.0.1:3306)/mydb?charset=utf8")

	// Register the User model
	orm.RegisterModel(new(models.User))

	// Create tables
	orm.RunSyncdb("default", false, true)
}
