package datastore

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/ved2pj/Duanlink/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _datastore *DataStore

type DataStore struct {
	MySQL *gorm.DB
	Redis *redis.Client
}

func NewDatastore(cfg *config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	redisCli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_datastore = &DataStore{
		MySQL: db,
		Redis: redisCli,
	}
	return nil
}

func Get() *DataStore {
	return _datastore
}
