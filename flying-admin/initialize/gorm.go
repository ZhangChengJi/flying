package initialize

import (
	"flying-admin/global"
	"flying-admin/model"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

//@author: SliverHorn
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Node{},
		model.App{},
		model.AppNode{},
		model.Namespace{},
		model.User{},
		model.JwtBlacklist{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}

//
//@author: SliverHorn
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	var dsn string
	m := global.GVA_CONFIG.Mysql
	//连接先加载环境变量中
	mysqlUrl := os.Getenv("MYSQL_URL")
	if mysqlUrl != "" {
		fmt.Printf("优先加载环境变量mysql配置🐬")
		mysqlUsername := os.Getenv("MYSQL_USERNAME")
		mysqlPassword := os.Getenv("MYSQL_PASSWORD")
		mysqlDbName := os.Getenv("MYSQL_DBNAME")
		dsn = mysqlUsername + ":" + mysqlPassword + "@tcp(" + mysqlUrl + ")/" + mysqlDbName + "?" + m.Config
	} else {
		dsn = m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	}

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@author: SliverHorn
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig(mod bool) *gorm.Config {
	switch global.GVA_CONFIG.Mysql.LogZap {
	case "Silent":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Error":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Error),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Warn":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Warn),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Info":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	default:
		if mod {
			return &gorm.Config{
				Logger:                                   logger.Default.LogMode(logger.Info),
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		} else {
			return &gorm.Config{
				Logger:                                   logger.Default.LogMode(logger.Silent),
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		}
	}
}
