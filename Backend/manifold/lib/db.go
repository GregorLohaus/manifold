package lib

import (
	"github.com/surrealdb/surrealdb.go"
	"strconv"
)

var db *surrealdb.DB

func GetDb(configPath *string) (*surrealdb.DB, error) {
	config, err := GetConfig(configPath)
	if err != nil {
		return nil, err
	}
	if db != nil {
		_, err = db.Signin(map[string]string{
			"user": config.Database.User,
			"pass": config.Database.Password,
		})
		return db, nil
	}
	dsn := "ws://" + config.Database.Host + ":" + strconv.Itoa(config.Database.Port) + "/rpc"
	i, err := surrealdb.New(dsn)
	if err != nil {
		return nil, err
	}
	db = i
	_, err = i.Signin(map[string]string{
		"user": config.Database.User,
		"pass": config.Database.Password,
	})
	if err != nil {
		return nil, err
	}
	return i, nil
}

func NewDb() (*surrealdb.DB, error) {
	config, err := GetConfig(nil)
	if err != nil {
		return nil, err
	}
	dsn := "ws://" + config.Database.Host + ":" + strconv.Itoa(config.Database.Port) + "/rpc"
	i, err := surrealdb.New(dsn)
	if err != nil {
		return nil, err
	}
	_, err = i.Signin(map[string]string{
		"user": config.Database.User,
		"pass": config.Database.Password,
	})
	if err != nil {
		return nil, err
	}
	return i, nil

}
