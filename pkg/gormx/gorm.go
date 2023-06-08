package gormx

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB(models []interface{}) error {
	dsn := "root:root@tcp(localhost:13306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 开启调试
	DB.Debug()

	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("db connect faile", err)
		panic(err)
	}

	// 设置连接可以重用的最长时间(单位：秒)
	sqlDB.SetConnMaxLifetime(7200)

	// 设置空闲连接池中的最大连接数
	sqlDB.SetMaxIdleConns(50)

	// 设置数据库的最大打开连接数
	sqlDB.SetMaxOpenConns(500)

	err = DB.AutoMigrate(models...)
	if err != nil {

		fmt.Println("db AutoMigrate faile", err)
		panic(err)
	}

	return nil
}

func Tx(txFunc func(tx *gorm.DB) error) error {
	tx := DB.Begin()
	var err error
	if err = tx.Error; err != nil {
		return err
	}
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//		panic(r)
	//	} else {
	//		err = tx.Commit().Error
	//	}
	//}()
	err = txFunc(tx)
	return err
}

//func autoMigrate(models []interface{}) error {
//	return DB.AutoMigrate(models...)
//}
