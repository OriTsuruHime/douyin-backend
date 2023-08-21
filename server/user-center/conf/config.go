package conf

import (
	"github.com/spf13/viper"
	"strings"
	"user-center/cache"
	"user-center/dao"
	"user-center/server"
)

var (
	Address string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb     string
	redisAddr   string
	RedisPw     string
	RedisDbName string
)

// Init 初始化配置文件与引擎
func Init() error {

	// 设置配置文件的名称和路径
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// 获取配置项的值并赋值给变量
	Address = viper.GetString("server.address")
	Db = viper.GetString("mysql.DB")
	DbHost = viper.GetString("mysql.DbHost")
	DbPort = viper.GetString("mysql.DbPort")
	DbUser = viper.GetString("mysql.DbUser")
	DbPassword = viper.GetString("mysql.DbPassword")
	DbName = viper.GetString("mysql.DbName")
	RedisDb = viper.GetString("redis.RedisDb")
	redisAddr = viper.GetString("redis.RedisAddr")
	RedisPw = viper.GetString("redis.RedisPw")
	RedisDbName = viper.GetString("redis.RedisDbName")

	//mysql连接信息
	conn := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	// gorm引擎初始化
	err = dao.Database(conn)
	// redis引擎初始化
	err = cache.Redis(RedisDb, redisAddr, RedisPw, RedisDbName)
	// grpc初始化
	err = server.Grpc(Address)
	return nil
}
