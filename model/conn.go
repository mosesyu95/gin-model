package model

import (
	"gin-model/config"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	DB     *gorm.DB
	Client redis.UniversalClient
)

func Init() {
	InitDB()
	InitRedis()
}

func InitDB() *gorm.DB {
	cfg := config.Config.DB
	db, err := gorm.Open("mysql", cfg.Dsn)
	if err != nil {
		log.Fatalln("fail to connect database, err: ", err.Error())
	}
	db.LogMode(config.Config.DB.Debug)
	db.DB().SetMaxIdleConns(cfg.MaxIdle)
	db.DB().SetMaxOpenConns(cfg.MaxConn)
	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}

func InitRedis() {
	cfg := config.Config.Redis
	if !cfg.Enable {
		return
	}
	var client redis.UniversalClient
	if len(cfg.Addrs) < 1 {
		log.Fatalln("please config redis address ! ")
	}
	switch cfg.Model {
	case "cluster":
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    cfg.Addrs,
			Password: cfg.Password,
		})

	case "normal":
		client = redis.NewClient(&redis.Options{
			Addr:     cfg.Addrs[0],
			Password: cfg.Password,
		})

	case "sentinel":
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    cfg.Master,
			SentinelAddrs: cfg.Addrs,
			Password:      cfg.Password,
		})
	default:
		log.Fatalln("unknown redis model,please select one of normal 、cluster、sentinel")
	}
	err := client.Ping().Err()
	if err != nil {
		log.Fatalln("connect redis failed please check configure")
	}

	Client = client
	log.Println("Init redis client success")
	return
}

func GetRedis() redis.UniversalClient {
	return Client
}
