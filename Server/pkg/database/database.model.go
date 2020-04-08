package database

type Database interface {
	Set(key string, value []byte) (string, error)
	Get(key string) ([]byte, error)
	Delete(key string) ([]byte, error)
}

func Factory(databaseName string) (Database, error) {
	switch databaseName {
	case "redis":
		return createRedisDatabase()
	default:
		return nil, &NotImplementedDatabaseError{database: databaseName}
	}
}
