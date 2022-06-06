package database

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	connectionStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DbUsername,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)

	DB, err := gorm.Open(mysql.Open(connectionStr))

	if err != nil {
		fmt.Println("error when open connection ", err)
	}

	err = DB.AutoMigrate(
		&model.Role{},
		&model.KategoriTherad{},
		&model.User{},
		&model.Ruang{},
		&model.Therad{},
		&model.Followers{},
		&model.MemberRuang{},
		&model.Comment{},
		&model.ReplyComment{},
		&model.Report{},
		&model.Like{},
	)
	if err != nil {
		fmt.Println("error when migration ", err)
	}

	return DB

}
