package dd

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"animeii.tech/model"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 设置日志参数 创建writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 启用日志文件
	if Cf.Mysql.EnableFileLogWriter {
		writer = &lumberjack.Logger{
			Filename:   Cf.Log.RootDir + "/" + Cf.Mysql.LogFilename,
			MaxSize:    Cf.Log.MaxSize,
			MaxBackups: Cf.Log.MaxBackups,
			MaxAge:     Cf.Log.MaxAge,
			Compress:   Cf.Log.Compress,
		}
	} else {
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

// 接管 gorm 日志
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch Cf.Mysql.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,        // 慢 SQL 阈值
		LogLevel:                  logMode,                       // 日志级别
		IgnoreRecordNotFoundError: false,                         // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !Cf.Mysql.EnableFileLogWriter, // 禁用彩色打印
	})
}

// 初始化 mysql gorm.DB
func initMySqlGorm() *gorm.DB {
	dbConfig := Cf.Mysql

	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
		Logger:                                   getGormLogger(), // 使用自定义 Logger
	}); err != nil {
		Log.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		// 创建数据库 迁移
		initMySqlTables(db)
		return db
	}
}

func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Article{},
		model.Comment{},
		model.Role{},
		model.UserRole{},
		model.Permission{},
		model.RolePermission{},
	)
	if err != nil {
		Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}

func InitGorm() {
	gormMysql := initMySqlGorm()
	DB = gormMysql
}
