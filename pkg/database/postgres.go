package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go-jwt-template/pkg/config"
)

func GetDB() *sqlx.DB {
	connString := fmt.Sprintf("user=%v host=%v port=%v dbname=%v password=%v sslmode=disable",
		viper.GetString(config.DBUser),
		viper.GetString(config.DBHost),
		viper.GetInt(config.DBPort),
		viper.GetString(config.DBName),
		viper.GetString(config.DBPassword),
	)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		panic(fmt.Sprintf("Error while connecting to DB. Error: %v", err.Error()))
	}

	return db
}
