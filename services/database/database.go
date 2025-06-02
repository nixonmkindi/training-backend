package database

import (
	"context"

	"sync"
	"training-backend/package/config"
	"training-backend/package/log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var once sync.Once

// variavel Global
var instance *pgxpool.Pool
var err error

func Connect() (*pgxpool.Pool, error) {
	//thread safe singletone initialised
	once.Do(func() {
		connectionString := config.GetDatabaseConnection()
		instance, err = pgxpool.Connect(context.Background(), connectionString)
		if err != nil {
			log.Errorf("unable to create a database instance")
		}

	})
	if err != nil {
		log.Errorf("unable to connect to database: %v\n", err)
		return nil, err
	}
	return instance, err
}

func Close() {
	instance.Close()
}
