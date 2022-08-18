package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

type Options struct {
	DSN                   string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              string
}

func NewSessionOrDie(opts *Options) *gorm.DB {

	var logLevel logger.LogLevel

	switch opts.LogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	default:
		logLevel = logger.Info
	}

	session, err := gorm.Open(
		mysql.Open(opts.DSN),
		&gorm.Config{
			SkipDefaultTransaction:                   false,
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   logger.Default.LogMode(logLevel),
			NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		},
	)
	if err != nil {
		panic(err)
	}

	sqlDB, err := session.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	db = session

	return session
}

// DB 获取 DB 实例, 如果没有初始化则会 panic. 最好不要直接使用此方法, 而是通过工厂方法获取.
func DB() *gorm.DB {
	return db
}
