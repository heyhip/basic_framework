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

	f := configs.Yaml.Db.Db1
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", f.Username, f.Password, f.Host, f.Port, f.Database)
	Orm = initDb(dsn, f.MaxIdleConns, f.MaxOpenConns, f.MaxLifetime, f.SlowThreshold, f.LogLevel, f.Prefix)

}

func initDb(dsn string, maxIdleConns, maxOpenConns, maxLifetime, slowThreshold, logLevel int, prefix string) (gdb *gorm.DB) {
	var err error

	// 数据库日志
	slowL := 5
	if slowThreshold != 0 {
		slowL = slowThreshold
	}
	logL := 4
	if logLevel != 0 {
		logL = logLevel
	}
	sqlLog := logger.New(gormLog{}, logger.Config{
		SlowThreshold: time.Duration(slowL) * time.Second,
		Colorful:      false,
		LogLevel:      logger.LogLevel(logL),
	})

	gdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix: prefix,
			// 全局禁用表名复数
			SingularTable: true,
		},
		Logger:               sqlLog,
		DisableAutomaticPing: true,
	})

	if err != nil {
		log.Error(err)
	}

	db, err1 := gdb.DB()
	if err1 != nil {
		log.Error(err1)
	}

	if maxIdleConns != 0 {
		db.SetMaxIdleConns(maxIdleConns)
	} else {
		db.SetMaxIdleConns(10)
	}

	if maxOpenConns != 0 {
		db.SetMaxOpenConns(maxOpenConns)
	} else {
		db.SetMaxOpenConns(1000)
	}

	if maxLifetime != 0 {
		db.SetConnMaxIdleTime(time.Duration(maxLifetime) * time.Second)
	} else {
		db.SetConnMaxLifetime(3600 * time.Second)
	}
	return
}

type gormLog struct {
}

func (this gormLog) Printf(_ string, data ...interface{}) {
	log.Info(data)
}
