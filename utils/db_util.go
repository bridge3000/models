package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type DbUtil struct {
}

func (this *DbUtil) OpenDB(user string, pwd string, host string, dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(user+":"+pwd+"@tcp("+host+":3306)/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err != nil {
		return nil, err
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

		return db, nil
	}
}
