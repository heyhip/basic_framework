package db

import (
	"basic_framework/configs"
	"basic_framework/core/log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Orm *gorm.DB

func init() {
	var err error

	// 数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.Yaml.Db.Username, configs.Yaml.Db.Password, configs.Yaml.Db.Host, configs.Yaml.Db.Port, configs.Yaml.Db.Database)

	// 数据库日志
	slowL := 1
	if configs.Yaml.Db.SlowThreshold != 0 {
		slowL = configs.Yaml.Db.SlowThreshold
	}
	logL := 4
	if configs.Yaml.Db.LogLevel != 0 {
		logL = configs.Yaml.Db.LogLevel
	}
	sqlLog := logger.New(gormLog{}, logger.Config{
		SlowThreshold: time.Duration(slowL) * time.Second,
		Colorful:      false,
		LogLevel:      logger.LogLevel(logL),
	})

	Orm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix: configs.Yaml.Db.Prefix,
			// 全局禁用表名复数
			SingularTable: true,
		},
		Logger:               sqlLog,
		DisableAutomaticPing: true,
	})

	if err != nil {
		log.Error(err)
	}

	db, err1 := Orm.DB()
	if err1 != nil {
		log.Error(err1)
	}

	if configs.Yaml.Db.MaxIdleConns != 0 {
		db.SetMaxIdleConns(configs.Yaml.Db.MaxIdleConns)
	} else {
		db.SetMaxIdleConns(10)
	}

	if configs.Yaml.Db.MaxOpenConns != 0 {
		db.SetMaxOpenConns(configs.Yaml.Db.MaxOpenConns)
	} else {
		db.SetMaxOpenConns(1000)
	}

	if configs.Yaml.Db.MaxLifetime != 0 {
		db.SetConnMaxIdleTime(time.Duration(configs.Yaml.Db.MaxLifetime) * time.Second)
	} else {
		db.SetConnMaxLifetime(3600 * time.Second)
	}

}

type gormLog struct {
}

func (this gormLog) Printf(_ string, data ...interface{}) {
	log.Info(data)
}
