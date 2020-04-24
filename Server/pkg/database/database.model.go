package database

//Database is an interface that defines the actions a database should be able to execute.
type Database interface {
	//Set takes a key and value to store in the database. If everything went well error is nil
	Set(key string, value string) (string, error)
	//Get takes a key and returns the value if error is not nil.
	Get(key string) (string, error)
	//Delete takes a key and deletes the value from the database. If everything went well error is nil
	Delete(key string) (string, error)
}

//Factory is a factory function that creates a database based on the databaseName param.
//eg. databaseName == 'redis' creates a Database that uses redis for it's backend.
func Factory(databaseName string) (Database, error) {
	switch databaseName {
	case "redis":
		return createRedisDatabase()
	default:
		return nil, &NotImplementedDatabaseError{database: databaseName}
	}
}
