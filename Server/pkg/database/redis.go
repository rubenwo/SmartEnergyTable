package database

import "github.com/go-redis/redis"

type redisDB struct {
	client *redis.Client
}

func createRedisDatabase() (Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "service.redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, &CreateDatabaseError{reason: err.Error()}
	}
	return &redisDB{client: client}, nil
}

func (r *redisDB) Set(key string, value []byte) (string, error) {
	return "", nil
}

func (r *redisDB) Get(key string) ([]byte, error) {
	return nil, nil
}

func (r *redisDB) Delete(key string) ([]byte, error) {
	return nil, nil
}
