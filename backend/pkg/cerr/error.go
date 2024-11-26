package cerr

import (
	"fmt"
)

//// UserType определяет тип пользователя
//type UserType string
//
//const (
//	User  UserType = "User"
//	Admin UserType = "Admin"
//
//	Work   UserType = "Work"
//)

//// LayerType определяет тип слоя
//type LayerType string
//
//const (
//	Repository LayerType = "REPO"
//	Service    LayerType = "SERV"
//	Handle     LayerType = "HAND"
//)

// ErrorType определяет тип ошибки
type ErrorType string

const (
	Transaction  ErrorType = "transaction error"
	Rollback     ErrorType = "rollback error"
	Commit       ErrorType = "commit error"
	Scan         ErrorType = "scan error"
	Execution    ErrorType = "execution error"
	ExecContext  ErrorType = "transaction.ExecContext error"
	Rows         ErrorType = "rows error"
	NoOneRow     ErrorType = "row count doesnt equals 1"
	InvalidLogin ErrorType = "invalid login"
	InvalidEmail ErrorType = "invalid email"
	InvalidPWD   ErrorType = "invalid password"
	InvalidCount ErrorType = "count more that have"
	InvalidType  ErrorType = "give not needn't name type"
	DiffPWD      ErrorType = "pwd not equal"
	Hash         ErrorType = "error in hashing time"
	NotFound     ErrorType = "this row not found"
	UserNotFound ErrorType = "this user not found"

	InvalidToken ErrorType = "invalid token"
	Parse        ErrorType = "parse error"
)

// CustomError структура для кастомной ошибки
type CustomError struct {
	Type ErrorType
	Err  error
}

// Error метод для реализации интерфейса error
func (e CustomError) Error() error {
	return fmt.Errorf("Error Type: %v, Error: %v", e.Type, e.Err)
}

// Str метод для реализации ошибки в строку
func (e CustomError) Str() string {
	return fmt.Sprintf("Error Type: %v, Error: %v", e.Type, e.Err)
}

// Err функция для создания новой кастомной ошибки
func Err(errType ErrorType, err error) CustomError {
	return CustomError{Type: errType, Err: err}
}
