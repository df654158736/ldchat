package utils

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MyDB *gorm.DB

func SystemInit() {
	// 初始化数据库连接
	InitDB()

	// // 初始化Redis连接
	// InitRedis()

	// // 初始化日志
	// InitLog()

	// // 初始化配置
	// InitConfig()

	// // 初始化路由
	// InitRouter()
}

func InitDB() {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	dsn := viper.GetString("mysql.dns")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL 阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //彩色打印
		},
	)

	MyDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("failed to connect database")
	}
	println("init DB:", MyDB)
	println("init mysql:", dsn)
	// db.AutoMigrate(&models.UserBasic{})

}

func InitRouter() {
	panic("unimplemented")
}

func InitConfig() {
	panic("unimplemented")
}

func InitLog() {
	panic("unimplemented")
}

func InitRedis() {
	panic("unimplemented")
}
