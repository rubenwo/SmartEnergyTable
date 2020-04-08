package database

import "fmt"

type OperationError struct {
	operation string
}

func (err *OperationError) Error() string {
	return fmt.Sprintf("could not perform the: %s operation", err.operation)
}

type DownError struct{}

func (err *DownError) Error() string {
	return "could not connect to the database"
}

type CreateDatabaseError struct{}

func (err *CreateDatabaseError) Error() string {
	return "could not create database"
}

type NotImplementedDatabaseError struct {
	database string
}

func (err *NotImplementedDatabaseError) Error() string {
	return fmt.Sprintf("%s not implemented", err.database)
}
