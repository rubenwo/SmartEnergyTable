package database

import (
	"github.com/go-redis/redis"
	"time"
)

type redisDB struct {
	client *redis.Client
}

func createRedisDatabase() (Database, error) {
	//create a new client with default options
	//Addr is defined by the docker-compose.yml file
	client := redis.NewClient(&redis.Options{
		Addr:     "service.redis:6379",
		Password: "",
		DB:       0,
	})
	var err error
	//Try to Ping the database, sometimes this doesn't work right away so we try it 10 times with a timeout of 1sec
	//between tries.
	for i := 0; i < 10; i++ {
		_, err = client.Ping().Result()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return nil, &CreateDatabaseError{reason: err.Error()}
	}
	return &redisDB{client: client}, nil
}

//Set: Implementation of the database interface
func (r *redisDB) Set(key string, value string) (string, error) {
	_, err := r.client.Set(key, value, 0).Result()
	if err != nil {
		return generateError("set", err)
	}
	return key, nil
}

//Get: Implementation of the database interface
func (r *redisDB) Get(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		return generateError("get", err)
	}
	return value, nil
}

//Delete: Implementation of the database interface
func (r *redisDB) Delete(key string) (string, error) {
	_, err := r.client.Del(key).Result()
	if err != nil {
		return generateError("delete", err)
	}
	return key, nil
}

func generateError(operation string, err error) (string, error) {
	if err == redis.Nil {
		return "", &OperationError{operation: operation}
	}
	return "", &DownError{}
}
